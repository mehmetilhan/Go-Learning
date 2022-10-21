package main

import (
	"github.com/streadway/amqp"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func exit_on_error(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	exit_on_error(err)
	defer conn.Close()

	ch, err := conn.Channel()
	exit_on_error(err)
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"example.fanout", //name
		"fanout",         //type
		true,             //durable
		false,            //auto-deleted
		false,            //internal
		false,            //no-wait
		nil,              //arguments
	)
	exit_on_error(err)

	for {
		message := "Procuder_" + strconv.Itoa(rand.Int())

		err = ch.Publish(
			"example.fanout", //exchange
			"",               //routing key
			false,            //mandatory
			false,            //immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(message),
			},
		)

		log.Printf(" [x] Sent %s", message)

		time.Sleep(3 * time.Second)
	}
}
