package main

import (
	"log"
	"net/http"
	"time"

	"github.com/WellyShadow/Golang_Homework/go-task4/repository"
	"github.com/WellyShadow/Golang_Homework/go-task4/router"
	"github.com/WellyShadow/Golang_Homework/go-task4/service"
)

//var usersMap map[string]string

func main() {

	//handler := http.NewServeMux()
	//handler.HandleFunc("/user", controller.Createuser)
	//handler.HandleFunc("/user/login", loginuser)
	//usersMap = make(map[string]string)
	service.PostgresConnection = repository.ConnectDBpostgres()
	//rep := repository.ConnectBD()
	handler := router.Createrouter()
	//var rep repository.Repository
	s := &http.Server{
		Addr:           ":8080",
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())

}
