package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/service/s3"
)

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func main() {
	if len(os.Args) < 4 {
		exitErrorf("usage: %s <region> <bucket> <key>\n", filepath.Base(os.Args[0]))
	}

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(os.Args[1]),
	}))

	svc := s3.New(sess)

	resp, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(os.Args[2]),
		Key:    aws.String(os.Args[3]),
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				exitErrorf("bucket %s doesnot exist\n", os.Args[2])
			case s3.ErrCodeNoSuchKey:
				exitErrorf("key %s doesnot exist\n", os.Args[3])
			}
		}
		exitErrorf("unknown error occured:%s\n", err)
	}

	defer resp.Body.Close()

	fmt.Printf("s3://%s/%s exists. size: %d\n", os.Args[2], os.Args[3],
		aws.Int64Value(resp.ContentLength),
	)
}
