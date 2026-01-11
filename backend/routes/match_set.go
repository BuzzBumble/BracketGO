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
