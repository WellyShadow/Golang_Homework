package controller

import (
	"fmt"
	"net/http"

	"github.com/WellyShadow/Golang_Homework/go-task4/service"
	"github.com/gorilla/mux"
)

type GetUserResponse struct {
	Id string `json:"id,omitempty"`
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	id := vars["id"]
	//user := service.User{}
	user := service.GetUser(id)
	fmt.Fprintf(w, "%+v", user)
}
