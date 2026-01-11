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
		ms := models.GetMatchSets(db, bid)
		json.NewEncoder(w).Encode(ms)
	}
}

func CreateMatchSet(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ms models.MatchSet
		vars := mux.Vars(r)
		bid := vars["bracket_id"]
		json.NewDecoder(r.Body).Decode(&ms)
		ms.BracketId, _ = strconv.Atoi(bid)

		models.CreateMatchSet(db, &ms)

		json.NewEncoder(w).Encode(ms)
	}
}

func GetMatchSet(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		ms := models.GetMatchSet(db, id)

		json.NewEncoder(w).Encode(ms)
	}
}

func UpdateMatchSet(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ms models.MatchSet
		json.NewDecoder(r.Body).Decode(&ms)
		vars := mux.Vars(r)
		id := vars["id"]

		models.UpdateMatchSet(db, id, &ms)

		json.NewEncoder(w).Encode(ms)
	}
}

func DeleteMatchSet(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		models.DeleteMatchSet(db, id)
	}
}
