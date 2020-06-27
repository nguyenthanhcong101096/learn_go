package main

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, message string) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://user:password@localhost:5672")

	defer conn.Close()

	failOnError(err, "Failed to connect to RabbitMQ ")

	ch, err := conn.Channel()

	defer ch.Close()

	failOnError(err, "Failed to Open channel ")

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	body := "Hello World!"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	failOnError(err, "Failed to publish a message")
}
