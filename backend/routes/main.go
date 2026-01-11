package routes

import (
	"bracketapi/middleware"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

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

	// GET /brackets/:id/participants
	r.HandleFunc("/api/go/brackets/{bracket_id}/participants", GetParticipants(db)).Methods("GET")
	// POST /brackets/:id/participants/
	r.HandleFunc("/api/go/brackets/{bracket_id}/participants", CreateParticipant(db)).Methods("POST")
	// GET /participants/:id
	r.HandleFunc("/api/go/participants/:id", GetParticipant(db)).Methods("GET")
	// PUT /participants/:id
	r.HandleFunc("/api/go/participants/:id", UpdateParticipant(db)).Methods("PUT")
	// DELETE /participants/:id
	r.HandleFunc("/api/go/participants/:id", DeleteParticipant(db)).Methods("DELETE")

}
