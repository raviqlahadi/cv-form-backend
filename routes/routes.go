package routes

import (
	"github.com/gorilla/mux"
	"github.com/raviqlahadi/cv-form-backend/internal/handlers"
	"github.com/raviqlahadi/cv-form-backend/internal/repositories"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	// Initialize repositories
	userRepo := repositories.NewUserRepository()

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userRepo)

	// Public routes
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/profile", userHandler.Create).Methods("POST")
	apiRouter.HandleFunc("/profile/{id}", userHandler.GetByID).Methods("GET")
	apiRouter.HandleFunc("/profile/{id}", userHandler.Update).Methods("PUT")

	return router
}
