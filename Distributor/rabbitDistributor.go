package Distributor

import (
	"log"
	"fmt"
	"github.com/streadway/amqp"
	"../filter"
	"../pgdatabase"
	"../models"
	"encoding/json"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func StartRabbit() {
	conn, err := amqp.Dial("amqp://byustudents:AlternativeFacts@rabbit-cluster-external-stage-1443209739.us-east-1.elb.amazonaws.com")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	twitter_q, err := ch.QueueDeclare(
		"byustudents-josh-twitter",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	instagram_q, err := ch.QueueDeclare(
		"byustudents-josh-instagram",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		twitter_q.Name, // queue name
		"social_activity_parser.twitter_activity.created",     // routing key
		"events", // exchange
		false,
		nil)
	failOnError(err, "Failed to bind a queue")

	err = ch.QueueBind(
		instagram_q.Name, // queue name
		"social_activity_parser.instagram_activity.created",     // routing key
		"events", // exchange
		false,
		nil)
	failOnError(err, "Failed to bind a queue")

	twitter_msgs, err := ch.Consume(
		twitter_q.Name, // queue
		"",     // consumer
		false,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	instagram_msgs, err := ch.Consume(
		instagram_q.Name, // queue
		"",     // consumer
		false,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	db := pgdatabase.NewDAO()
	twitterFilter := new(filter.Filter)
	instragramFilter := new(filter.Filter)
	twitterFilter.InitFilter("danger.txt")
	instragramFilter.InitFilter("danger.txt")
	twitterFilter.SetExceptionsFilter("exceptions.csv")
	instragramFilter.SetExceptionsFilter("exceptions.csv")
	forever := make(chan bool)

	go func() {
		for d := range twitter_msgs {
			var post models.Post
			json.Unmarshal(d.Body, &post)
			if twitterFilter.ContainsDangerWord(post) {
				db.AddPost(d.Body)
			}
		}
	}()

	go func() {
		for d := range instagram_msgs {
			var post models.Post
			json.Unmarshal(d.Body, &post)
			if instragramFilter.ContainsDangerWord(post) {
				db.AddPost(d.Body)
			}
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}