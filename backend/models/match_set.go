package models

import (
	"log/slog"

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
var updateMatchSet = "UPDATE match_sets SET participanta_id = $2, participantb_id = $3 WHERE id = $1 RETURNING participanta_id, participantb_id"
var deleteMatchSet = "DELETE FROM match_sets WHERE id = $1"

func GetMatchSets(db *sqlx.DB, bid string) []MatchSet {
	ms := []MatchSet{}
	err := db.Select(&ms, getMatchSets, bid)
	if err != nil {
		slog.Error(err.Error())
	}

	return ms
}

func CreateMatchSet(db *sqlx.DB, data *MatchSet) *MatchSet {
	err := db.QueryRowx(createMatchSet, data.BracketId, data.ParticipantAId, data.ParticipantBId).Scan(&data.Id)
	if err != nil {
		slog.Error(err.Error())
	}
	return data
}

func GetMatchSet(db *sqlx.DB, id string) *MatchSet {
	var ms MatchSet
	err := db.Get(&ms, getMatchSet, id)
	if err != nil {
		slog.Error(err.Error())
	}
	return &ms
}

func UpdateMatchSet(db *sqlx.DB, id string, data *MatchSet) {
	db.QueryRowx(updateMatchSet, id, data.ParticipantAId, data.ParticipantBId)

	db.QueryRowx(getMatchSet, id).StructScan(data)
}

func DeleteMatchSet(db *sqlx.DB, id string) {
	_, err := db.Exec(deleteMatchSet, id)
	if err != nil {
		slog.Error(err.Error())
	}
}
