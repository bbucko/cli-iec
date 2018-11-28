package main

import (
	"github.com/urfave/cli"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"log"
)

var commandPublish = cli.Command {
	Name:        "publish",
	ArgsUsage:   "[qos][namespace][jurisdiction][host][client-id][key-name][topic][message]",
	Description: "Publishes message to the topic",
	HideHelp:    true,
	Action:      callPublish,
	Flags:       []cli.Flag {
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
			Value: "HackathonPub",
        },
		cli.StringFlag {
			Name: "client-id-claim",
			Usage: "Optional client id claim name",
			Value: "clientId",
		},
		cli.StringFlag {
			Name: "auth-groups",
			Usage: "Optional authorized groups",
			Value: "WebSockePSub",
		},
		cli.StringFlag {
			Name: "auth-groups-claim",
			Usage: "Optional authorized groups claim name",
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
		cli.StringFlag {
			Name: "message",
			Usage: "Message to be published",
			Value: "Hello Akamai",
        },
    },
}

func callPublish(context *cli.Context) error {
	return connectAndPublish(buildMQTTParameters(context), context.String("topic"), context.String("message"))
}

func connectAndPublish(mqttParams MQTTParameters, topicName string, message string) error {
	client, error := mqttParams.connectSSLClient()
	if error != nil {
		log.Fatalf("Unable to publish message. Host '%s' connection error: %s\n", mqttParams.host, error)
		return error
	}

	log.Println("Publisher successfully connected to server:", mqttParams.host)

	error = publishMessage(client, mqttParams, topicName, message)
	if error != nil {
		log.Fatalln("Unable to publish message, caused by:", error)
		return error
	}

	log.Printf("Message '%s' sent to topic '%s' with qos: %d.\n", message, topicName, mqttParams.qos)

	return nil
}

func publishMessage(client MQTT.Client, mqttParams MQTTParameters, topicName string, message string) error {
	token := client.Publish(topicName, mqttParams.qos, false, message)
	token.Wait()
	return token.Error()
}