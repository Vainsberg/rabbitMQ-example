package main

import (
	"log"
	"net/http"

	"github.com/Vainsberg/rabbitMQ-example/handler"
	"github.com/Vainsberg/rabbitMQ-example/rabbitMQ"
)

func main() {
	conn, err :=
		rabbitMQ.ConnectToRabbitMQ()
	if err != nil {
		panic("Error create connection Rabbit MQ")
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("failed to open channel. Error: %s", err)
	}

	repositoryRabbitMQ := rabbitMQ.NewRepositoryRabbitMQ(ch, conn)

	handler := handler.NewHandler(repositoryRabbitMQ)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler.ExampleHanlder(w, r)
	})

	http.ListenAndServe(":8080", nil)
}
