package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("<region>"),
	})
	if err != nil {
		log.Fatalln(err)
	}

	svc := s3.New(sess)

	svc.PutObject((&s3.PutObjectInput{}).
		SetBucket("<bucketname>").
		SetKey("myKey").
		SetBody(strings.NewReader("object body new")).
		SetWebsiteRedirectLocation("https://example.com/something"),
	)

	fmt.Println("operation completed")
}
