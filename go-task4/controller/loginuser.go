package controller

import (
	"encoding/json"
	"io"
	"net/http"
)

type LoginUserRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	Url string `json:"url"`
}

func loginuser(w http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodPost {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "Internal Server Error : "+err.Error(), http.StatusInternalServerError)
	}

	var reqBodyLogin LoginUserRequest
	//var resBodyLogin LoginUserResponse
	err = json.Unmarshal(body, &reqBodyLogin)

	if err != nil {
		http.Error(w, "Internal Server Error : "+err.Error(), http.StatusInternalServerError)
	}

	//_, ok := usersMap[reqBodyLogin.UserName] //check including
	url := "ws://fancy-chat.io/ws&"
	//resBodyLogin.Url = url + usersMap[reqBodyLogin.UserName]
	b, err := json.Marshal(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(b)

}
