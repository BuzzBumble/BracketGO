package routes

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var queryGetBrackets = "SELECT * FROM brackets"
var queryGetBracket = "SELECT * FROM brackets WHERE id = $1"
var queryCreateBracket = "INSERT INTO brackets (name) VALUES ($1) RETURNING id"
var queryUpdateBracket = "UPDATE brackets SET name = $1 WHERE id = $3"
var queryDeleteBracket = "DELETE FROM brackets WHERE id = $1"

type Bracket struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
}

// get all brackets
func GetBrackets(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(queryGetBrackets)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		brackets := []Bracket{} // array of brackets
		for rows.Next() {
			var b Bracket
			if err := rows.Scan(&b.Id, &b.Name); err != nil {
				log.Fatal(err)
			}
			brackets = append(brackets, b)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(brackets)

	}

}

// get Bracket by id
func GetBracket(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var b Bracket
		err := db.QueryRow(queryGetBracket, id).Scan(&b.Id, &b.Name)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(b)
	}
}

// create Bracket
func CreateBracket(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var b Bracket
		json.NewDecoder(r.Body).Decode(&b)

		err := db.QueryRow(queryCreateBracket, b.Name).Scan(&b.Id)
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(b)
	}
}

// update Bracket
func UpdateBracket(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var b Bracket
		json.NewDecoder(r.Body).Decode(&b)

		vars := mux.Vars(r)
		id := vars["id"]

		// Execute the update query
		_, err := db.Exec(queryUpdateBracket, b.Name, id)
		if err != nil {
			log.Fatal(err)
		}

		// Retrieve the updated Bracket data from the database
		var updatedBracket Bracket
		err = db.QueryRow(queryGetBracket, id).Scan(&updatedBracket.Id, &updatedBracket.Name)
		if err != nil {
			log.Fatal(err)
		}

		// Send the updated Bracket data in the response
		json.NewEncoder(w).Encode(updatedBracket)
	}
}

// delete Bracket
func DeleteBracket(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var b Bracket
		err := db.QueryRow(queryGetBracket, id).Scan(&b.Id, &b.Name)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		} else {
			_, err := db.Exec(queryDeleteBracket, id)
			if err != nil {
				//todo : fix error handling
				w.WriteHeader(http.StatusNotFound)
				return
			}

			json.NewEncoder(w).Encode("Bracket deleted")
		}
	}
}
