package main

import (
	"bracketapi/models"
	"bracketapi/routes"
	"flag"
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
	dropFlag := flag.Bool("drop", false, "whether to drop tables before starting")

	flag.Parse()
	if *dropFlag {
		slog.Info("Dropping all tables...")
		for _, query := range models.SchemaDropQueries {
			tx.MustExec(query)
		}
	}

	slog.Info("Creating database schema...")
	for _, query := range models.SchemaCreateQueries {
		tx.MustExec(query)
	}
	tx.Commit()
	slog.Info("Schema successfully created")

	// create router
	router := mux.NewRouter()

	router.NotFoundHandler = http.HandlerFunc(NotFoundHandler)
	routes.RegisterMiddleware(router)
	slog.Info("Middleware successfully registered")
	routes.RegisterRoutes(router, db)
	slog.Info("Routes successfully registered")

	// start server
	slog.Info("Starting server on port " + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	slog.Info("Request",
		slog.String("Method", r.Method),
		slog.String("URL", r.URL.String()),
		slog.String("Status", "404 Not Found"),
	)
}
