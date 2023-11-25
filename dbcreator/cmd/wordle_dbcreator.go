package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/philhanna/wordle_history/dbcreator"
)

const (
	URL   = "https://wordfinder.yourdictionary.com/wordle/answers"
	USAGE = `usage: wordle_dbcreator [filename]
Downloads wordle history and creates an sqlite3 database from it.

positional arguments:
  filename      (optional) name of database file to be created.
                Default is "wordle_history.db" in the system
                temporary directory.

options:
  -h, --help	displays this help text and exits
`
)

// Internal function to handle errors consistently
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Mainline
func main() {
	var (
		body []byte
		data string
		err  error
		resp *http.Response
	)

	// Set the log flags to include source file and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Get the name to be used for the database file
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, USAGE)
	}
	flag.Parse()

	var filename string
	if flag.NArg() > 0 {
		filename = flag.Arg(0)
	} else {
		tempDir := os.TempDir()
		filename = filepath.Join(tempDir, "wordle_history.db")
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
	err = dbcreator.CreateDatabase(data, filename)
	checkError(err)
}
