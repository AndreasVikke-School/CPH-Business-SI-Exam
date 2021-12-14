package main

import (
	"encoding/json"
	"fmt"

	eh "github.com/andreasvikke-school/CPH-Bussiness-SI-Exam/applications/services/api/errorhandler"
	"github.com/streadway/amqp"
)

func ProduceLoanEntryToRabbitmq(loanEntry LoanEntry) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s/", configuration.Rabbitmq.Username, configuration.Rabbitmq.Password, configuration.Rabbitmq.Service))
	eh.PanicOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	eh.PanicOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"LoanQueue", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	eh.PanicOnError(err, "Failed to declare a queue")

	body, err := json.Marshal(loanEntry)
	eh.PanicOnError(err, "Can't convert to JSON")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/json",
			Body:        []byte(body),
		})
	eh.PanicOnError(err, "Failed to publish a message")

}
