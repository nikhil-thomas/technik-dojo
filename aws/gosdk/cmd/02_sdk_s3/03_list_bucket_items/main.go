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

	resp, err := svc.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(bName),
	},
	)

	if err != nil {
		errhndl.ExitErrorf("error in listing items in bucket %q: %v", bName, err)
	}

	fmt.Println(resp)

	for _, item := range resp.Contents {
		fmt.Println("Key (Name): ", *item.Key)
		fmt.Println("Last modified:", *item.LastModified)
		fmt.Println("Size:", *item.Size)
		fmt.Println("Storage class:", *item.StorageClass)
		fmt.Println("")
	}
}
