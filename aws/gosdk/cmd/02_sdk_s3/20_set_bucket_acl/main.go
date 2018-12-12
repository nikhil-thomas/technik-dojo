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
	if len(os.Args) < 3 {
		log.Fatalln("please provide bucket name, email and permission")
	}

	bucket := os.Args[1]
	email := os.Args[2]

	permission := "READ"

	if len(os.Args) == 4 {
		permission = os.Args[3]
		if !(permission == "FULL_CONTROL" || permission == "WRITE" || permission == "WRITE_ACP" || permission == "READ" || permission == "READ_ACP") {
			fmt.Println("Illegal permission value")
			fmt.Println("Expected: FULL_CONTROL, WRITE, WRITE_ACP, READ, or READ_ACP")
			os.Exit(1)
		}
	}

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

	ownerID := *result.Owner.ID

	grants := result.Grants
	userType := "AmazonCustomerByEmail"
	var newGrantee = s3.Grantee{EmailAddress: &email, Type: &userType}
	var newGrant = s3.Grant{Grantee: &newGrantee, Permission: &permission}

	grants = append(grants, &newGrant)

	params := &s3.PutBucketAclInput{
		Bucket: aws.String(bucket),
		AccessControlPolicy: &s3.AccessControlPolicy{
			Grants: grants,
			Owner: &s3.Owner{
				ID: aws.String(ownerID),
			},
		},
	}

	_, err = svc.PutBucketAcl(params)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("ACL updated successfulle")
}
