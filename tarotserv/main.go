package main

import (
	"fmt"
	"log"
)

func main() {
	var cfg ConfJSON
	cfg.PSQLURI = "host=postgres port=5432 user=postgres password=secret dbname=profilerDB sslmode=disable"
	cfg.ServerAddress = ":8080"
	cfg.RmqURI = "amqp://guest:guest@rabbit:5672/xr"
	cfg.ExchangeName = "tracking"

	if err := CheckPSQL(cfg.GetPSQLURI()); err != nil {
		log.Fatal(err)
	}

	if err := CheckRMQ(cfg.GetRmqURI()); err != nil {
		log.Fatal(err)
	}

	// Prepare publisher
	publisher, err := NewRmqPublisher(cfg.GetRmqURI(), cfg.GetExchangeName())
	if err != nil {
		fmt.Println("Prepare publisher err = ", err)
	}

	// Start server
	server := newServer(cfg, publisher)
	fmt.Println("start")
	log.Fatal(server.Start())
}
