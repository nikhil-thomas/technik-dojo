package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("<region>"),
	})

	if err != nil {
		log.Fatalln("Error in creating aws session:", err)
	}

	var wg sync.WaitGroup

	keysChan := make(chan string, 10)

	svc := s3.New(sess)
	buckets := []string{"<bucketname>", "<bucketname>"}

	for _, bucket := range buckets {
		params := &s3.ListObjectsInput{
			Bucket:  aws.String(bucket),
			MaxKeys: aws.Int64(100),
		}

		wg.Add(1)
		go func(param *s3.ListObjectsInput) {
			defer wg.Done()

			err := svc.ListObjectsPages(param,
				func(page *s3.ListObjectsOutput, last bool) bool {
					for _, object := range page.Contents {
						keysChan <- fmt.Sprintf("%s: %s", *param.Bucket, *object.Key)
					}
					return true
				},
			)

			if err != nil {
				fmt.Println("error listing", *params.Bucket, "objects:", err)
			}
		}(params)
	}

	go func() {
		wg.Wait()
		close(keysChan)
	}()

	for key := range keysChan {
		fmt.Println(key)
	}
}
