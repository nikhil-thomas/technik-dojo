package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("not enough arguments")
	}

	bucket := os.Args[1]
	region := os.Getenv("AWS_REGION")

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	if err != nil {
		log.Fatalln(err)
	}

	svc := s3.New(sess)

	result, err := svc.GetBucketWebsite(&s3.GetBucketWebsiteInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok && awsErr.Code() == "NoSuchWebsiteConfiguration" {
			log.Fatalln(bucket, "doesnot have website configuration")
		}
		log.Fatalln(err)
	}

	fmt.Println(bucket, ": websute configuration")
	fmt.Println(result)
}
