package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	queryCreateTablePhones = `CREATE TABLE phones (
		id serial PRIMARY KEY,
		userid  varchar,
		phones varchar,
		FOREIGN KEY (userid) REFERENCES accounts (userid) 
)`
	queryCreateTableAccount = `CREATE TABLE accounts (
	userid varchar PRIMARY KEY,
    name varchar not null,
    surname varchar not null
)`

	queryDeleteTable = `DROP TABLE accounts`

	queryInsertAccount = `INSERT INTO accounts (userid, name, surname) VALUES ($1, $2, $3)`
	queryInsertPhone   = `INSERT INTO phones (userid, phones) VALUES ($1, $2)`

	//queryEditSelect = `SELECT * FROM Products WHERE id = (id) VALUES ($1)`
	//queryEditInsert = `INSERT INTO my_table (name, surname, phones) VALUES ($1, $2, $3)`

	querySelect      = `SELECT userid, name, surname FROM accounts WHERE userid = ($1)`
	querySelectPhone = `SELECT phones FROM phones WHERE userid = ($1)`
)

type User struct {
	ID      string `bson:"_id" json:"id"`
	Name    string
	SurName string
	Phones  []string
}

type record struct {
	id      string
	name    string
	surname string
}
type Repository struct {
	db *sql.DB
}
type Collection struct {
	coll *mongo.Collection
}

func ConnectDBmongo() *Collection {
	connStr := "mongodb://localhost:27017"
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connStr))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	coll := client.Database("go").Collection("users")

	return &Collection{coll: coll}
}
func (coll *Collection) InputDBmongo(id, name, surname string, phones []string) {
	ctx := context.Background()
	u := &User{
		ID:      id,
		Name:    name,
		SurName: surname,
		Phones:  phones,
	}

	fmt.Println(u)
	_, err := coll.coll.InsertOne(ctx, u)
	if err != nil {
		log.Fatal(err)
	}

}

func (coll *Collection) OutputMongoDB(id string) (string, string, string, []string) {
	ctx := context.Background()
	data, err := coll.coll.Find(ctx, bson.D{{"_id", id}})
	if err != nil {
		log.Fatal(err)
	}

	//output := json.MarshalIndent(data, "")
	log.Printf("data: %+v\n", data)
	us := []*User{}
	err = data.All(ctx, &us)
	if err != nil {
		log.Fatal(err)
	}
	return us[0].ID, us[0].Name, us[0].SurName, us[0].Phones

}

func ConnectDBpostgres() *Repository {
	connStr := "user=postgres dbname=postgres sslmode=disable password=root"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	/*
		_, err = db.Exec(queryCreateTableAccount)
		if err != nil {
			log.Fatal(err)
		}

		_, err = db.Exec(queryCreateTablePhones)
		if err != nil {
			log.Fatal(err)
		}*/
	return &Repository{db: db}

}

func (rep *Repository) InputDBpostgres(id, name, surname, phone string) {
	_, err := rep.db.Exec(queryInsertAccount, id, name, surname)
	if err != nil {
		log.Fatal(err)
	}

	if phone == "" {
		return
	}
	_, err = rep.db.Exec(queryInsertPhone, id, phone)
	if err != nil {
		log.Fatal(err)
	}

}

func (rep *Repository) OutputDBpostgres(id string) (string, string, string) {
	rows, err := rep.db.Query(querySelect, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	r := record{}
	for rows.Next() {
		err := rows.Scan(&r.id, &r.name, &r.surname)
		if err != nil {
			log.Fatal(err)
		}
	}
	return r.id, r.name, r.surname
}

func (rep *Repository) OutputPhonesDBpostgres(id string) []string {
	rows, err := rep.db.Query(querySelectPhone, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var phone string
	var phones []string
	for rows.Next() {
		err := rows.Scan(&phone)
		if err != nil {
			log.Fatal(err)
		}
		phones = append(phones, phone)
	}
	return phones
}

func (rep *Repository) InputPhoneDBpostgres(id, phone string) {
	_, err := rep.db.Exec(queryInsertPhone, id, phone)
	if err != nil {
		log.Fatal(err)
	}
}

func (rep *Repository) DeleteDBpostgres() {
	_, err := rep.db.Exec(queryDeleteTable)
	if err != nil {
		log.Fatal(err)
	}
}
