package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

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

	downloader := s3manager.NewDownloader(sess)

	file, err := os.Create(objectKey)
	if err != nil {
		errhndl.ExitErrorf("error in creating file %s\n", err)
	}
	defer file.Close()

	numBytes, err := downloader.Download(
		file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(objectKey),
		},
	)
	if err != nil {
		errhndl.ExitErrorf("error in downloading file %s\n", err)
	}

	fmt.Println("download successfull", file.Name(), "Size:", numBytes)
}
