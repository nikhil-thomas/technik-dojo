package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/nikhil-thomas/technik-dojo/aws/gosdk/internal/errhndl"
)

func main() {
	if len(os.Args) != 3 {
		errhndl.ExitErrorf(
			"Bucket and item names required\nusage: %s bucket_name object_name",
			filepath.Base(os.Args[0]),
		)
	}
	bucket := os.Args[1]
	objectKey := os.Args[2]
	region := os.Getenv("AWS_REGION")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		errhndl.ExitErrorf("error in creating session %s\n", err)
	}

	svc := s3.New(sess)

	fmt.Println("deleting object")
	_, err = svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(objectKey),
	})

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("waiting for delete")

	err = svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(objectKey),
	})

	fmt.Println("delete successful")
}
