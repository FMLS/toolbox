package rabbitmq

import (
	"log"

	"github.com/urfave/cli/v2"
)

func Publish(ctx *cli.Context) error {
	rmq, err := NewRmq(
		ctx.Int("port"),
		ctx.String("host"),
		ctx.String("username"),
		ctx.String("password"),
		ctx.String("vhost"),
	)
	if err != nil {
		log.Fatal(err)
	}

	err = rmq.Publish(ctx.String("exchange"), ctx.String("key"), []byte(ctx.String("data")))
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func Consume(ctx *cli.Context) error {
	rmq, err := NewRmq(
		ctx.Int("port"),
		ctx.String("host"),
		ctx.String("username"),
		ctx.String("password"),
		ctx.String("vhost"),
	)
	if err != nil {
		log.Fatal(err)
	}

	q, err := rmq.Chan.QueueDeclare(
		ctx.String("queue"), // name
		true,                // durable
		false,               // delete when unused
		false,               // exclusive
		false,               // no-wait
		nil,                 // arguments
	)
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := rmq.Chan.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("waitting message")
	for msg := range msgs {
		log.Println(string(msg.Body))
	}
	return nil
}
