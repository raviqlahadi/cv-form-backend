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
	employmentRepo := repositories.NewEmploymentRepository()
	// Initialize handlers
	userHandler := handlers.NewUserHandler(userRepo)

	employmentHandler := handlers.NewEmploymentHandler(employmentRepo)

	// Public routes
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/profile", userHandler.Create).Methods("POST")
	apiRouter.HandleFunc("/profile/{id}", userHandler.GetByID).Methods("GET")
	apiRouter.HandleFunc("/profile/{id}", userHandler.Update).Methods("PUT")

	apiRouter.HandleFunc("/employment/{user_id}", employmentHandler.Create).Methods("POST")
	apiRouter.HandleFunc("/employment/{user_id}", employmentHandler.GetByUserID).Methods("GET")
	apiRouter.HandleFunc("/employment/{user_id}", employmentHandler.Update).Methods("PUT")
	apiRouter.HandleFunc("/employment/{user_id}", employmentHandler.Delete).Methods("DELETE")

	return router
}
