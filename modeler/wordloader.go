package modeler

import "database/sql"

// ---------------------------------------------------------------------
// Functions
// ---------------------------------------------------------------------

// LoadWords get the list of words from the history database
func LoadWords(filename string) ([]string, error) {

	// Connect to the database
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Issue a query for the words in the history table
	rows, err := db.Query(`SELECT word FROM history ORDER BY date`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Read from the query result into a slice of strings
	words := make([]string, 0)
	for rows.Next() {
		var word string
		if err := rows.Scan(&word); err != nil {
			return nil, err
		}
		words = append(words, word)
	}

	return words, nil
}
