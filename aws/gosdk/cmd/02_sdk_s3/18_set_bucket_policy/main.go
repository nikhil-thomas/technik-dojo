package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
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

	readOnlyAnonUserPolicy := map[string]interface{}{
		"Version": "2012-10-17",
		"Statement": []map[string]interface{}{
			{
				"Sid":       "AddPerm",
				"Effect":    "Allow",
				"Principal": "*",
				"Action": []string{
					"s3:GetObject",
				},
				"Resource": []string{
					fmt.Sprintf("arn:aws:s3:::%s/test/*", bucket),
				},
			},
		},
	}

	policy, err := json.Marshal(readOnlyAnonUserPolicy)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = svc.PutBucketPolicy(&s3.PutBucketPolicyInput{
		Bucket: aws.String(bucket),
		Policy: aws.String(string(policy)),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == s3.ErrCodeNoSuchBucket {
			log.Fatalln("bucket:", bucket, "does not exist")
		}
		log.Fatalln(err)
	}
	fmt.Println("successfully set bucket policy for bucket:", bucket)
}
