package repository

import (
	"database/sql"
	"log"
	"time"

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

func CreateBD() {
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
	/*
		_, err = db.Exec(queryDeleteTable)
		if err != nil {
			log.Fatal(err)
		}
	*/
}
