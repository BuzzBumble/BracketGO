package main

import (
	"bracketapi/models"
	"bracketapi/routes"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// main function
func main() {
	log.Println("Start Server")
	// connect to database
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for _, query := range models.SchemaCreateQueries {
		_, err = db.Exec(query)
		if err != nil {
			log.Fatal(err)
		}
	}

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

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set JSON Content-Type
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
