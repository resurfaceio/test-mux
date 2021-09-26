package news

import (
	"log"

	database "github.com/resurfaceio/test-mux/internal/pkg/db"
)

type News struct {
	ID    string
	Title string
	Body  string
}

func (news News) Save() int64 {
	stmt, err := database.Db.Prepare("INSERT INTO News(Title,Body) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(news.Title, news.Body)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("Row Inserted in News with ID-", id, "\n")
	return id
}

func GetAll() []News {
	stmt, err := database.Db.Prepare("select ID, Title, Body from News")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var newss []News
	for rows.Next() {
		var nnews News
		err := rows.Scan(&nnews.ID, &nnews.Title, &nnews.Body)
		if err != nil {
			log.Fatal(err)
		}
		newss = append(newss, nnews)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return newss
}
