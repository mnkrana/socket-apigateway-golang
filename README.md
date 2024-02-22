# socket-apigateway-golang
api gateway for socket apis

Add environment variables in the lambda function configuration
```
KEY_REGION   = "apigateway_region"
KEY_ENDPOINT = "apigateway_endpoint"
```

- First function call should be Init() to get the evn vars and create a session

