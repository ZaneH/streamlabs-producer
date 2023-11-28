package streamlabs

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"socketrabbit/internal/entities"
	"socketrabbit/internal/rabbit"

	gosocketio "github.com/ambelovsky/gosf-socketio"
	"github.com/ambelovsky/gosf-socketio/transport"
)

var (
	MarshalError = errors.New("Error marshalling data to JSON")
)

type StreamlabsClient struct {
	socket *gosocketio.Client
}

func NewClient() *StreamlabsClient {
	return &StreamlabsClient{}
}

func (*StreamlabsClient) Setup(rabbitClient *rabbit.Client) error {
	c, err := gosocketio.Dial(
		fmt.Sprintf("wss://sockets.streamlabs.com/socket.io/?EIO=3&transport=websocket&token=%s", os.Getenv("STREAMLABS_SOCKET_TOKEN")),
		transport.GetDefaultWebsocketTransport())

	if err != nil {
		log.Fatal(err)
		return err
	}

	err = c.On(gosocketio.OnConnection, func(h *gosocketio.Channel) {
		log.Println("Connected to Streamlabs")
	})

	if err != nil {
		log.Fatal(err)
		return err
	}

	err = c.On("event", func(c *gosocketio.Channel, data interface{}) {
		// Assert the data to a map
		dataMap, ok := data.(map[string]interface{})
		if !ok {
			log.Println("Error asserting data to map")
			return
		}

		// Extract event type
		eventType, ok := dataMap["type"].(string)
		if !ok {
			log.Println("Error extracting event type")
			return
		}

		event, err := handleEvent(eventType, dataMap)

		if err != nil {
			log.Printf("Error handling event: %v\n", err)
			log.Println("Raw Data:", dataMap)
			return
		}

		switch e := event.(type) {
		case *DonationEvent:
			log.Println("Received a DonationEvent:", e)
			err = rabbitClient.PublishEvent(1, e)
		case *FollowEvent:
			log.Println("Received a FollowEvent:", e)
			err = rabbitClient.PublishEvent(1, e)
		case *SuperchatEvent:
			log.Println("Received a SuperchatEvent:", e)
			err = rabbitClient.PublishEvent(1, e)
		case *SubscriptionEvent:
			log.Println("Received a SubscriptionEvent:", e)
			err = rabbitClient.PublishEvent(1, e)
		default:
			log.Println("Received an unknown event:", e)
		}

		if err != nil {
			log.Printf("Error publishing event: %v\n", err)
		}
	})

	return nil
}

func (s *StreamlabsClient) Close() {
	if s.socket != nil {
		s.socket.Close()
	}

	log.Println("Disconnected from Streamlabs")
}

func handleEvent(eventType string, dataMap map[string]interface{}) (entities.Event, error) {
	switch eventType {
	case "donation":
		return handleDonationEvent(dataMap)
	case "follow":
		return handleFollowEvent(dataMap)
	case "superchat":
		return handleSuperchatEvent(dataMap)
	case "subscription":
		return handleSubscriptionEvent(dataMap)
	default:
		log.Printf("Unknown event type: %s", eventType)
		return entities.GenericEvent{EventType: eventType}, fmt.Errorf("Unknown event type: %s", eventType)
	}
}

func handleDonationEvent(dataMap map[string]interface{}) (*DonationEvent, error) {
	// Convert dataMap to JSON
	jsonData, err := json.Marshal(dataMap)
	if err != nil {
		return nil, MarshalError
	}

	// Decode JSON data into DonationEvent
	var donationEvent DonationEvent
	if err := json.Unmarshal(jsonData, &donationEvent); err != nil {
		return nil, fmt.Errorf("Error decoding data into DonationEvent: %v", err)
	}

	return &donationEvent, nil
}

func handleSuperchatEvent(dataMap map[string]interface{}) (*SuperchatEvent, error) {
	// Convert dataMap to JSON
	jsonData, err := json.Marshal(dataMap)
	if err != nil {
		return nil, MarshalError
	}

	// Decode JSON data into SuperchatEvent
	var superchatEvent SuperchatEvent
	if err := json.Unmarshal(jsonData, &superchatEvent); err != nil {
		return nil, fmt.Errorf("Error decoding data into SuperchatEvent: %v", err)
	}

	return &superchatEvent, nil
}

func handleFollowEvent(dataMap map[string]interface{}) (*FollowEvent, error) {
	// Convert dataMap to JSON
	jsonData, err := json.Marshal(dataMap)
	if err != nil {
		return nil, MarshalError
	}

	// Decode JSON data into FollowEvent
	var followEvent FollowEvent
	if err := json.Unmarshal(jsonData, &followEvent); err != nil {
		return nil, fmt.Errorf("Error decoding data into FollowEvent: %v", err)
	}

	return &followEvent, nil
}

func handleSubscriptionEvent(dataMap map[string]interface{}) (*SubscriptionEvent, error) {
	// Convert dataMap to JSON
	jsonData, err := json.Marshal(dataMap)
	if err != nil {
		return nil, MarshalError
	}

	// Decode JSON data into SubscriptionEvent
	var subscriptionEvent SubscriptionEvent
	if err := json.Unmarshal(jsonData, &subscriptionEvent); err != nil {
		return nil, fmt.Errorf("Error decoding data into SubscriptionEvent: %v", err)
	}

	return &subscriptionEvent, nil
}
