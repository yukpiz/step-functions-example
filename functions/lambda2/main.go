package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Lambda2Response struct {
	Success  bool `json:"success"`
	NextLoop bool `json:"next_loop"`
}

var i int = 0

func Handler() (Lambda2Response, error) {
	fmt.Println("START: lambda2.main.go")

	nextLoop := true
	if i == 2 {
		nextLoop = false
	}

	i += 1

	return Lambda2Response{
		Success:  true,
		NextLoop: nextLoop,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
