package models

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type Bracket struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var getBrackets = "SELECT * FROM brackets ORDER BY id ASC"
var getBracket = "SELECT * FROM brackets WHERE id = $1"
var createBracket = "INSERT INTO brackets (name) VALUES ($1) RETURNING id"
var updateBracket = "UPDATE brackets SET name = $1 WHERE id = $2"
var deleteBracket = "DELETE FROM brackets WHERE id = $1"

func GetBrackets(db *sqlx.DB) []Bracket {
	brackets := []Bracket{} // array of brackets
	db.Select(&brackets, getBrackets)
	return brackets
}

func GetBracket(db *sqlx.DB, id string) Bracket {
	var b Bracket
	err := db.QueryRow(getBracket, id).Scan(&b.Id, &b.Name)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func CreateBracket(db *sqlx.DB, data *Bracket) *Bracket {
	err := db.QueryRow(createBracket, data.Name).Scan(&data.Id)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func UpdateBracket(db *sqlx.DB, id string, data *Bracket) Bracket {
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

func DeleteBracket(db *sqlx.DB, id string) string {
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
