package models

import (
	"log/slog"

	"github.com/jmoiron/sqlx"
)

type Bracket struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var getBrackets = "SELECT * FROM brackets ORDER BY id ASC"
var getBracket = "SELECT * FROM brackets WHERE id = $1 LIMIT 1"
var createBracket = "INSERT INTO brackets (name) VALUES ($1) RETURNING id"
var updateBracket = "UPDATE brackets SET name = $2 WHERE id = $1"
var deleteBracket = "DELETE FROM brackets WHERE id = $1"

func GetBrackets(db *sqlx.DB) []Bracket {
	brackets := []Bracket{} // array of brackets
	err := db.Select(&brackets, getBrackets)
	if err != nil {
		slog.Error(err.Error())
	}
	return brackets
}

func CreateBracket(db *sqlx.DB, data *Bracket) *Bracket {
	err := db.QueryRow(createBracket, data.Name).Scan(&data.Id)
	if err != nil {
		slog.Error(err.Error())
	}
	return data
}

func GetBracket(db *sqlx.DB, id string) Bracket {
	var b Bracket
	err := db.Get(&b, getBracket, id)
	if err != nil {
		slog.Error(err.Error())
	}
	return b
}

func UpdateBracket(db *sqlx.DB, id string, data *Bracket) {
	_, err := db.Exec(updateBracket, id, data.Name)
	if err != nil {
		slog.Error(err.Error())
	}

	err = db.QueryRowx(getBracket, id).StructScan(data)
	if err != nil {
		slog.Error(err.Error())
	}
}

func DeleteBracket(db *sqlx.DB, id string) {
	_, err := db.Exec(deleteBracket, id)
	if err != nil {
		slog.Error(err.Error())
	}
}
