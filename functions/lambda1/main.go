package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Lambda1Response struct {
	Success bool `json:"success"`
}

func Handler() (Lambda1Response, error) {
	fmt.Println("START: lambda1.main.go")
	return Lambda1Response{
		Success: true,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
