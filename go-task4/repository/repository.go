package repository

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
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

	querySelect = `SELECT userid, name, surname FROM accounts`
)

type record struct {
	id      string
	name    string
	surname string
}
type repository struct {
	db *sql.DB
}

func ConnectBD() *repository {
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
	return &repository{db: db}

}

func (rep *repository) InputBD(id, name, surname, phone string) {
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

func (rep *repository) OutputBD() {
	rows, err := rep.db.Query(querySelect)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		r := record{}
		err := rows.Scan(&r.id, &r.name, &r.surname)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("record: %v", r)
	}

}

func (rep *repository) EditBD() {
	_, err := rep.db.Exec(queryDeleteTable)
	if err != nil {
		log.Fatal(err)
	}
}

func (rep *repository) DeleteBD() {
	_, err := rep.db.Exec(queryDeleteTable)
	if err != nil {
		log.Fatal(err)
	}
}
