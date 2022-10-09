package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type GetUserResponse struct {
	Id string `json:"id,omitempty"`
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	/*
		if r.Method != http.MethodGet {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}*/
	params, found := mux.Vars(r)["id"]
	if !found {
		log.Println(": [INFO] Id not found ")
	}
	fmt.Println(`id := `, params)
}
