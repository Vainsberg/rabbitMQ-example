package main

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ConsumeMessages(ch *amqp.Channel, queueName string) {
	msgs, err := ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Printf("err rab.ch.Consume ", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			fmt.Println(string(d.Body))
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func PublishMessageсh(ch *amqp.Channel, queueName string, body string) error {
	_, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
		return err
	}

	err = ch.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {
		log.Fatalf("Failed to publish a message: %s", err)
		return err
	}

	log.Printf(" [x] Sent %s", body)
	return nil
}

func ConnectToRabbitMQ() (*amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
		return nil, err
	}

	log.Println("Successfully connected to RabbitMQ")
	return conn, nil
}

func main() {
	conn, err :=
		ConnectToRabbitMQ()
	if err != nil {
		panic("Error create connection Rabbit MQ")
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("failed to open channel. Error: %s", err)
	}

	PublishMessageсh(ch, "text", "hello")

	ConsumeMessages(ch, "text")

}
