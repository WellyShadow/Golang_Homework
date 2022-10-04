package main

import (
	"log"
	"net/http"
	"time"
	"github.com/WellyShadow/Golang_Homework/go-task4/controller"
	_ "github.com/lib/pq"
)

//var usersMap map[string]string

func main() {

	handler := http.NewServeMux()
	handler.HandleFunc("/user", controller.createuser)
	//handler.HandleFunc("/user/login", loginuser)
	//usersMap = make(map[string]string)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
