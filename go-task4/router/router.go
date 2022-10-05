package router

import (
	"net/http"

	"github.com/WellyShadow/Golang_Homework/go-task4/controller"
)

func Createrouter() *http.ServeMux {
	handler := http.NewServeMux()
	handler.HandleFunc("/user", controller.Createuser)
	handler.HandleFunc("/user/login", controller.Loginuser)
	return handler
}
