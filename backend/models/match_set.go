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

func GetMatchSets(db *sqlx.DB, bid string) []MatchSet {
	ms := []MatchSet{}
	err := db.Select(&ms, getMatchSets, bid)
	if err != nil {
		slog.Info(err.Error())
	}

	return ms
}

func CreateMatchSet(db *sqlx.DB, data *MatchSet) *MatchSet {
	db.QueryRowx(createMatchSet, data.BracketId, data.ParticipantAId, data.ParticipantBId).Scan(&data.Id)

	return data
}
