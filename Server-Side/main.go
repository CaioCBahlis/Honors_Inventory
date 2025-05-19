package main

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	ConnectToDatabase()
	defer Connection.Close()

	ReactBuildPath := filepath.Join("..", "honors-client-side", "build")
	BuildDir := http.Dir(ReactBuildPath)
	StaticFileDir := http.Dir(filepath.Join(ReactBuildPath, "static"))

	r.Handle("/build/*", http.StripPrefix("/build/", http.FileServer(BuildDir)))
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(StaticFileDir)))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(ReactBuildPath, "index.html"))
	})

	r.Get("/heartbeat", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<3"))
	})

	http.ListenAndServe(":8080", r)
}
