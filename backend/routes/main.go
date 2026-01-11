package routes

import (
	"bracketapi/middleware"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type StandardResponse struct {
	Data  interface{} `json:"data,omitempty"`
	Error interface{} `json:"error,omitempty"`
}

func SendJSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	response := StandardResponse{Data: data}
	json.NewEncoder(w).Encode(response)
}

func SendErrorResponse(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	e := map[string]string{"message": message}
	response := StandardResponse{Error: e}
	json.NewEncoder(w).Encode(response)
}

func RegisterMiddleware(r *mux.Router) {
	r.Use(middleware.RequestLoggerMiddleware)
	r.Use(middleware.EnableCORS)
	r.Use(middleware.JSONContentTypeMiddleware)
}

func RegisterRoutes(r *mux.Router, db *sqlx.DB) {
	// GET /brackets
	r.HandleFunc("/api/go/brackets", GetBrackets(db)).Methods("GET")
	// POST /brackets
	r.HandleFunc("/api/go/brackets", CreateBracket(db)).Methods("POST")
	// GET /brackets/:id
	r.HandleFunc("/api/go/brackets/{id}", GetBracket(db)).Methods("GET")
	// PUT /brackets/:id
	r.HandleFunc("/api/go/brackets/{id}", UpdateBracket(db)).Methods("PUT")
	// DELETE /brackets/:id
	r.HandleFunc("/api/go/brackets/{id}", DeleteBracket(db)).Methods("DELETE")

	// GET /brackets/:bracket_id/participants
	r.HandleFunc("/api/go/brackets/{bracket_id}/participants", GetParticipants(db)).Methods("GET")
	// POST /brackets/:bracket_id/participants/
	r.HandleFunc("/api/go/brackets/{bracket_id}/participants", CreateParticipant(db)).Methods("POST")
	// GET /participants/:id
	r.HandleFunc("/api/go/participants/{id}", GetParticipant(db)).Methods("GET")
	// PUT /participants/:id
	r.HandleFunc("/api/go/participants/{id}", UpdateParticipant(db)).Methods("PUT")
	// DELETE /participants/:id
	r.HandleFunc("/api/go/participants/{id}", DeleteParticipant(db)).Methods("DELETE")

	// GET /brackets/:bracket_id/match_sets
	r.HandleFunc("/api/go/brackets/{bracket_id}/match_sets", GetMatchSets(db)).Methods("GET")
	// POST /brackets/:bracket_id/match_sets
	r.HandleFunc("/api/go/brackets/{bracket_id}/match_sets", CreateMatchSet(db)).Methods("POST")
	// GET /match_sets/:id
	r.HandleFunc("/api/go/match_sets/{id}", GetMatchSet(db)).Methods("GET")
	// PUT /match_sets/:id
	r.HandleFunc("/api/go/match_sets/{id}", UpdateMatchSet(db)).Methods("PUT")
	// DELETE /match_sets/:id
	r.HandleFunc("/api/go/match_sets/{id}", DeleteMatchSet(db)).Methods("DELETE")

}
