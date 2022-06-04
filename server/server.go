package server

import (
	"net/http"
	"seniorproject-backend/controller"
	"seniorproject-backend/repository"
	"seniorproject-backend/service"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) StartServer() error {
	repository := repository.NewAnimalRepository()
	service := service.NewAnimalService(repository)
	controller := controller.NewAnimalController(service)
	router := mux.NewRouter()

	// GET `/animals` = Returns list of `Animal`
	router.HandleFunc("/animals", controller.GetAnimals).Methods("GET")

	// GET `/animal/:animal_id/symptoms` = Returns list of `Symptom`
	router.HandleFunc("/animal/{animal_id}/symptoms", controller.GetSymptoms).Methods("GET")

	// GET `/action/:id` = Returns `Actionable`
	router.HandleFunc("/action/{id}", controller.GetActionable).Methods("GET")
	
	c := cors.New(cors.Options{
			AllowedOrigins: []string{
				"https://vet-heal.web.app",
				"http://vetheal.app",
				"https://vetheal.app",
				"http://localhost:3000",
			},
			AllowCredentials: true,
			// Enable Debugging for testing, consider disabling in production
	})

	handler := c.Handler(router)

	err := http.ListenAndServe(":8000", handler)
	return err
}
