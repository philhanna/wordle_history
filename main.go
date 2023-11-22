package main

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Get the history data
	log.Printf("Downloading history data...\n")
	data, _ := os.ReadFile("testdata/answers.html")
	body := string(data)
	history := GetScrapes(body)
	
	// Create the database
	tempDir := os.TempDir()
	filename := filepath.Join(tempDir, "wordle_history.db")
	log.Printf("Creating %q for %d history records...\n", filename, len(history))

	// Delete any old version
	os.Remove(filename)

	// Create the new database
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the wordle history table
	sql_create := `CREATE TABLE history (
		date	TEXT,
		puzzle	TEXT,
		word	TEXT
	);`
	_, err = db.Exec(sql_create)
	if err != nil {
		log.Fatal(err)
	}

	// Insert all the records
	sql_insert := `INSERT INTO history VALUES(?, ?, ?);`
	stmt, err := db.Prepare(sql_insert)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for _, answer := range history {
		_, err := stmt.Exec(answer.Date, answer.Index, answer.Answer)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Done
	log.Printf("Done.\n")
}