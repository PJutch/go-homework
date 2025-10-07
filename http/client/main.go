package main

import (
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

func main() {
	version, err := GetVersion("http://:8080/version")
	if err == nil {
		fmt.Println(version)
	} else {
		fmt.Println(err)
	}
}
