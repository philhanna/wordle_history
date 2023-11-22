# wordle

This website: [https://wordfinder.yourdictionary.com/wordle/answers/](https://wordfinder.yourdictionary.com/wordle/answers/) keeps a history of every Wordle word featured in the New York Times since 2021.  This project downloads and parses that data into an sqlite3 database with this table:
```sql
CREATE TABLE wordle_history (
    date    text, -- The puzzle date in YYYY-MM-DD format
    puzzle  text, -- The puzzle number. Usually numeric; some are alphabetic
    word    text  -- The 5-letter word, uppercase
)
```
