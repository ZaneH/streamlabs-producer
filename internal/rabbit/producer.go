package rabbit

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Client struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func (c *Client) Setup() error {
	// Connect to RabbitMQ server
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}

	// Open a channel
	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	c.Connection = conn
	c.Channel = ch
	return nil
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) NewExchange(exchangeName string) error {
	return c.Channel.ExchangeDeclare(
		exchangeName, // name
		"direct",     // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
}

func (c *Client) NewQueue(queueName string) (amqp.Queue, error) {
	return c.Channel.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
}

func (c *Client) Close() {
	// Close the channel
	c.Channel.Close()

	// Close the connection
	c.Connection.Close()
}
