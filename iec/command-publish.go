package main

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/urfave/cli"
	"log"
)

var commandPublish = cli.Command{
	Name:        "publish",
	ArgsUsage:   "[qos][host][client-id][key-name][topic][message]",
	Description: "Publishes message to the topic",
	HideHelp:    true,
	Action:      callPublish,
	Flags: []cli.Flag{
		cli.IntFlag{
			Name:  "qos",
			Usage: "qos - Quality of service. Allowed values 0, 1, 2.",
			Value: 0,
		},
		cli.StringFlag{
			Name:  "host",
			Usage: "Url to host",
			Value: "qa4.dcp-test.com",
		},
		cli.StringFlag{
			Name:  "client-id",
			Usage: "Client id",
			Value: "WebSocketPub",
		},
		cli.StringFlag{
			Name:  "key-name",
			Usage: "Name of the generated key",
		},
		cli.StringFlag{
			Name:  "topic",
			Usage: "Name of the topic to which message will be published",
			Value: "test",
		},
		cli.StringFlag{
			Name:  "message",
			Usage: "Message to be published",
			Value: "Hello Akamai",
		},
	},
}

func callPublish(context *cli.Context) error {
	parameters := MqttParameters{}
	parameters.qos = byte(context.Int("qos"))
	parameters.host = context.String("host")
	parameters.clientId = context.String("client-id")
	parameters.token = getPublishersJWT(context.String("key-name"))

	return connectAndPublish(parameters, context.String("topic"), context.String("message"))
}

func connectAndPublish(mqttParams MqttParameters, topicName string, message string) error {

	client, error := mqttParams.connectSSLClient()
	if error != nil {
		log.Fatalln("Unable to publish message, caused by connection error:", error)
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

func publishMessage(client MQTT.Client, mqttParams MqttParameters, topicName string, message string) error {
	token := client.Publish(topicName, mqttParams.qos, false, message)
	token.Wait()
	return token.Error()
}
