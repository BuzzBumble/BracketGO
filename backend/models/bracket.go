package models

import (
	"database/sql"
	"log"
)

type Bracket struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var getBrackets = "SELECT * FROM brackets"
var getBracket = "SELECT * FROM brackets WHERE id = $1"
var createBracket = "INSERT INTO brackets (name) VALUES ($1) RETURNING id"
var updateBracket = "UPDATE brackets SET name = $1 WHERE id = $2"
var deleteBracket = "DELETE FROM brackets WHERE id = $1"

func GetBrackets(db *sql.DB) []Bracket {
	rows, err := db.Query(getBrackets)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	brackets := []Bracket{} // array of brackets
	for rows.Next() {
		var b Bracket
		if err := rows.Scan(&b.Id, &b.Name); err != nil {
			log.Fatal(err)
		}
		brackets = append(brackets, b)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return brackets
}

func GetBracket(db *sql.DB, id string) Bracket {
	var b Bracket
	err := db.QueryRow(getBracket, id).Scan(&b.Id, &b.Name)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func CreateBracket(db *sql.DB, data *Bracket) *Bracket {
	err := db.QueryRow(createBracket, data.Name).Scan(&data.Id)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func UpdateBracket(db *sql.DB, id string, data *Bracket) Bracket {
	_, err := db.Exec(updateBracket, data.Name, id)
	if err != nil {
		log.Fatal(err)
	}

	var updatedBracket Bracket
	err = db.QueryRow(getBracket, id).Scan(&updatedBracket.Id, &updatedBracket.Name)
	if err != nil {
		log.Fatal(err)
	}

	return updatedBracket
}

func DeleteBracket(db *sql.DB, id string) string {
	var b Bracket
	err := db.QueryRow(getBracket, id).Scan(&b.Id, &b.Name)
	if err != nil {
		log.Fatal(err)
	} else {
		_, err := db.Exec(deleteBracket, id)
		if err != nil {
			log.Fatal(err)
		}

		return "Bracket Deleted"
	}
	return "Could not delete"
}
