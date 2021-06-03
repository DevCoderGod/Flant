package main

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

// NewPSQL opens connection
func NewPSQL(config string) (db *PSQL, err error) {
	connect, err := sql.Open("postgres", config)
	if err != nil {
		return nil, fmt.Errorf("connect.Open error:  %v\n", err)
	}

	err = connect.Ping()
	if err != nil {
		return nil, fmt.Errorf("database.Ping error:  %v\n", err)
	}
	return &PSQL{
		connect: connect,
		config:  config,
	}, nil
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

func (db *PSQL) Close() {
	db.connect.Close()
	return
}
