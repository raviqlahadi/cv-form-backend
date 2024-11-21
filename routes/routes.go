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
	educationRepo := repositories.NewEducationRepository()
	skillRepo := repositories.NewSkillRepository()

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userRepo)
	employmentHandler := handlers.NewEmploymentHandler(employmentRepo)
	educationHandler := handlers.NewEducationHandler(educationRepo)
	skillHandler := handlers.NewSkillHandler(skillRepo)

	// Public routes
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/profile", userHandler.Create).Methods("POST")
	apiRouter.HandleFunc("/profile/{id}", userHandler.GetByID).Methods("GET")
	apiRouter.HandleFunc("/profile/{id}", userHandler.Update).Methods("PUT")

	apiRouter.HandleFunc("/employment/{user_id}", employmentHandler.Create).Methods("POST")
	apiRouter.HandleFunc("/employment/{user_id}", employmentHandler.GetByUserID).Methods("GET")
	apiRouter.HandleFunc("/employment/{user_id}", employmentHandler.Update).Methods("PUT")
	apiRouter.HandleFunc("/employment/{user_id}", employmentHandler.Delete).Methods("DELETE")

	apiRouter.HandleFunc("/education/{user_id}", educationHandler.Create).Methods("POST")
	apiRouter.HandleFunc("/education/{user_id}", educationHandler.GetByUserID).Methods("GET")
	apiRouter.HandleFunc("/education/{user_id}", educationHandler.Update).Methods("PUT")
	apiRouter.HandleFunc("/education/{user_id}", educationHandler.Delete).Methods("DELETE")

	apiRouter.HandleFunc("/skill/{user_id}", skillHandler.Create).Methods("POST")
	apiRouter.HandleFunc("/skill/{user_id}", skillHandler.GetByUserID).Methods("GET")
	apiRouter.HandleFunc("/skill/{user_id}", skillHandler.Update).Methods("PUT")
	apiRouter.HandleFunc("/skill/{user_id}", skillHandler.Delete).Methods("DELETE")

	apiRouter.HandleFunc("/working-experience/{user_id}", userHandler.UpdateWorkingExperience).Methods("PUT")
	apiRouter.HandleFunc("/working-experience/{user_id}", userHandler.GetWorkingExperience).Methods("GET")

	return router
}
