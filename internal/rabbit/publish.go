package rabbit

import (
	"context"
	"encoding/json"
	"fmt"
	"socketrabbit/internal/entities"
	"strconv"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (c *Client) PublishEvent(userId int, e entities.Event) {
	ctx := context.Background()
	publishData, err := json.Marshal(e)
	if err != nil {
		fmt.Printf("Error marshalling data to JSON for RabbitMQ: %v\n", err)
		return
	}

	err = c.Channel.PublishWithContext(ctx, "server.streamlabs", strconv.Itoa(userId), false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        publishData,
	})
}
