package main

import (
	"log"

	"github.com/streadway/amqp"
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

	q, err := ch.QueueDeclare(
		"",    //name
		false, //durable
		false, //delete when usused
		true,  //exclusive
		false, //no-wait
		nil,   //arguments
	)
	exit_on_error(err)

	err = ch.QueueBind(
		q.Name,           //queue name
		"",               //routing key
		"example.fanout", //exchange
		false,
		nil,
	)
	exit_on_error(err)

	msgs, err := ch.Consume(
		q.Name, //queue
		"",     //consumer
		true,   //auto-ack
		false,  //exclusive
		false,  //no-local
		false,  //no-wait
		nil,    //args
	)
	exit_on_error(err)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" Consumer 2 [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
