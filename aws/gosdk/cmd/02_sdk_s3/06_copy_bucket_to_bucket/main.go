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
	if len(os.Args) != 4 {
		errhndl.ExitErrorf(
			"bucket, key, and other bucket names required\nUsage: go run %s_object bucket item other-bucket",
			filepath.Base(os.Args[0]),
		)
	}

	bucketSrc := os.Args[1]
	objectName := os.Args[2]
	bucketDst := os.Args[3]

	copySrc := bucketSrc + "/" + objectName

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	})
	if err != nil {
		log.Fatalln(err)
	}

	svc := s3.New(sess)

	fmt.Println("copy object")

	_, err = svc.CopyObject(&s3.CopyObjectInput{
		Bucket:     aws.String(bucketDst),
		CopySource: aws.String(copySrc),
		Key:        aws.String(objectName),
	})

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("wait until object exists")

	err = svc.WaitUntilObjectExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucketDst),
		Key:    aws.String(objectName),
	})

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(
		objectName,
		"successfully copied from",
		bucketSrc,
		"to",
		bucketDst,
	)
}
