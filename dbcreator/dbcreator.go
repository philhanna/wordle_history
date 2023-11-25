package dbcreator

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// CreateDatabase creates an sqlite3 database in the system temporary
// directory (/tmp on Linux) from the specified HTML data.
func CreateDatabase(body string, filename string) error {

	// Extract history data
	history := GetScrapes(body)

	log.Printf("Creating %q for %d history records...\n", filename, len(history))

	// Delete any old version
	os.Remove(filename)

	// Create the new database
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return err
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
		return err
	}

	// Insert all the records
	sql_insert := `INSERT INTO history VALUES(?, ?, ?);`
	stmt, err := db.Prepare(sql_insert)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, answer := range history {
		_, err := stmt.Exec(answer.Date, answer.Index, answer.Answer)
		if err != nil {
			return err
		}
	}

	// Done
	log.Printf("Done.\n")
	return nil
}
