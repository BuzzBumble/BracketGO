package main

import (
	"database/sql"
	"log/slog"
	"log"
	"net/http"
	"os"
	"api/routes"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// main function
func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug, // This enables debug logs to be outputted
	}))
	slog.SetDefault(logger)
	slog.Debug("Test test")
	// connect to database
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	slog.Debug("Test test 2")

	// create table if it doesn't exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS brackets (id SERIAL PRIMARY KEY, name TEXT)")
	if err != nil {
		slog.Debug("Issue");
		log.Fatal(err)
	}
	
	slog.Debug("Test test 3")

	// create router
	router := mux.NewRouter()
	router.HandleFunc("/api/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Test!"))
	})
	router.HandleFunc("/api/go/brackets", routes.GetBrackets(db)).Methods("GET")
	router.HandleFunc("/api/go/brackets", routes.CreateBracket(db)).Methods("POST")
	router.HandleFunc("/api/go/brackets/{id}", routes.GetBracket(db)).Methods("GET")
	router.HandleFunc("/api/go/brackets/{id}", routes.UpdateBracket(db)).Methods("PUT")
	router.HandleFunc("/api/go/brackets/{id}", routes.DeleteBracket(db)).Methods("DELETE")

	// wrap the router with CORS and JSON content type middlewares
	enhancedRouter := enableCORS(jsonContentTypeMiddleware(router))
	// start server
	log.Fatal(http.ListenAndServe(":8000", enhancedRouter))
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow any origin
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Check if the request is for CORS preflight
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Pass down the request to the next middleware (or final handler)
		next.ServeHTTP(w, r)
	})
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	slog.Debug("Request handled")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set JSON Content-Type
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

