package main

import (
	"fmt"
	"net/http"
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
	fmt.Println(ts)
	return http.ListenAndServe(ts.config.GetServerAddress(), nil)
}

func (ts *Server) rootHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Path

	switch r.Method {
	case "GET":
		if query == "/" {
			query = "/index.html"
		}
		query = "./TARO" + query
		fmt.Println("it is GET!!!", query)
		http.ServeFile(w, r, query)

	case "POST":
		fmt.Println("it is POST!!!")
		err := r.ParseMultipartForm(0)
		fmt.Println("it is POST!!!", r.Form)
		if err != nil {
			fmt.Println("r.Form error = ", err)
			break
		}

		event := Event{
			ID: r.PostFormValue("mail"),
		}

		err = ts.publisher.Publish(event)
		if err != nil {
			fmt.Printf("Publish event: %v error: %v\n", event, err)
			return
		}
	}
}
