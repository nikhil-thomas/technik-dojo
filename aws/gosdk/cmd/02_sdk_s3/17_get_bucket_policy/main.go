package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws/awserr"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("provide bucket name")
	}
	bucket := os.Args[1]

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	})
	if err != nil {
		log.Fatalln(err)
	}
	svc := s3.New(sess)

	result, err := svc.GetBucketPolicy(&s3.GetBucketPolicyInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				log.Fatalln("no such bucket,", bucket)
			case "NoSuchBucketPolicy":
				log.Fatalln("NoSuchBucketPolicy, bucket:", bucket)
			}
		}
		log.Fatalln("unable to get bucket policy, bucket:", bucket)
	}

	out := bytes.Buffer{}
	policyStr := aws.StringValue(result.Policy)
	if err := json.Indent(&out, []byte(policyStr), "", " "); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("bucket :", bucket, ": policy")
	fmt.Println(out.String())
}
