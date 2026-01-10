package routes

import (
	"bracketapi/models"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// get all brackets
func GetBrackets(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		brackets := models.GetBrackets(db)
		json.NewEncoder(w).Encode(brackets)
	}
}

// get Bracket by id
func GetBracket(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		b := models.GetBracket(db, id)

		json.NewEncoder(w).Encode(b)
	}
}

// create Bracket
func CreateBracket(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var b models.Bracket
		json.NewDecoder(r.Body).Decode(&b)

		models.CreateBracket(db, &b)

		json.NewEncoder(w).Encode(b)
	}
}

// update Bracket
func UpdateBracket(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var b models.Bracket
		json.NewDecoder(r.Body).Decode(&b)

		vars := mux.Vars(r)
		id := vars["id"]

		updatedBracket := models.UpdateBracket(db, id, &b)

		// Send the updated Bracket data in the response
		json.NewEncoder(w).Encode(updatedBracket)
	}
}

// delete Bracket
func DeleteBracket(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		res := models.DeleteBracket(db, id)

		json.NewEncoder(w).Encode(res)
	}
}
