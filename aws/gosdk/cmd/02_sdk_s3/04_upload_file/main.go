package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nikhil-thomas/technik-dojo/aws/gosdk/internal/errhndl"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("specify a bucket name and a file name")
	}

	bName := os.Args[1]
	fName := os.Args[2]
	region := os.Getenv("AWS_REGION")

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.Open(fName)
	if err != nil {
		errhndl.ExitErrorf("error opening file %q: %v", fName, err)
	}
	defer file.Close()

	uploader := s3manager.NewUploader(sess)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bName),
		Key:    aws.String(fName),
		Body:   file,
	},
	)
	if err != nil {
		errhndl.ExitErrorf("error uploading file %v to bucket %v: %v", fName, bName, err)
	}

	fmt.Printf("successfully uploaded file %v to bucket %v\n", fName, bName)
}
