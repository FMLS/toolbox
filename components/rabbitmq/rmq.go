package rabbitmq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func NewRmq(port int, host, username, password, vhost string) (*Rmq, error) {
	rmqOption := &RmqOption{
		Host:     host,
		Port:     port,
		UserName: username,
		Password: password,
		Vhost:    vhost,
	}
	log.Print("connect to ", rmqOption.Url())
	conn, err := amqp.Dial(rmqOption.Url())
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &Rmq{
		rmqOption: rmqOption,
		Conn:      conn,
		Chan:      ch,
	}, nil
}

type RmqOption struct {
	Host     string
	Port     int
	UserName string
	Password string
	Vhost    string
}

func (option *RmqOption) Url() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/%s",
		option.UserName,
		option.Password,
		option.Host,
		option.Port,
		option.Vhost,
	)
}

type Rmq struct {
	rmqOption *RmqOption
	Conn      *amqp.Connection
	Chan      *amqp.Channel
}

func (rmq *Rmq) Close() {
	_ = rmq.Chan.Close()
	_ = rmq.Conn.Close()
}

func (rmq *Rmq) Publish(exchange, routingKey string, data []byte) error {
	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        data,
	}
	if err := rmq.Chan.Publish(exchange, routingKey, false, false, msg); err != nil {
		return err
	}
	return nil
}
