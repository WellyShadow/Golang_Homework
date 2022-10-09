package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/WellyShadow/Golang_Homework/go-task4/service"
)

type PhoneUserRequest struct {
	UserId string `json:"userId"`
	Phone  string `json:"phone"`
}

func Phoneuser(w http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodPost {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "Internal Server Error : "+err.Error(), http.StatusInternalServerError)
	}

	var reqBody PhoneUserRequest

	err = json.Unmarshal(body, &reqBody)

	service.AddPhone(reqBody.UserId, reqBody.Phone)

	if err != nil {
		http.Error(w, "Internal Server Error : "+err.Error(), http.StatusInternalServerError)
	}

}
