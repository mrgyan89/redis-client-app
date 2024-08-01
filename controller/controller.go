package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mrgyan89/redis-client-app/model"
	"github.com/mrgyan89/redis-client-app/service"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the home page of redis-client-app"))
}
func GetMessageByKey(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	userKey := params["key"]
	userEntry := service.GetMessageByKey(userKey)
	json.NewEncoder(w).Encode(userEntry)
}

func GetAllMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userEntries := service.GetAllMessages()
	json.NewEncoder(w).Encode(userEntries)

}
func SetMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userEntry model.UserEntry
	json.NewDecoder(r.Body).Decode(&userEntry)
	userEntry = service.SetMessage(userEntry)
	json.NewEncoder(w).Encode(userEntry)
}
func UpdateMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	userKey := params["key"]
	var userEntryUpdate model.UserEntryUpdate
	json.NewDecoder(r.Body).Decode(&userEntryUpdate)
	var userEntry = model.UserEntry{
		Key:   userKey,
		Value: userEntryUpdate.Value,
	}
	userEntry = service.UpdateMessage(userEntry)
	json.NewEncoder(w).Encode(userEntry)
}
func DeleteMessageByKey(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	userKey := params["key"]
	service.DeleteMessageByKey(userKey)
	json.NewEncoder(w).Encode("Successfully deleted key: " + userKey)

}
