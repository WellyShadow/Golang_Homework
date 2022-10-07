package repository

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var (
	queryCreateTable = `CREATE TABLE my_table (
    id  serial PRIMARY KEY,
    name varchar not null,
    surname varchar not null,
	phones varchar

)`
	queryDeleteTable = `DROP TABLE my_table`

	queryInsert = `INSERT INTO my_table (name, surname) VALUES ($1, $2)`

	querySelect = `SELECT id, name, surname, phones FROM my_table`
)

type record struct {
	id      int
	name    string
	surname string
	phones  sql.NullString
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
	/*_, err = db.Exec(queryCreateTable)
	if err != nil {
		log.Fatal(err)
	}*/
	return &repository{db: db}
}

func (rep *repository) InputBD(name, surname string) {
	_, err := rep.db.Exec(queryInsert, name, surname)
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
		err := rows.Scan(&r.id, &r.name, &r.surname, &r.phones)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("record: %v", r)
	}
}
func (rep *repository) DeleteBD() {
	_, err := rep.db.Exec(queryDeleteTable)
	if err != nil {
		log.Fatal(err)
	}
}
