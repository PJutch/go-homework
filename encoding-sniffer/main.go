package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"unicode"
	"unicode/utf16"
	"unicode/utf8"
)

func decodeUtf8(encoded []byte) []rune {
	runes := make([]rune, 0)
	for len(encoded) > 0 {
		rune_, runeLen := utf8.DecodeRune(encoded)
		runes = append(runes, rune_)
		encoded = encoded[runeLen:]
	}
	return runes
}

func decodeUtf16(encoded []byte) []rune {
	codePoints := make([]uint16, 0)
	for i := 0; i < len(encoded); i += 2 {
		var codePoint uint16
		if i+1 < len(encoded) {
			codePoint = uint16(encoded[i])<<8 + uint16(encoded[i+1])
		} else {
			// i assume we should be resistant encoding errors so just generate some garbage
			codePoint = uint16(encoded[i]) << 8
		}

		codePoints = append(codePoints, codePoint)
	}
	return utf16.Decode(codePoints)
}

func toLowercase(runes []rune) []rune {
	lowercase := make([]rune, len(runes))
	for i, rune_ := range runes {
		lowercase[i] = unicode.ToLower(rune_)
	}
	return lowercase
}

func computeFrequencies(decoded []rune) map[rune]float64 {
	freqs := make(map[rune]float64)

	for _, rune_ := range decoded {
		freqs[rune_] += 1
	}

	for rune_ := range freqs {
		freqs[rune_] /= float64(len(decoded))
		freqs[rune_] *= 100
	}

	return freqs
}

func square(x float64) float64 {
	return x * x
}

func keyUnion(map1 map[rune]float64, map2 map[rune]float64) []rune {
	keys := make([]rune, 0)

	for key := range map1 {
		keys = append(keys, key)
	}

	for key := range map2 {
		keys = append(keys, key)
	}

	return keys
}

func computeError(found_freqs map[rune]float64, expected_freqs map[rune]float64) float64 {
	sumSquareDiff := 0.0
	keys := keyUnion(found_freqs, expected_freqs)
	for _, rune_ := range keys {
		sumSquareDiff += square(found_freqs[rune_] - expected_freqs[rune_])
	}
	return math.Sqrt(sumSquareDiff / float64(len(keys)))
}

func sniffDecode(encoded []byte) (decoderName string, decoded string) {
	decoders := map[string]func([]byte) []rune{"utf-8": decodeUtf8, "utf-16": decodeUtf16}
	russianFreqs := map[rune]float64{
		'о': 10.97, 'е': 8.45, 'а': 8.01, 'и': 7.35, 'н': 6.70, 'т': 6.26,
		'с': 5.47, 'р': 4.73, 'в': 4.54, 'л': 4.40, 'к': 3.49, 'м': 3.21,
		'д': 2.98, 'п': 2.81, 'у': 2.62, 'я': 2.01, 'ы': 1.90, 'ь': 1.74,
		'г': 1.70, 'з': 1.65, 'б': 1.59, 'ч': 1.44, 'й': 1.21, 'х': 0.97,
		'ж': 0.94, 'ш': 0.73, 'ю': 0.64, 'ц': 0.48, 'щ': 0.36, 'э': 0.32,
		'ф': 0.26, 'ъ': 0.04, 'ё': 0.04,
	}

	minError := 1e10
	bestDecoderName := "none"
	bestDecoding := make([]rune, 0)

	for decoderName, decoder := range decoders {
		runes := decoder(encoded)
		preprocessedRunes := toLowercase(runes)
		freqs := computeFrequencies(preprocessedRunes)
		error_ := computeError(freqs, russianFreqs)
		if error_ < minError {
			minError = error_
			bestDecoderName = decoderName
			bestDecoding = runes
		}
	}

	return bestDecoderName, string(bestDecoding)
}

func main() {
	if len(os.Args) > 2 || len(os.Args) == 2 && os.Args[1] == "--help" {
		programName := filepath.Base(os.Args[0])
		fmt.Printf("Usage:\n    %s <filename>\nor\n    %s\nand provide input in stdin\n", programName, programName)
	} else if len(os.Args) < 2 {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		decoderName, decoded := sniffDecode([]byte(input))
		fmt.Println(decoderName)
		fmt.Println(decoded)
	} else {
		input, err := os.ReadFile(os.Args[1])
		if err != nil {
			fmt.Printf("error reading the file %s: %e", input, err)
			return
		}

		decoderName, decoded := sniffDecode(input)
		fmt.Println(decoderName)
		fmt.Println(decoded)
	}
}
