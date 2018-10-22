package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION"))},
	)
	if err != nil {
		log.Fatalln(err)
	}

	bucket := os.Args[1]
	key := os.Args[2]

	svc := s3.New(sess)

	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   strings.NewReader("EXPECTED CONTENTS"),
	})

	if err != nil {
		log.Fatalln(err)
	}

	urlStr, err := req.Presign(1 * time.Minute)
	fmt.Println("url")
	fmt.Println(urlStr)
}
