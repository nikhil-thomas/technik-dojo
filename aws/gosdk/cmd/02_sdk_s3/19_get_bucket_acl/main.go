package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("please privide bucket name")
	}

	bucket := os.Args[1]

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := s3.New(sess)

	result, err := svc.GetBucketAcl(&s3.GetBucketAclInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Owner:", *result.Owner.ID)
	fmt.Print("\n")
	fmt.Println("Grants")

	for _, g := range result.Grants {
		if g.Grantee.DisplayName == nil {
			fmt.Println(" Grantee: EVERYONE")
		} else {
			fmt.Println(" Grantee:", *g.Grantee.DisplayName)
		}
		fmt.Println("Type:", *g.Grantee.Type)
		fmt.Println("Permission:", *g.Permission)
		fmt.Print("\n")
	}

}
