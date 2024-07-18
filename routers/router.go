package routers

import (
	"go-template/controllers"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/items/get", controllers.GetItems).Methods("GET")
	router.HandleFunc("/items/getbyid/{id}", controllers.GetItem).Methods("GET")
	router.HandleFunc("/items/create", controllers.CreateItem).Methods("POST")
	router.HandleFunc("/items/update/{id}", controllers.UpdateItem).Methods("PUT")
	router.HandleFunc("/items/delete/{id}", controllers.DeleteItem).Methods("DELETE")

	return router
}
