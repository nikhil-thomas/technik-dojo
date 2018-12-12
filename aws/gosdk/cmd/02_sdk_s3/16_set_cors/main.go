package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func filterMethods(methods []string) []string {
	filtered := make([]string, 0, len(methods))
	for _, m := range methods {
		v := strings.ToUpper(m)
		switch v {
		case http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete:
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func main() {
	var bucket string

	flag.StringVar(&bucket, "b", "", "bucket to set CORS (required)")
	flag.Parse()

	if len(bucket) == 0 {
		log.Fatalln("specify a bucket")
	}

	methods := filterMethods(flag.Args())

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	})
	if err != nil {
		log.Fatalln(err)
	}

	svc := s3.New(sess)

	rule := s3.CORSRule{
		AllowedHeaders: aws.StringSlice([]string{"Authorization"}),
		AllowedOrigins: aws.StringSlice([]string{"*"}),
		AllowedMethods: aws.StringSlice(methods),
		MaxAgeSeconds:  aws.Int64(3000),
	}

	params := s3.PutBucketCorsInput{
		Bucket: aws.String(bucket),
		CORSConfiguration: &s3.CORSConfiguration{
			CORSRules: []*s3.CORSRule{&rule},
		},
	}

	_, err = svc.PutBucketCors(&params)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("updated bucket: %s for CORS, %v\n", bucket, methods)

}
