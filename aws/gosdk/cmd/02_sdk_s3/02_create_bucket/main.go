package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nikhil-thomas/technik-dojo/aws/gosdk/internal/errhndl"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("specify a bucket name")
	}

	bName := os.Args[1]
	region := os.Getenv("AWS_REGION")

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	if err != nil {
		log.Fatalln(err)
	}

	svc := s3.New(sess)

	_, err = svc.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bName),
	})
	if err != nil {
		errhndl.ExitErrorf("unable to create bucket %q: %v", bName, err)
	}

	fmt.Printf("waiting for bucket %q to be created...\n", bName)

	err = svc.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: aws.String(bName),
	})

	if err != nil {
		errhndl.ExitErrorf("error while waiting for bucket creation: %v", err)
	}

	fmt.Printf("bucket %q successfully created\n", bName)
}
