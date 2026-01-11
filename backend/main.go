package main

import (
	"bracketapi/models"
	"bracketapi/routes"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// main function
func main() {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	slog.SetDefault(slog.New(handler))
	slog.Info("Server starting")
	// connect to database
	db, err := sqlx.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx := db.MustBegin()

	slog.Info("Creating database schema...")
	for _, query := range models.SchemaCreateQueries {
		tx.MustExec(query)
	}

	tx.Commit()
	slog.Info("Schema successfully created")

	// create router
	router := mux.NewRouter()

	routes.RegisterMiddleware(router)
	slog.Info("Middleware successfully registered")
	routes.RegisterRoutes(router, db)
	slog.Info("Routes successfully registered")

	// start server
	log.Fatal(http.ListenAndServe(":8000", router))
}
