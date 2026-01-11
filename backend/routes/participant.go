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
		p, err := models.GetParticipants(db, bracket_id)
		if err != nil {
			SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		} else {
			SendJSONResponse(w, http.StatusOK, p)
		}
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

		err := models.CreateParticipant(db, &p)

		if err != nil {
			SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		} else {
			SendJSONResponse(w, http.StatusOK, p)
		}
	}
}

func GetParticipant(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		p, err := models.GetParticipant(db, id)

		if err != nil {
			SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		} else {
			SendJSONResponse(w, http.StatusOK, p)
		}
	}
}

func UpdateParticipant(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p models.Participant
		json.NewDecoder(r.Body).Decode(&p)

		vars := mux.Vars(r)
		id := vars["id"]

		err := models.UpdateParticipant(db, id, &p)
		if err != nil {
			SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		} else {
			SendJSONResponse(w, http.StatusOK, p)
		}
	}
}

func DeleteParticipant(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		err := models.DeleteParticipant(db, id)
		if err != nil {
			SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		} else {
			SendJSONResponse(w, http.StatusOK, nil)
		}
	}
}
