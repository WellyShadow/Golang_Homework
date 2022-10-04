package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

var (
	queryCreateTable = `CREATE TABLE my_table (
    id  serial PRIMARY KEY,
    name varchar not null,
    created_at TIMESTAMP with time zone,
    description varchar
)`
	queryDeleteTable = `DROP TABLE my_table`

	queryInsert = `INSERT INTO my_table (name, created_at, description) VALUES ($1, $2, $3)`

	querySelect = `SELECT id, name, created_at, description FROM my_table`
)

type record struct {
	id          int
	name        string
	createdAt   time.Time
	description sql.NullString
}

type LoginUserRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	Url string `json:"url"`
}

type CreateUserRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	Id       string `json:"id,omitempty"`
	UserName string `json:"userName,omitempty"`
}

var usersMap map[string]string

func createuser(w http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodPost {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "Internal Server Error : "+err.Error(), http.StatusInternalServerError)
	}

	var reqBody CreateUserRequest

	err = json.Unmarshal(body, &reqBody)

	if err != nil {
		http.Error(w, "Internal Server Error : "+err.Error(), http.StatusInternalServerError)
	}

	if len(reqBody.Password) >= 4 && len(reqBody.UserName) >= 8 {
		//usersMap[reqBody.Password] = reqBody.UserName
		id := uuid.New()
		resBodyCreate := CreateUserResponse{id.String(), reqBody.UserName}
		usersMap[resBodyCreate.UserName] = resBodyCreate.Id
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
	var resBodyLogin LoginUserResponse
	err = json.Unmarshal(body, &reqBodyLogin)

	if err != nil {
		http.Error(w, "Internal Server Error : "+err.Error(), http.StatusInternalServerError)
	}

	_, ok := usersMap[reqBodyLogin.UserName] //check including
	if ok {
		url := "ws://fancy-chat.io/ws&"
		resBodyLogin.Url = url + usersMap[reqBodyLogin.UserName]
		b, err := json.Marshal(resBodyLogin)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(b)
	} else {
		http.Error(w, "Invalid username/password", http.StatusBadRequest)
	}

}

func main() {

	handler := http.NewServeMux()
	handler.HandleFunc("/user", createuser)
	handler.HandleFunc("/user/login", loginuser)
	usersMap = make(map[string]string)

	connStr := "user=postgres dbname=postgres sslmode=disable password=root"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(queryCreateTable)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(queryInsert, "user", time.Now(), nil)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(querySelect)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		r := record{}
		err := rows.Scan(&r.id, &r.name, &r.createdAt, &r.description)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("record: %+v", r)
	}

	_, err = db.Exec(queryDeleteTable)
	if err != nil {
		log.Fatal(err)
	}

	s := &http.Server{
		Addr:           ":8080",
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
