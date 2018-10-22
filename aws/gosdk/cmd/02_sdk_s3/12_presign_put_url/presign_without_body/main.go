package main

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// If you omit the Body field, users can write any contents to the given object.
func main() {
	if len(os.Args) < 3 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION"))},
	)
	if err != nil {
		log.Fatalln(err)
	}

	bucket := os.Args[1]
	key := os.Args[2]

	svc := s3.New(sess)

	h := md5.New()
	content := strings.NewReader("")
	content.WriteTo(h)

	resp, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	md5s := base64.StdEncoding.EncodeToString(h.Sum(nil))
	resp.HTTPRequest.Header.Set("Content-MD5", md5s)

	url, err := resp.Presign(1 * time.Minute)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("\n", url, "\n")

	req, err := http.NewRequest(http.MethodPut, url, strings.NewReader(""))
	req.Header.Set("Content-MD5", md5s)

	if err != nil {
		log.Fatalln(err)
	}

	defClient, err := http.DefaultClient.Do(req)
	fmt.Println(defClient, err)
}
