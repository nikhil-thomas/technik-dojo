package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("<region>"),
	})
	if err != nil {
		log.Fatalln("error creating aws session:", err)
	}

	svc := s3.New(sess)

	inputparams := &s3.ListObjectsInput{
		Bucket:  aws.String("<bucketname>"),
		MaxKeys: aws.Int64(10),
	}

	pageNum := 0
	svc.ListObjectsPages(inputparams,
		func(page *s3.ListObjectsOutput, lastPage bool) bool {
			pageNum++
			fmt.Printf("\npage: %d\n", pageNum)
			for _, value := range page.Contents {
				fmt.Println(*value.Key)
			}
			fmt.Print("\n")
			return pageNum < 3
		},
	)

	fmt.Println("page num:", pageNum)
}
