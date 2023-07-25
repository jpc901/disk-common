package rabbitmq

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/jpc901/disk-common/conf"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Send(config *conf.RabbitMQConfig, msg []byte) error {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", config.Username, config.Password, config.Host, config.Port))
	if err != nil {
		log.Errorf("Failed to connect to RabbitMQ: %s", err)
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Errorf("Failed to open a channel: %s", err)
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		config.TransOSSQueueName, // name
		true,                     // durable
		false,                    // delete when unused
		false,                    // exclusive
		false,                    // no-wait
		nil,                      // arguments
	)
	if err != nil {
		log.Errorf("Failed to declare a queue: %s", err)
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		})
	if err != nil {
		log.Errorf("Failed to publish a message: %s", err)
		return err
	}
	log.Printf(" [x] Sent %s\n", msg)
	return nil
}

func Receive(config *conf.RabbitMQConfig, msgChan chan []byte) error {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", config.Username, config.Password, config.Host, config.Port))
	if err != nil {
		log.Errorf("Failed to connect to RabbitMQ: %s", err)
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Errorf("Failed to open a channel: %s", err)
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		config.TransOSSQueueName, // name
		true,                     // durable
		false,                    // delete when unused
		false,                    // exclusive
		false,                    // no-wait
		nil,                      // arguments
	)
	if err != nil {
		log.Errorf("Failed to declare a queue: %s", err)
		return err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Errorf("Failed to register a consumer: %s", err)
		return err
	}

	for {
		select {
		case d := <-msgs:
			if len(d.Body) == 0 {
				continue
			}
			log.Printf(fmt.Sprintf("Received a message from %s: [%s]", q.Name, string(d.Body)))
			msgChan <- d.Body
		}
	}

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	return nil
}
