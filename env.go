package socket_apigateway_golang

import "os"

const (
	KEY_REGION   = "apigateway_region"
	KEY_ENDPOINT = "apigateway_endpoint"
)

var (
	REGION   string
	ENDPOINT string
)

func getAWSConfig() {
	if len(REGION) == 0 {
		REGION = os.Getenv(KEY_REGION)
		ENDPOINT = os.Getenv(KEY_ENDPOINT)
	}
}
