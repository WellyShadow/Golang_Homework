package controller

import (
	"io"
	"net/http"
)

func Createuser(w http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodPost {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "Internal Server Error : "+err.Error(), http.StatusInternalServerError)
	}
	service.CreateUser(body, w)

}
