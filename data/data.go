package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/manifoldco/promptui"
)

var db *sql.DB

func OpenDatabase() error {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	pgurl := os.Getenv("PROCESS.ENV.pgURL")
	var err error
	db, err = sql.Open("postgres", pgurl)
	if err != nil {
		log.Fatal(err)
	}
	return db.Ping()
}

func CreateTable() {

	createTableSql := `CREATE TABLE studybuddy(
		"IdNote" SERIAL PRIMARY KEY,
		"word" TEXT,
		"definition" TEXT,
		"category" TEXT
	);`

	_, err := db.Exec(createTableSql)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Studybuddy table created")
}

func InsertNote(word string, definition string, category string) {
	insertNoteSQL := `INSERT INTO studybuddy(word,definition,category) VALUES($1,$2,$3)`
	statement, err := db.Prepare(insertNoteSQL)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = statement.Exec(word, definition, category)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Inserted study note successfully")

}

func DisplayAllNotes() {
	row, err := db.Query(`SELECT * FROM studybuddy ORDER BY word;`)
	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()

	for row.Next() {
		var idNote int
		var word string
		var definition string
		var category string
		row.Scan(&idNote, &word, &definition, &category)
		log.Println("[", category, "]", word, "-", definition)
	}
}

func DeleteNote() {
	row, err := db.Query(`SELECT * FROM studybuddy;`)
	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()

	type Note struct {
		IdNote     int
		Word       string
		Definition string
		Category   string
	}

	notes := make([]Note, 0)

	for row.Next() {
		var note Note
		err = row.Scan(&note.IdNote, &note.Word, &note.Definition, &note.Category)
		if err != nil {
			log.Fatal(err)
		}
		notes = append(notes, note)
	}

	if err := row.Err(); err != nil {
		log.Fatal(err)
	}

	items := make([]string, len(notes))
	for i, note := range notes {
		items[i] = fmt.Sprintf("%d. %s - %s (%s)", note.IdNote, note.Word, note.Definition, note.Category)
	}

	prompt := promptui.Select{
		Label: "Select a note to delete",
		Items: items,
	}
	index, _, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	noteID := notes[index].IdNote

	statement, err := db.Prepare(`DELETE FROM studybuddy WHERE idNote = $1`)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = statement.Exec(noteID)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Deleted study note successfully")
}
