package router

import (
	"net/http"

	"goflix/internal/config"
	"goflix/internal/database"
	"goflix/internal/handlers"
	"goflix/internal/middleware"

	"github.com/gorilla/mux"
)

type Router struct {
	config *config.Config
}

func NewRouter(cfg *config.Config) *Router {
	return &Router{config: cfg}
}

func (r *Router) Start() error {
	db, err := database.Connect(r.config.DBConnection)
	if err != nil {
		return err
	}
	defer db.Close()

	movieHandler := handlers.NewMovieHandler(db, r.config.JWTSecret)

	router := mux.NewRouter()
	router.Use(middleware.LoggingMiddleware)
	router.HandleFunc("/login", movieHandler.HandleLogin).Methods("POST")
	router.HandleFunc("/api/movies", middleware.JWTAuthMiddleware(movieHandler.HandleGetMovies, r.config.JWTSecret)).Methods("GET")
	router.HandleFunc("/api/movies", middleware.JWTAuthMiddleware(movieHandler.HandleCreateMovie, r.config.JWTSecret)).Methods("POST")

	return http.ListenAndServe(r.config.ServerAddress, router)
}
