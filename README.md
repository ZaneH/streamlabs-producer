# Streamlabs Producer

Connects to Streamlabs API to parse incoming events and sends them to a queue.
The events are never consumed in this proof of concept.

## Usage

```bash
$ git clone https://github.com/ZaneH/streamlabs-producer.git
$ cd streamlabs-producer
$ go mod download
$ export STREAMLABS_SOCKET_TOKEN=<your token> # required

$ docker run --rm -d \
             --hostname my-rabbit \
             --name streamaze-rabbit \
             -p 15672:15672 \
             -p 5672:5672 \
             rabbitmq:3-management

$ go run cmd/socketrabbit.go
```

You can navigate to the [Streamlabs Alert Box page](https://streamlabs.com/dashboard#/alertbox)
to test the connection. You should see the events being printed to the console.

Navigate to `localhost:15672` to see the RabbitMQ management console. The default
username and password are both `guest`. You should be able to see the queue
`frontend.consumer.1` and the messages that were sent to it.