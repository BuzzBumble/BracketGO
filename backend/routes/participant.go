package routes

import (
	"bracketapi/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"

	"github.com/gorilla/mux"
)

// get all Participants
func GetParticipants(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bracket_id := vars["bracket_id"]
		participants := models.GetParticipants(db, bracket_id)
		json.NewEncoder(w).Encode(participants)
	}
}

// create Participant
func CreateParticipant(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bracket_id := vars["bracket_id"]

		var p models.Participant
		json.NewDecoder(r.Body).Decode(&p)
		p.BracketId, _ = strconv.Atoi(bracket_id)

		models.CreateParticipant(db, &p)

		json.NewEncoder(w).Encode(p)
	}
}
