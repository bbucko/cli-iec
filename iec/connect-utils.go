package main

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type MqttParameters struct {
	qos      byte
	host     string
	clientId string
	token    string
}

func (params MqttParameters) connectWSSClient() (MQTT.Client, error) {
	opts := MQTT.NewClientOptions().AddBroker(params.getWSSBroker())
	opts.SetClientID(params.clientId)

	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return c, nil
}

func (params MqttParameters) connectSSLClient() (MQTT.Client, error) {
	opts := MQTT.NewClientOptions().AddBroker(params.getSSLBroker())
	opts.SetClientID(params.clientId)
	opts.SetUsername("dummyUsername")
	opts.SetPassword(params.token)

	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return c, nil
}

func (params MqttParameters) getSSLBroker() string {
	hostName := "ssl://" + params.host + ":8883"
	return hostName
}

func (params MqttParameters) getWSSBroker() string {
	hostName := "wss://" + params.host + ":443?jwt=" + params.token
	return hostName
}

func getSubscribersJWT(keyName string) string {
	return "eyJhbGciOiJSUzI1NiJ9.eyJjbGllbnRJZCI6IldlYlNvY2tldFN1YiIsImdyb3VwcyI6IldlYlNvY2tldFN1YiIsImV4cCI6MTU0NDEwNzk2N30.Tir1a75DzikdE8gE09vARL7brmQVdPljjyTyXdbEtRDglrTAPqRqF340gSm6WZd10GL07IeLrM3Y-BgIJX1RiiMTC2NziV0kPE3Ix7Sd-n6bYaZpqeQIjCEvR6s7uJdLfoPdshBj8AKSg9HoAo_KXZjsHuRNkIie4eMU_WbRLACGl-S5dgkUWt9pb9mcBMnGt4t6g0upcInyteykd8V89NgYjv-wmkm0UwVczheFhhffodi_XycgrbB3mIlsMXkET5CZdGXgfNfh07yi8QGdryZCJaajc_II8spExu5-LaRViUHokqcKsyaAnMgJ-s_rpFfYzbWzm9TIJEVkKQTPZw"
}

func getPublishersJWT(keyName string) string {
	return "eyJhbGciOiJSUzI1NiJ9.eyJjbGllbnRJZCI6IldlYlNvY2tldFB1YiIsImdyb3VwcyI6IldlYlNvY2tldFB1YiIsImV4cCI6MTU0NDEwNTcyN30.A_u0irpMhGTp-_kvyRZyzwwIeMVJWGFbwAEtLaRtha1echGsCAhWK_S0r4dz4LwtTbzsdlLJRDzv23GosHqu2jX8U6ZLkA9ihLe7MAT2CrypiPO4N2iqlpBHpnQSxJP7sbPjyRlhlPRSW-kZZCaH7bmAii0Zx77l9lMDji6yStf1OxIKp4P2YHTY5cMALp5AbPBKmD8UTU6yl5h9d1Me9xvMrGmhVPRbc37zFbRv1b3q8135vaYPeC9_7pIQFPGZf-Q7q6Wu7e3H_IenKB65Vy-q2xWbmuqDVeXetYL1jsKNzwKZyyjHejqJEX645Xeb40qYzW4qsKf66J9eD2aSbQ"
}
