package main

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
	"path/filepath"
)

var Connection *sql.DB

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	//Default CORS setup, exactly as specified in the chi/Cors documentation

	ConnectToDatabase()
	RouterSetup(r, Connection)
	defer Connection.Close()

	ReactBuildPath := "static"
	BuildDir := http.Dir(ReactBuildPath)
	StaticFileDir := http.Dir(filepath.Join(ReactBuildPath, "static"))

	r.Handle("/build/*", http.StripPrefix("/build/", http.FileServer(BuildDir)))
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(StaticFileDir)))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(ReactBuildPath, "index.html"))
	})

	http.ListenAndServe(":8080", r)
}
