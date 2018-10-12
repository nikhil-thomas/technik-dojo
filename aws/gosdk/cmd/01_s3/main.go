package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	bucket := "<bucket_name>"
	tagName1 := "Cost Center"
	tagValue1 := "123456"
	tagName2 := "Stack"
	tagValue2 := "MyTestStack"

	// https://docs.aws.amazon.com/general/latest/gr/rande.html
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("<region>"),
	})
	if err != nil {
		log.Fatalln(err)
	}
	svc := s3.New(sess)

	input := &s3.PutBucketTaggingInput{
		Bucket: aws.String(bucket),
		Tagging: &s3.Tagging{
			TagSet: []*s3.Tag{
				{
					Key:   aws.String(tagName1),
					Value: aws.String(tagValue1),
				},
				{
					Key:   aws.String(tagName2),
					Value: aws.String(tagValue2),
				},
			},
		},
	}

	_, err = svc.PutBucketTagging(input)
	if err != nil {
		log.Fatalln(err)
	}

	inputBucketTagging := &s3.GetBucketTaggingInput{
		Bucket: aws.String(bucket),
	}

	result, err := svc.GetBucketTagging(inputBucketTagging)
	if err != nil {
		log.Fatalln(err)
	}
	numTags := len(result.TagSet)

	if numTags > 0 {
		fmt.Printf("Found %d tag(s)\n\n", numTags)

		for _, t := range result.TagSet {
			fmt.Printf("Key: %v, Value: %v\n", *t.Key, *t.Value)
		}

	} else {
		fmt.Println("did not find any tags")
	}
}
