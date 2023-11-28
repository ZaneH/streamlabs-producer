package rabbit

import (
	"context"
	"encoding/json"
	"log"
	"socketrabbit/internal/entity"
	"strconv"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (c *Client) PublishEvent(userId int, e entity.Event) error {
	ctx := context.Background()
	publishData, err := json.Marshal(e)
	if err != nil {
		log.Printf("Error marshalling data to JSON for RabbitMQ: %v\n", err)
		return err
	}

	return c.Channel.PublishWithContext(ctx, "server.streamlabs", strconv.Itoa(userId), false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        publishData,
	})
}
