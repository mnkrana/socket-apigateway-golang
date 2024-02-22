package socket_apigateway_golang

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
)

func Init() {
	getAWSConfig()
}

func PostToConnection(connectionId string, data []byte) {

	input := &apigatewaymanagementapi.PostToConnectionInput{
		ConnectionId: aws.String(connectionId),
		Data:         data,
	}

	post(input)
}
