package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/nikhil-thomas/technik-dojo/aws/gosdk/internal/errhndl"
)

func main() {
	if len(os.Args) != 2 {
		errhndl.ExitErrorf(
			"Bucket and item names required\nusage: %s bucket_name",
			filepath.Base(os.Args[0]),
		)
	}
	bucket := os.Args[1]
	region := os.Getenv("AWS_REGION")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		errhndl.ExitErrorf("error in creating session %s\n", err)
	}

	svc := s3.New(sess)

	iter := s3manager.NewDeleteListIterator(svc, &s3.ListObjectsInput{
		Bucket: aws.String(bucket),
	})

	err = s3manager.NewBatchDeleteWithClient(svc).Delete(aws.BackgroundContext(), iter)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("deleted all objects from", bucket)
}
