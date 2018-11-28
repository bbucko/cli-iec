package main

import (
	"fmt"
	akamai "github.com/akamai/cli-common-golang"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/urfave/cli"
	"log"
)

var commandSubscribe = cli.Command{
	Name:        "subscribe",
	ArgsUsage:   "[qos][host][client-id][key-name][topic]",
	Description: "Subscribes to a topic and waits for a messages published to this topic.",
	HideHelp:    true,
	Action:      callSubscribe,
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
			Value: "WebSocketSub",
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
	},
}

func callSubscribe(context *cli.Context) error {
	mqttParams := MqttParameters{}
	mqttParams.qos = byte(context.Int("qos"))
	mqttParams.host = context.String("host")
	mqttParams.clientId = context.String("client-id")
	mqttParams.token = getSubscribersJWT(context.String("key-name"))

	return connectAndSubscribe(mqttParams, context.String("topic"))
}

func connectAndSubscribe(mqttParams MqttParameters, topicName string) error {
	akamai.StartSpinner("Connecting to: "+mqttParams.host, "Connected")

	client, error := mqttParams.connectSSLClient()
	if error != nil {
		log.Fatalln("Unable to publish message, caused by connection error:", error)
		akamai.StopSpinnerFail()
		return error
	}

	akamai.StopSpinnerOk()

	akamai.StartSpinner("Subscribing to topic: "+topicName, "Subscribed")
	error = subscribeToTopic(client, mqttParams, topicName)
	if error != nil {
		log.Fatalf("Unable to subscribe topic '%s', caused by: %s\n", topicName, error)
		akamai.StopSpinnerFail()
		return error
	}

	akamai.StopSpinnerOk()

	waitForQuit()

	return nil
}

func subscribeToTopic(client MQTT.Client, mqttParams MqttParameters, topicName string) error {
	token := client.Subscribe(topicName, mqttParams.qos, onMessage)
	token.Wait()
	return token.Error()
}

func onMessage(client MQTT.Client, message MQTT.Message) {
	akamai.StartSpinner(fmt.Sprintf("Message '%s' received", string(message.Payload())), "Done")
	akamai.StopSpinnerOk()
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
