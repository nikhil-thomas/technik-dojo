package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	// https://docs.aws.amazon.com/general/latest/gr/rande.html
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("<region>"),
	})

	if err != nil {
		fmt.Println("Error gettng session:")
		fmt.Println(err)
		os.Exit(1)
	}

	// create DynamoDB client and expose HTTP requests/responses
	svc := dynamodb.New(sess, aws.NewConfig().WithLogLevel(aws.LogDebugWithHTTPBody))

	// add custom header value of 10
	svc.Handlers.Send.PushFront(func(r *request.Request) {
		r.HTTPRequest.Header.Set("CustomHeader", fmt.Sprintf("%d", 10))
	})

	// call listTables just to see HTTP request/response
	svc.ListTables(&dynamodb.ListTablesInput{})
}
