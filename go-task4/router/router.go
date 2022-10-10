package router

import (
	"github.com/WellyShadow/Golang_Homework/go-task4/controller"
	"github.com/gorilla/mux"
)

func Createrouter() *mux.Router {
	handler := mux.NewRouter()
	handler.HandleFunc("/users/{id}", controller.GetUser)
	handler.HandleFunc("/user", controller.User)
	handler.HandleFunc("/user/phone", controller.Phoneuser)

	return handler
}
