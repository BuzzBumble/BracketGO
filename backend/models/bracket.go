package models

import (
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

func GetBrackets(db *sqlx.DB) ([]Bracket, error) {
	brackets := []Bracket{} // array of brackets
	err := db.Select(&brackets, getBrackets)
	return brackets, err
}

func CreateBracket(db *sqlx.DB, data *Bracket) error {
	err := db.QueryRow(createBracket, data.Name).Scan(&data.Id)
	return err
}

func GetBracket(db *sqlx.DB, id string) (Bracket, error) {
	var b Bracket
	err := db.Get(&b, getBracket, id)
	return b, err
}

func UpdateBracket(db *sqlx.DB, id string, data *Bracket) error {
	_, err := db.Exec(updateBracket, id, data.Name)
	if err != nil {
		return err
	}

	return db.QueryRowx(getBracket, id).StructScan(data)
}

func DeleteBracket(db *sqlx.DB, id string) error {
	_, err := db.Exec(deleteBracket, id)
	return err
}
