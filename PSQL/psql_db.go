package PSQL

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// Postgresql Database`s opened connection
type PSQL struct {
	config  string
	connect *sql.DB
}

func NewPSQL(config string) (db *PSQL, err error) {
	fmt.Println("PSQL.Check")

	connect, err := sql.Open("postgres", config)
	if err != nil {
		fmt.Printf("NewPSQL.Open err =  %v\n", err)
		return nil, err
	}

	fmt.Println("PSQL.Ping start")
	err = connect.Ping()
	if err != nil {
		fmt.Printf("PSQL.Ping err =  %v\n", err)
		return nil, err
	}
	fmt.Println("PSQL.Ping OK")

	fmt.Println("PSQL.Check OK")
	return &PSQL{
		connect: connect,
		config:  config}, nil
}

func (db *PSQL) Query(query string) (*sql.Rows, error) {

	rows, err := db.connect.Query(query)
	if err != nil {
		fmt.Printf("db.Query (%v) --- ERROR\n", query)
		log.Fatalf("Query Fatal error: %v", err)
	}

	return rows, err
}

func (db *PSQL) QueryRow(query string) *sql.Row {

	return db.connect.QueryRow(query)
}

func (db *PSQL) Exec(query string, time time.Time) bool {
	_, err := db.connect.Exec(query, time)
	if err != nil {
		fmt.Println("PSQL.connect.Exec err = ", err)
		return false
	}
	return true
}
