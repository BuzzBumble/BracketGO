package routes

import (
	"bracketapi/models"
	"encoding/json"
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/gorilla/mux"
)

// get all brackets
func GetBrackets(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		brackets := models.GetBrackets(db)
		SendJSONResponse(w, http.StatusOK, brackets)
	}
}

// get Bracket by id
func GetBracket(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		b := models.GetBracket(db, id)

		SendJSONResponse(w, http.StatusOK, b)
	}
}

// create Bracket
func CreateBracket(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var b models.Bracket
		json.NewDecoder(r.Body).Decode(&b)

		models.CreateBracket(db, &b)

		SendJSONResponse(w, http.StatusOK, b)
	}
}

// update Bracket
func UpdateBracket(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var b models.Bracket
		json.NewDecoder(r.Body).Decode(&b)

		vars := mux.Vars(r)
		id := vars["id"]

		models.UpdateBracket(db, id, &b)

		// Send the updated Bracket data in the response
		SendJSONResponse(w, http.StatusOK, b)
	}
}

// delete Bracket
func DeleteBracket(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		models.DeleteBracket(db, id)
		SendJSONResponse(w, http.StatusOK, nil)
	}
}
