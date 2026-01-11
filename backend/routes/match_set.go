package routes

import (
	"bracketapi/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func GetMatchSets(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bid := vars["bracket_id"]
		ms, err := models.GetMatchSets(db, bid)
		if err != nil {
			SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		} else {
			SendJSONResponse(w, http.StatusOK, ms)
		}
	}
}

func CreateMatchSet(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ms models.MatchSet
		vars := mux.Vars(r)
		bid := vars["bracket_id"]
		json.NewDecoder(r.Body).Decode(&ms)
		ms.BracketId, _ = strconv.Atoi(bid)

		err := models.CreateMatchSet(db, &ms)
		if err != nil {
			SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		} else {
			SendJSONResponse(w, http.StatusOK, ms)
		}
	}
}

func GetMatchSet(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		ms, err := models.GetMatchSet(db, id)
		if err != nil {
			SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		} else {
			SendJSONResponse(w, http.StatusOK, ms)
		}
	}
}

func UpdateMatchSet(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ms models.MatchSet
		json.NewDecoder(r.Body).Decode(&ms)
		vars := mux.Vars(r)
		id := vars["id"]

		err := models.UpdateMatchSet(db, id, &ms)
		if err != nil {
			SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		} else {
			SendJSONResponse(w, http.StatusOK, ms)
		}
	}
}

func DeleteMatchSet(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		err := models.DeleteMatchSet(db, id)
		if err != nil {
			SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		} else {
			SendJSONResponse(w, http.StatusOK, nil)
		}
	}
}
