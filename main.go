package main

import "os"

func main() {
	data, _ := os.ReadFile("testdata/answers.html")
	body := string(data)
	GetScrapes(body)
}