package routes

import (
	"bracketapi/models"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// get all Participants
func GetParticipants(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bracket_id := vars["bracket_id"]
		participants := models.GetParticipants(db, bracket_id)
		json.NewEncoder(w).Encode(participants)
	}
}

// create Participant
func CreateParticipant(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		vars := mux.Vars(r)
		bracket_id := vars["bracket_id"]

		var p models.Participant
		json.NewDecoder(r.Body).Decode(&p)

		p.BracketId, err = strconv.Atoi(bracket_id)
		if err != nil {
			log.Fatal(err)
		}

		models.CreateParticipant(db, &p)

		json.NewEncoder(w).Encode(p)
	}
}
