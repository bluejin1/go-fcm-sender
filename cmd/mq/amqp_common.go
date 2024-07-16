package mq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

var MqConnection *amqp.Connection
var MqChannel *amqp.Channel

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
