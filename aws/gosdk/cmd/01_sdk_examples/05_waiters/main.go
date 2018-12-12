package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("<region>"),
	})
	if err != nil {
		log.Fatalln("error in creating aws session:", err)
	}
	ec2Client := ec2.New(sess)

	instanceIDsToStop := aws.StringSlice([]string{
		"<instance id>",
		"<instance id>",
		"<instance id>",
	})

	_, err = ec2Client.StopInstances(&ec2.StopInstancesInput{
		InstanceIds: instanceIDsToStop,
	})

	if err != nil {
		panic(err)
	}

	describeInstancesInput := &ec2.DescribeInstancesInput{
		InstanceIds: instanceIDsToStop,
	}

	if err := ec2Client.WaitUntilInstanceStopped(describeInstancesInput); err != nil {
		panic(err)
	}

	fmt.Println("instances stopped")
}
