package database

import (
	"fmt"
	"log"
)

func GetAllTimezones() {
	var (
		id int
		created string
		modified string
		name string
		identifier string
	)

	rows, err := db.Query("SELECT * FROM timezones")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &created, &modified, &name, &identifier)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name, identifier)
	}
}
