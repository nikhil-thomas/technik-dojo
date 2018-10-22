package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("not enough arguments")
	}

	region := os.Getenv("AWS_REGION")
	bucket := os.Args[1]

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		log.Fatalln(err)
	}
	svc := s3.New(sess)

	_, err = svc.DeleteBucketWebsite(&s3.DeleteBucketWebsiteInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("successfully deleted website configuration to bucket:", bucket)
}
