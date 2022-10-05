package service

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type CreateUserRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	Id       string `json:"id,omitempty"`
	UserName string `json:"userName,omitempty"`
}

func CreateUser(body []byte, w http.ResponseWriter) {

	var reqBody CreateUserRequest

	err := json.Unmarshal(body, &reqBody)

	if err != nil {
		http.Error(w, "Internal Server Error : "+err.Error(), http.StatusInternalServerError)
	}

	if len(reqBody.Password) >= 4 && len(reqBody.UserName) >= 8 {
		id := uuid.New()
		resBodyCreate := CreateUserResponse{id.String(), reqBody.UserName}
		b, err := json.Marshal(resBodyCreate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(b)
	} else {
		http.Error(w, "Minimal length for password - 4 and for username - 8", http.StatusBadRequest)
		return
	}
}
