package socket_apigateway_golang

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
)

var awsSession *session.Session
var apiGateway *apigatewaymanagementapi.ApiGatewayManagementApi

func getSession() *session.Session {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(REGION),
	})
	if err != nil {
		log.Fatalln("Unable to create AWS session", err.Error())
	}

	return sess
}

func post(input *apigatewaymanagementapi.PostToConnectionInput) {

	if awsSession == nil {
		awsSession = getSession()
	}

	if apiGateway == nil {
		apiGateway = apigatewaymanagementapi.New(awsSession, aws.NewConfig().WithEndpoint(ENDPOINT))
	}

	_, err := apiGateway.PostToConnection(input)
	if err != nil {
		log.Fatalln("Error posting to connection", err.Error())
	} else {
		log.Println("Successfully posted to connection")
	}
}
