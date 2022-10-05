package router

import (
	"net/http"

	"github.com/WellyShadow/Golang_Homework/go-task4/controller"
)

func Createrouter() {
	handler := http.NewServeMux()
	handler.HandleFunc("/user", controller.Createuser)
}
