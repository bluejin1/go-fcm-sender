package mq

import (
	"bytes"
	"fcm-sender/configs"
	"fcm-sender/helper/env"
	"fcm-sender/helper/zlog"
	"fcm-sender/internal/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

func CmdReceiverKiccOrder() error {

	var err error

	MqConnection, err = amqp.Dial(configs.RabbitMQServer)
	//failOnError(err, "Failed to connect to RabbitMQ")
	if err != nil {
		failOnError(err, "Failed to connect to RabbitMQ")
		return err
	}

	//fmt.Printf("MqConnection %C", MqConnection)
	defer MqConnection.Close()

	MqChannel, err = MqConnection.Channel()
	//failOnError(err, "Failed to open a channel")
	if err != nil {
		failOnError(err, "Failed to open a channel")
		return err
	}
	defer MqChannel.Close()

	queuePrefix := configs.QueuePrefixLive
	zlog.Info().Msgf("env.GetEnvEnvironment() %c", env.GetEnvEnvironment())

	if env.GetEnvEnvironment() == "development" {
		queuePrefix = configs.QueuePrefixTest
	}
	queueName := queuePrefix + configs.PosOrderQueueName
	zlog.Info().Msgf("CmdReceiverAppOrder queueName %c", queueName)

	q, err := MqChannel.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}
	failOnError(err, "Failed to declare a queue")

	msgs, err := MqChannel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return err
	}
	//failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			queErr := rabbitmq.NewKiccOrderReceiver(&d)
			if queErr != nil {
				zlog.Error().Msgf("NewKiccOrderReceiver Err %s", queErr.Error())
			}

			dotCount := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dotCount)
			time.Sleep(t * time.Second)
			log.Printf("Done")
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

	return nil
}
