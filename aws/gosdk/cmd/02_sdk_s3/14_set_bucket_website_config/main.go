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
	if len(os.Args) != 4 {
		log.Fatalln("not enough arguments")
	}

	region := os.Getenv("AWS_REGION")
	bucket := os.Args[1]
	indexSuffix := os.Args[2]
	errorPage := os.Args[3]

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		log.Fatalln(err)
	}
	svc := s3.New(sess)

	params := s3.PutBucketWebsiteInput{
		Bucket: aws.String(bucket),
		WebsiteConfiguration: &s3.WebsiteConfiguration{
			IndexDocument: &s3.IndexDocument{
				Suffix: aws.String(indexSuffix),
			},
		},
	}

	if len(errorPage) > 0 {
		params.WebsiteConfiguration.ErrorDocument = &s3.ErrorDocument{
			Key: aws.String(errorPage),
		}
	}

	_, err = svc.PutBucketWebsite(&params)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("successfully updated website configuration to bucket:", bucket)
}
