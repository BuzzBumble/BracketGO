package models

import (
	"log"

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
var updateParticipant = "UPDATE participants SET name = $1 WHERE id = $2"
var deleteParticipant = "DELETE FROM participants WHERE id = $1"

func GetParticipants(db *sqlx.DB, bracketId string) []Participant {
	participants := []Participant{} // array of participants
	err := db.Select(&participants, getParticipants, bracketId)
	if err != nil {
		log.Fatal(err)
	}
	return participants
}

func GetParticipant(db *sqlx.DB, id string) Participant {
	var b Participant
	err := db.Get(&b, getBracket, id)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func CreateParticipant(db *sqlx.DB, data *Participant) *Participant {
	db.QueryRowx(createParticipant, data.Name, data.BracketId).Scan(&data.Id)

	return data
}

func UpdateParticipant(db *sqlx.DB, id string, data *Participant) Participant {
	_, err := db.Exec(updateBracket, data.Name, id)
	if err != nil {
		log.Fatal(err)
	}

	var updatedBracket Participant
	err = db.QueryRowx(getBracket, id).StructScan(&updatedBracket)
	if err != nil {
		log.Fatal(err)
	}

	return updatedBracket
}

func DeleteParticipant(db *sqlx.DB, id string) string {
	var b Participant
	err := db.QueryRowx(getBracket, id).StructScan(&b)
	if err != nil {
		log.Fatal(err)
	} else {
		_, err := db.Exec(deleteBracket, id)
		if err != nil {
			log.Fatal(err)
		}

		return "Participant Deleted"
	}
	return "Could not delete"
}
