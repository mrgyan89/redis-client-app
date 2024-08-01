package router

import (
	"github.com/gorilla/mux"
	"github.com/mrgyan89/redis-client-app/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/entries", controller.GetAllMessages).Methods("GET")
	router.HandleFunc("/api/entries", controller.SetMessage).Methods("POST")
	router.HandleFunc("/api/entries/{key}", controller.GetMessageByKey).Methods("GET")
	router.HandleFunc("/api/entries/{key}", controller.DeleteMessageByKey).Methods("DELETE")
	router.HandleFunc("/api/entries/{key}", controller.UpdateMessage).Methods("PATCH")
	return router
}
