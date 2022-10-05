package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/uuid"
)

type UserRequest struct {
	Name    string `json:"name"`
	SurName string `json:"surName"`
}

type UserResponse struct {
	Id string `json:"id,omitempty"`
}

func User(w http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodPost {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "Internal Server Error : "+err.Error(), http.StatusInternalServerError)
	}
	var reqBody UserRequest

	err = json.Unmarshal(body, &reqBody)

	if err != nil {
		http.Error(w, "Internal Server Error : "+err.Error(), http.StatusInternalServerError)
	}
	id := uuid.New()
	resBodyCreate := UserResponse{id.String()}
	b, err := json.Marshal(resBodyCreate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(b)

}
