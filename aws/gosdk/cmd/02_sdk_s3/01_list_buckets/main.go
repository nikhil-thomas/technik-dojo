package main

import (
	"fmt"
	"os"

	"github.com/nikhil-thomas/technik-dojo/aws/gosdk/internal/errhndl"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	region := os.Getenv("AWS_REGION")
	fmt.Println("region:", region)
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		errhndl.ExitErrorf("error creating session", err)
	}

	svc := s3.New(sess)
	result, err := svc.ListBuckets(nil)

	if err != nil {
		errhndl.ExitErrorf("unable to list buckets", err)
	}

	fmt.Println(result)

	fmt.Println("Buckets")
	for _, b := range result.Buckets {
		fmt.Printf("%s created on %s\n", aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}
}
