package models

import (
	"log"
	"log/slog"

	"github.com/jmoiron/sqlx"
)

type Participant struct {
	Id        int    `json:"id"`
	BracketId int    `json:"bracket_id" db:"bracket_id"`
	Name      string `json:"name"`
}

var getParticipants = "SELECT * FROM participants WHERE bracket_id = $1 ORDER BY id ASC"
var getParticipant = "SELECT * FROM participants WHERE id = $1 LIMIT 1"
var createParticipant = "INSERT INTO participants (name, bracket_id) VALUES ($1, $2) RETURNING id"
var updateParticipant = "UPDATE participants SET name = $2 WHERE id = $1"
var deleteParticipant = "DELETE FROM participants WHERE id = $1"

func GetParticipants(db *sqlx.DB, bid string) []Participant {
	p := []Participant{}
	err := db.Select(&p, getParticipants, bid)
	if err != nil {
		slog.Error(err.Error())
	}
	return p
}

func CreateParticipant(db *sqlx.DB, data *Participant) *Participant {
	err := db.QueryRowx(createParticipant, data.Name, data.BracketId).Scan(&data.Id)

	if err != nil {
		slog.Error(err.Error())
	}
	return data
}

func GetParticipant(db *sqlx.DB, id string) Participant {
	var p Participant
	err := db.Get(&p, getParticipant, id)
	if err != nil {
		slog.Error(err.Error())
	}
	return p
}

func UpdateParticipant(db *sqlx.DB, id string, data *Participant) {
	_, err := db.Exec(updateParticipant, id, data.Name)
	if err != nil {
		log.Fatal(err)
	}

	err = db.QueryRowx(getParticipant, id).StructScan(data)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteParticipant(db *sqlx.DB, id string) {
	_, err := db.Exec(deleteParticipant, id)
	if err != nil {
		slog.Error(err.Error())
	}
}
