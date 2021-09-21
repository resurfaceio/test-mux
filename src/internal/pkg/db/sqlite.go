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
	CREATE TABLE IF NOT EXISTS Links(
		ID INTEGER PRIMARY KEY UNIQUE,
		Title VARCHAR (255) ,
		Address VARCHAR (255)
	)`)

	for _, c := range commands {
		_, err := Db.Exec(c)
		if err != nil {
			log.Fatal("DB table creation failed, error: ", err)
		}
	}
}

func Populate() {

	type link struct {
		title   string
		address string
	}

	linksToAdd := make([]link, 0)
	linksToAdd = append(linksToAdd, link{
		title:   "Resurface Site",
		address: "https://resurface.io",
	})
	linksToAdd = append(linksToAdd, link{
		title:   "Google Home",
		address: "https://google.com",
	})
	linksToAdd = append(linksToAdd, link{
		title:   "Splunk Site",
		address: "https://www.splunk.com/",
	})
	linksToAdd = append(linksToAdd, link{
		title:   "Datadog Site",
		address: "https://www.datadoghq.com/",
	})
	linksToAdd = append(linksToAdd, link{
		title:   "GraphQL Site",
		address: "https://graphql.org/",
	})

	for _, li := range linksToAdd {
		stmt, err := Db.Prepare("INSERT INTO Links(Title,Address) VALUES(?,?)")
		if err != nil {
			log.Fatal(err)
		}

		_, err = stmt.Exec(li.title, li.address)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Print("Successfuly added ", len(linksToAdd), " rows to the Links table...\n")
}

func Clear() {
	commands := make([]string, 0)
	commands = append(commands, `DROP TABLE Links`)

	for _, c := range commands {
		_, err := Db.Exec(c)
		if err != nil {
			log.Fatal("DB table deletion failed, error: ", err)
		}
	}
}
