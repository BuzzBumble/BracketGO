package models

import (
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
)

type Participant struct {
	Id        int    `json:"id"`
	BracketId int    `json:"bracket_id"`
	Name      string `json:"name"`
}

var getParticipants = "SELECT * FROM participants WHERE bracket_id = $1"
var getParticipant = "SELECT * FROM participants WHERE id = $1"
var createParticipant = "INSERT INTO participants (name, bracket_id) VALUES ($1, $2) RETURNING id"
var updateParticipant = "UPDATE participants SET name = $1 WHERE id = $2"
var deleteParticipant = "DELETE FROM participants WHERE id = $1"

func GetParticipants(db *sqlx.DB, bracketId string) []Participant {
	rows, err := db.Query(getParticipants, bracketId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	participants := []Participant{} // array of brackets
	for rows.Next() {
		var p Participant
		if err := rows.Scan(&p.Id, &p.BracketId, &p.Name); err != nil {
			log.Fatal(err)
		}
		participants = append(participants, p)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return participants
}

func GetParticipant(db *sql.DB, id string) Participant {
	var b Participant
	err := db.QueryRow(getBracket, id).Scan(&b.Id, &b.Name)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func CreateParticipant(db *sqlx.DB, data *Participant) *Participant {
	err := db.QueryRow(createParticipant, data.Name, data.BracketId).Scan(&data.Id)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func UpdateParticipant(db *sql.DB, id string, data *Participant) Participant {
	_, err := db.Exec(updateBracket, data.Name, id)
	if err != nil {
		log.Fatal(err)
	}

	var updatedBracket Participant
	err = db.QueryRow(getBracket, id).Scan(&updatedBracket.Id, &updatedBracket.Name)
	if err != nil {
		log.Fatal(err)
	}

	return updatedBracket
}

func DeleteParticipant(db *sql.DB, id string) string {
	var b Participant
	err := db.QueryRow(getBracket, id).Scan(&b.Id, &b.Name)
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
