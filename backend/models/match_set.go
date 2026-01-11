package models

import (
	"github.com/jmoiron/sqlx"
)

type MatchSet struct {
	Id             int `json:"id"`
	BracketId      int `json:"bracket_id" db:"bracket_id"`
	ParticipantAId int `json:"participantA_id" db:"participanta_id"`
	ParticipantBId int `json:"participantB_id" db:"participantb_id"`
}

var getMatchSets = "SELECT * FROM match_sets WHERE bracket_id = $1 ORDER BY id ASC"
var createMatchSet = "INSERT INTO match_sets (bracket_id, participanta_id, participantb_id) VALUES ($1, $2, $3) RETURNING id"
var getMatchSet = "SELECT * FROM match_sets WHERE id = $1 LIMIT 1"
var updateMatchSet = "UPDATE match_sets SET participanta_id = $2, participantb_id = $3 WHERE id = $1"
var deleteMatchSet = "DELETE FROM match_sets WHERE id = $1"

func GetMatchSets(db *sqlx.DB, bid string) ([]MatchSet, error) {
	ms := []MatchSet{}
	err := db.Select(&ms, getMatchSets, bid)

	return ms, err
}

func CreateMatchSet(db *sqlx.DB, data *MatchSet) error {
	err := db.QueryRowx(createMatchSet, data.BracketId, data.ParticipantAId, data.ParticipantBId).Scan(&data.Id)
	return err
}

func GetMatchSet(db *sqlx.DB, id string) (MatchSet, error) {
	var ms MatchSet
	err := db.Get(&ms, getMatchSet, id)
	return ms, err
}

func UpdateMatchSet(db *sqlx.DB, id string, data *MatchSet) error {
	_, err := db.Exec(updateMatchSet, id, data.ParticipantAId, data.ParticipantBId)
	if err != nil {
		return err
	}

	err = db.QueryRowx(getMatchSet, id).StructScan(data)
	return err
}

func DeleteMatchSet(db *sqlx.DB, id string) error {
	_, err := db.Exec(deleteMatchSet, id)
	return err
}
