package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const URL = "https://wordfinder.yourdictionary.com/wordle/answers"

func FromTestData() (string, error) {
	data, err := os.ReadFile("testdata/answers.html")
	return string(data), err
}

func FromWebSite() (string, error) {
	resp, err := http.Get(URL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("return code %d from website", resp.StatusCode)
	}
	data, err := io.ReadAll(resp.Body)
	return string(data), err
}
