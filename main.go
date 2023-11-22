package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const URL = "https://wordfinder.yourdictionary.com/wordle/answers"

func main() {
	var (
		body []byte
		data string
		err  error
		resp *http.Response
	)

	// Set the log flags to include source file and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Internal function to handle errors consistently
	checkError := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}

	// Download the data
	log.Printf("Downloading answer data HTML from website...\n")
	resp, err = http.Get(URL)
	checkError(err)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		err = fmt.Errorf("return code %d from website", resp.StatusCode)
		log.Fatal(err)
	}
	body, err = io.ReadAll(resp.Body)
	checkError(err)

	// Convert the byte slice to string
	data = string(body)

	// Create the database from the body
	err = CreateDatabase(data)
	checkError(err)
}
