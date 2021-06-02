package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("start")

	var cfg ConfJSON
	cfg.PSQLURI = "host=postgres port=5432 user=postgres password=secret dbname=profilerDB sslmode=disable"
	cfg.ServerAddress = ":8080"
	cfg.RmqURI = "amqp://guest:guest@rabbit:5672/xr"
	cfg.QueueName = "tracking"

	// Prepare publisher
	publisher, err := NewRmqPublisher(cfg.RmqURI, cfg.QueueName)
	if err != nil {
		fmt.Println("Prepare publisher err = ", err)
	}

	// Start server
	server := newServer(cfg, publisher)
	log.Fatal(server.Start())
	fmt.Println("end")
}
