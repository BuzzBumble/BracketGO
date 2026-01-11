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
		b, err := models.GetBrackets(db)

		if err != nil {
			SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		} else {
			SendJSONResponse(w, http.StatusOK, b)
		}
	}
}

// get Bracket by id
func GetBracket(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		b, err := models.GetBracket(db, id)
		if err != nil {
			SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		} else {
			SendJSONResponse(w, http.StatusOK, b)
		}
	}
}

// create Bracket
func CreateBracket(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var b models.Bracket
		json.NewDecoder(r.Body).Decode(&b)

		err := models.CreateBracket(db, &b)
		if err != nil {
			SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		} else {
			SendJSONResponse(w, http.StatusOK, b)
		}
	}
}

// update Bracket
func UpdateBracket(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var b models.Bracket
		json.NewDecoder(r.Body).Decode(&b)

		vars := mux.Vars(r)
		id := vars["id"]

		err := models.UpdateBracket(db, id, &b)
		if err != nil {
			SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		} else {
			SendJSONResponse(w, http.StatusOK, b)
		}
	}
}

// delete Bracket
func DeleteBracket(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		err := models.DeleteBracket(db, id)
		if err != nil {
			SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		} else {
			SendJSONResponse(w, http.StatusOK, nil)
		}
	}
}
