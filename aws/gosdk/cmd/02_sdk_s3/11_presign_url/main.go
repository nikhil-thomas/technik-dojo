package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("not enough arguments")
	}

	bucket := os.Args[1]
	key := os.Args[2]

	region := os.Getenv("AWS_REGION")

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		log.Fatalln(err)
	}

	svc := s3.New(sess)

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	urlStr, err := req.Presign(30 * time.Second)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("pre-signed url")
	fmt.Println(urlStr)
}
