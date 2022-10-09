package controller

import (
	"fmt"
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
	vars := mux.Vars(r)
	id := vars["id"]
	response := fmt.Sprintf("Product %s", id)
	fmt.Fprint(w, response)
}
