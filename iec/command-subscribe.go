package main

import (
	"github.com/urfave/cli"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"log"
	"fmt"
)

var commandSubscribe = cli.Command{
	Name:        "subscribe",
	ArgsUsage:   "[qos][namespace][jurisdiction][host][client-id][key-name][topic]",
	Description: "Subscribes to a topic and waits for a messages published to this topic.",
	HideHelp:    true,
	Action:      callSubscribe,
	Flags:       []cli.Flag{
		cli.IntFlag {
			Name: "qos",
			Usage: "qos - Quality of service. Allowed values 0, 1, 2.",
			Value: 0,
		},
		cli.StringFlag {
			Name: "namespace",
			Usage: "Namespace assigned of the property.",
			Value: "test",
		},
		cli.StringFlag {
			Name: "jurisdiction",
			Usage: "Jurisdiction assigned of the property.",
			Value: "eu",
		},
		cli.StringFlag {
			Name: "host",
			Usage: "Optional url to host. If not specified hostname will be retrieved from configuration using namespace/jurisdiction",
		},
		cli.StringFlag {
			Name: "client-id",
			Usage: "Client id",
			Value: "HackathonSub",
		},
		cli.StringFlag {
			Name: "client-id-claim",
			Usage: "Client id claim name",
			Value: "clientId",
		},
		cli.StringFlag {
			Name: "auth-groups",
			Usage: "Authorized groups",
			Value: "WebSocketSub",
		},
		cli.StringFlag {
			Name: "auth-groups-claim",
			Usage: "Authorized groups claim name",
			Value: "groups",
		},
		cli.StringFlag {
			Name: "key-name",
			Usage: "Name of the generated key",
		},
		cli.StringFlag {
			Name: "topic",
			Usage: "Name of the topic to which message will be published",
			Value: "test",
		},
	},
}

func callSubscribe(context *cli.Context) error {
	return connectAndSubscribe(buildMQTTParameters(context), context.String("topic"))
}

func connectAndSubscribe(mqttParams MQTTParameters, topicName string) error {
	client, error := mqttParams.connectSSLClient()
	if error != nil {
		log.Fatalf("Unable to subscribe to topic. Host '%s' connection error: %s\n", mqttParams.host, error)
		return error
	}

	log.Println("Subscriber successfully connected to server:", mqttParams.host)

	error = subscribeToTopic(client, mqttParams, topicName)
	if error != nil {
		log.Fatalf("Unable to subscribe topic '%s', caused by: %s\n", topicName, error)
		return error
	}

	log.Printf("Successfully subscribed to a topic: '%s'.\n", topicName)

	waitForQuit()

	return nil
}

func subscribeToTopic(client MQTT.Client, mqttParams MQTTParameters, topicName string) error {
	token := client.Subscribe(topicName, mqttParams.qos, onMessage)
	token.Wait()
	return token.Error()
}

func onMessage(client MQTT.Client, message MQTT.Message) {
	log.Printf("Message '%s' received.\n", string(message.Payload()))
}

func waitForQuit() {
	fmt.Println("Type q to quit the subscriber")
	for {
		var input string
		fmt.Scanln(&input)
		if input == "q" {
			break
		}
	}
	fmt.Println("Bye!")
}



