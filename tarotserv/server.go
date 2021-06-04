package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Server represents http server.
type Server struct {
	config    Config    // !Interface
	publisher Publisher // !Interface
}

// newServer creates new Server instance.
func newServer(config Config, publisher Publisher) *Server {
	return &Server{
		config:    config,
		publisher: publisher,
	}
}

// Start runs server.
func (ts *Server) Start() error {
	http.HandleFunc("/", ts.rootHandler)
	return http.ListenAndServe(ts.config.GetServerAddress(), nil)
}

// rootHandler.
func (ts *Server) rootHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Path
	switch r.Method {
	case "GET":
		if query == "/" {
			query = "/index.html"
		}
		query = "./TARO" + query
		http.ServeFile(w, r, query)

	case "POST":
		contentType := r.Header.Get("content-type")

		if strings.Contains(contentType, "application/json") {
			var dataIn = map[string][]int{
				"ids": {},
			}
			if err := json.NewDecoder(r.Body).Decode(&dataIn); err != nil {
				fmt.Println("wrong data error: ", err)
			}
			var dataOut = getCardsFromDB(dataIn["ids"], ts.config.GetPSQLURI())

			w.Header().Set("content-type", "application/json")
			json.NewEncoder(w).Encode(dataOut)
		}
		if strings.Contains(contentType, "multipart/form-data") {
			err := r.ParseMultipartForm(0)
			if err != nil {
				fmt.Println("parse formdata error: ", err)
				break
			}

			event := Event{
				ID: r.PostFormValue("mail"),
			}

			err = ts.publisher.Publish(event)
			if err != nil {
				fmt.Printf("Publish event: %v error: %v\n", event, err)
			}
		}
	}
}

// Card represents card data.
type Card struct {
	Id         int    `json:"id"`
	Image      string `json:"image"`
	ImageAlt   string `json:"imageAlt"`
	Title      string `json:"title"`
	Subtitle   string `json:"subtitle"`
	Prediction string `json:"prediction"`
	Prop1      string `json:"prop1"`
	Prop2      string `json:"prop2"`
}

// getCardsFromDB.
func getCardsFromDB(ids []int, PSQLURI string) []Card {
	ids = selectCards(ids)

	db, err := NewPSQL(PSQLURI)
	if err != nil {
		fmt.Println("database is unavailable. error: ", err)
	}
	defer db.Close()

	var cards []Card
	for _, id := range ids {
		query := fmt.Sprintf("SELECT * FROM Cards WHERE id='%d'", id)
		row := db.QueryRow(query)
		var card Card

		err = row.Scan(&card.Id, &card.Image, &card.ImageAlt, &card.Title, &card.Subtitle, &card.Prediction, &card.Prop1, &card.Prop2)
		if err == sql.ErrNoRows {
			fmt.Printf("there is no such %d in database. error: %v", id, err)
			return nil
		}
		cards = append(cards, card)
	}
	return cards
}

// selectCards is temp func. not all cards yet.
func selectCards(numbers []int) (cards []int) {
	var desk = map[int]int{
		0: 0,
		1: 1,
		2: 4,
		3: 7,
		4: 8,
		5: 36,
		6: 69,
	}
	for _, v := range numbers {
		cards = append(cards, desk[v])
	}
	return
}
