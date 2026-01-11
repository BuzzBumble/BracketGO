package models

import (
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

func GetParticipants(db *sqlx.DB, bid string) ([]Participant, error) {
	p := []Participant{}
	err := db.Select(&p, getParticipants, bid)

	return p, err
}

func CreateParticipant(db *sqlx.DB, data *Participant) error {
	err := db.QueryRowx(createParticipant, data.Name, data.BracketId).Scan(&data.Id)
	return err
}

func GetParticipant(db *sqlx.DB, id string) (Participant, error) {
	var p Participant
	err := db.Get(&p, getParticipant, id)
	return p, err
}

func UpdateParticipant(db *sqlx.DB, id string, data *Participant) error {
	_, err := db.Exec(updateParticipant, id, data.Name)
	if err != nil {
		return err
	}
	err = db.QueryRowx(getParticipant, id).StructScan(data)
	return err

}

func DeleteParticipant(db *sqlx.DB, id string) error {
	_, err := db.Exec(deleteParticipant, id)
	return err
}
