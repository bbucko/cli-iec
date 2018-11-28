package main

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/urfave/cli"
)

type MQTTParameters struct {
	qos byte
	host string
	clientId string
	clientIdClaim string
	authGroups string
	authGroupsClaim string
	token string
	namespace string
	jurisdiction string
}

func (params MQTTParameters)connectWSSClient() (MQTT.Client, error) {
	opts := MQTT.NewClientOptions().AddBroker(params.getWSSBroker())
	opts.SetClientID(params.clientId)

	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return c, nil
}

func (params MQTTParameters)connectSSLClient() (MQTT.Client, error) {
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

func buildMQTTParameters(context *cli.Context) MQTTParameters {
	parameters := MQTTParameters{}
	parameters.qos = byte(context.Int("qos"))
	parameters.host = context.String("host")
	parameters.clientId = context.String("client-id")
	parameters.clientIdClaim = context.String("client-id-claim")
	parameters.authGroups = context.String("auth-groups")
	parameters.authGroupsClaim = context.String("auth-groups-claim")
	parameters.namespace = context.String("namespace")
	parameters.jurisdiction = context.String("jurisdiction")

	parameters.resolveAuthGroups()
	parameters.resolveHost()
	parameters.resolveJWT()

	return parameters
}

func (params MQTTParameters) getSSLBroker() (string) {
	hostName :=  "ssl://" + params.host + ":8883"
	return hostName
}

func (params MQTTParameters) getWSSBroker() (string) {
	hostName :=  "wss://" + params.host + ":443?jwt=" + params.token
	return hostName
}

func (params *MQTTParameters) resolveAuthGroups() {
	if len(params.authGroups) == 0 {
		params.authGroups = params.clientId
	}
}

func (params *MQTTParameters) resolveHost()  {
	if len(params.host) == 0 {
		switch params.namespace {
		case "hackathon_dcp_test_com":
			params.host = "qa4.dcp-test.com"
		case "wiktor2_dcp_test_com":
			params.host = "wbaranow2.dcp-test.com"
		default:
			params.host = "qa4.dcp-test.com"
		}
	}
}

func (params *MQTTParameters) resolveJWT() {
	params.token = GenerateTokenLocal(params.buildJWTParameters())
}

type JWTParamsLocal struct {
	namespace string
	jurisdiction string
	clientId string
	clientIdClaim string
	authGroups string
	authGroupsClaim string
}

func (mqttParams MQTTParameters) buildJWTParameters() JWTParamsLocal {
	jwtParams := JWTParamsLocal{}

	jwtParams.namespace = mqttParams.namespace
	jwtParams.jurisdiction = mqttParams.jurisdiction
	jwtParams.clientId = mqttParams.clientId
	jwtParams.clientIdClaim = mqttParams.clientIdClaim
	jwtParams.authGroups = mqttParams.authGroups
	jwtParams.authGroupsClaim = mqttParams.authGroupsClaim

	return jwtParams
}

func GenerateTokenLocal(params JWTParamsLocal) (string) {
	switch params.clientId + params.namespace {
	case "HackathonPub" + "hackathon_dcp_test_com":
		return "eyJhbGciOiJSUzI1NiJ9.eyJjbGllbnRJZCI6IkhhY2thdGhvblB1YiIsImdyb3VwcyI6IldlYlNvY2tldFB1YiIsImV4cCI6MTU0NjAxNzA1M30.1OqlwaF-4E7AuejIHrPKnmwjZ11G5qbaE8Rv7GZz1tp3IkFvxgF3tEkovNlnDJrCyXseUWUb7Ep8TAFYF34SZ8IlA7D1t78p8lhd7zN5ZGH73-6MYHzB_Vr_zg4pixmPS4uPlDrJxgPDzIdjIlAcbGsng8eVIpVOpibZ4BL05S2pdn5X2AxS3s3BB_niXWii13GQMvO0i_6dkOC2rjdmI14IK2i0zA9kEYkmxE_qFqi_UxKc8ns6CLWoq03O5gwlr-faJE6m-U3oaxgWseh6JbPowV7QKHIEg-odFRUgj0n9UMXnamA86h1MH1ldXClP_Q5MFgOp5hv0oeH5bZ0uZA"
	case "HackathonSub"  + "hackathon_dcp_test_com":
		return "eyJhbGciOiJSUzI1NiJ9.eyJjbGllbnRJZCI6IkhhY2thdGhvblN1YiIsImdyb3VwcyI6IldlYlNvY2tldFN1YiIsImV4cCI6MTU0NjAxNzA1M30.VGeSMWLZtyaAxuZojgeezLCVHmNjltihTFHBz99dQkW7vk4_PeQXamtEUtEVwsEG6zUelxa87WGibNqdOYIpxLhaU6mFAMnI0vuWERMEg6Ue5PmkaC1oV3lAaMgkWLimRt5fr5jECBTX6iVtagX5eoN_9MEjHWwLQjyuy4i3iRAb4SfN5IrZAgqxzx33Rr8c2vkfmJiTzmmU5lJLJlDM1cH3JG7FoYRY9bM3LH7Iis9ypyMvXj50WGL0NCrPiwSGgqcyTSJKYE4GaflAGGN7jx6XAlMIX4hBtFpUbfJLSn_1_Y-l0e0KfqTWawt_EBtgSXcZwPxeaUsV3-9Fh-IGMw"
	case "HackathonPub" + "wiktor2_dcp_test_com":
		return "eyJhbGciOiJSUzI1NiJ9.eyJjbGllbnRJZCI6IkhhY2thdGhvblB1YiIsImdyb3VwcyI6IldlYlNvY2tldFB1YiIsImV4cCI6MTU0NjAxMjAzOH0.Rt394Cfw4wUxKPyFyy-THm2iErzmkJo2s9cVd-QrwFI9c97BGBpOLHE9ZDfaponVDYbVmxflUtBcuFHVezV2hU6X5eQcUm4HifW_L4yl7Ubmgp_V0MtHC_K_cQQBcNjU4oxWwgkIQjk49f2vEIR60266HpbUysvZaapNRnlFoS2bROBs_ufHSC8Gqq7j2dH1D5lTB1EZ_zeIb19Qkct0cs9RoxRjc46-ZV7utEusLtn77prsOhj7uRXFzCaoIhWUpS7NtYnkCexXlc37_3xbng71u2veUxrXWwhbQnsemKXXlmCQm4GiwDZ5LH3jTR7uEhWAxBwAUo8Vf97F0NhVLg"
	case "HackathonSub" + "wiktor2_dcp_test_com":
		return "eyJhbGciOiJSUzI1NiJ9.eyJjbGllbnRJZCI6IkhhY2thdGhvblN1YiIsImdyb3VwcyI6IldlYlNvY2tldFN1YiIsImV4cCI6MTU0NjAxMjAzOH0.xzCvOegPWC4TVCMjZe7rvd8jqmKPdNUej12n2o0G1LhxaFJoStlPhahgG7wThCJ61FYzCE-5_xjYHzLkm2bgrbFDf8WsJfgiTTYWQK7eA-6Evuz6gYTTMPcMsAJlJBY3EqbmcI1m7HQSIfIDwC1C5MeY7srZu4xEEWX665dVjJE70aGlQnnWs0vu9NOsbzIhw57weDHZ0_ghy9VzGgWQuW7pfK0jnHq3cfGxJUJ4LNmLpQJFg6BuNsdfkSxJkMUAnZLDsySIcJEt-nEBdnO39U1m8pNtndtuT11mYr0yDykxryeZGFMT_UEYJ2LCwLDIMTSC20-jTg-c4cTWBH-LXg"
	default:
		return ""
	}
}