package rabbitMQ

import (
	"log"

	"github.com/Vainsberg/rabbitMQ-example/service"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RepositoryRabbitMQ struct {
	ch   *amqp.Channel
	conn *amqp.Connection
}

func NewRepositoryRabbitMQ(ch *amqp.Channel, conn *amqp.Connection) *RepositoryRabbitMQ {
	return &RepositoryRabbitMQ{ch: ch, conn: conn}
}

func (rab *RepositoryRabbitMQ) ConsumeMessages(queueName string) {
	msgs, err := rab.ch.Consume(
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
			service.ProcessMessage(d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func (rab *RepositoryRabbitMQ) PublishMessage(queueName string, body string) error {
	_, err := rab.ch.QueueDeclare(
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

	err = rab.ch.Publish(
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