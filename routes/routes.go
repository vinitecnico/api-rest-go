package routes

import (
	"api-rest-go/controllers"

	"github.com/gorilla/mux"
)

func LoadRoat() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/contato", controllers.GetPeople).Methods("GET")
	router.HandleFunc("/contato/{id}", controllers.GetPerson).Methods("GET")
	router.HandleFunc("/contato/{id}", controllers.CreatePerson).Methods("POST")
	router.HandleFunc("/contato/{id}", controllers.DeletePerson).Methods("DELETE")
	return router
}
