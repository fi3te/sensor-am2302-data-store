package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/fi3te/sensor-am2302-data-store/pkg/aws"
)

func main() {
	lambda.Start(aws.HandleRequest)
}
