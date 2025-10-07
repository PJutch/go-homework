package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetVersion(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(body), nil
}

type jsonInput struct {
	InputString string `json:"inputString"`
}

type jsonOutput struct {
	OutputString string `json:"outputString"`
}

func SendDecodeRequest(url string, input string) (string, error) {
	reqBody, err := json.Marshal(jsonInput{input})
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(reqBody))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var output jsonOutput
	if err := json.Unmarshal(respBody, &output); err != nil {
		fmt.Println(err)
		return "", nil
	}

	return string(output.OutputString), nil
}

func RequestHardOp(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	return resp.Status, nil
}

func main() {
	version, err := GetVersion("http://:8080/version")
	if err == nil {
		fmt.Println(version)
	} else {
		fmt.Println(err)
	}

	decoded, err := SendDecodeRequest("http://:8080/decode", "dGVzdA==")
	if err == nil {
		fmt.Println(decoded)
	} else {
		fmt.Println(err)
	}

	status, err := RequestHardOp("http://:8080/hard-op")
	if err == nil {
		fmt.Println(status)
	} else {
		fmt.Println(err)
	}
}
