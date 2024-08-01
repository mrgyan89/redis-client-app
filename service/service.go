package service

import (
	"github.com/mrgyan89/redis-client-app/model"
	"github.com/mrgyan89/redis-client-app/repository"
)

func GetMessageByKey(key string) model.UserEntry {
	return repository.GetMessageByKey(key)
}

func GetAllMessages() []model.UserEntry {
	return repository.GetMessages()
}
func SetMessage(userEntry model.UserEntry) model.UserEntry {
	return repository.SetMessage(userEntry)
}
func UpdateMessage(userEntry model.UserEntry) model.UserEntry {
	return repository.SetMessage(userEntry)
}
func DeleteMessageByKey(key string) {
	repository.DeleteByKey(key)
}
