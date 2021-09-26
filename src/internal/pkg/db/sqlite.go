package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func InitDB() {

	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatal("DB creation failed, error: ", err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal("DB ping failed, error: ", err)
	}

	Db = db
}

func Migrate() {
	commands := make([]string, 0)

	commands = append(commands, `
	CREATE TABLE IF NOT EXISTS News(
		ID INTEGER PRIMARY KEY UNIQUE,
		Title VARCHAR (255) ,
		Body VARCHAR (255)
	)`)

	for _, c := range commands {
		_, err := Db.Exec(c)
		if err != nil {
			log.Fatal("DB table creation failed, error: ", err)
		}
	}
}

func Populate() {

	type news struct {
		title string
		body  string
	}

	newsToAdd := make([]news, 0)
	newsToAdd = append(newsToAdd, news{
		title: "Resurface Site",
		body:  "Hello from the other side",
	})

	for _, li := range newsToAdd {
		stmt, err := Db.Prepare("INSERT INTO News(Title,Body) VALUES(?,?)")
		if err != nil {
			log.Fatal(err)
		}

		_, err = stmt.Exec(li.title, li.body)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Print("Successfuly added ", len(newsToAdd), " rows to the News table...\n")
}

func Clear() {
	commands := make([]string, 0)
	commands = append(commands, `DROP TABLE News`)

	for _, c := range commands {
		_, err := Db.Exec(c)
		if err != nil {
			log.Fatal("DB table deletion failed, error: ", err)
		}
	}
}

func Truncate() {
	commands := make([]string, 0)
	commands = append(commands, `Delete from News`)

	for _, c := range commands {
		_, err := Db.Exec(c)
		if err != nil {
			log.Fatal("DB table deletion failed, error: ", err)
		}
	}
}
