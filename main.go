package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
)

func CreateS3Client(region string) (*s3.S3, error) {
	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}
	svc := s3.New(sess, aws.NewConfig().WithRegion(region))
	return svc, nil
}

func ListBuckets(svc *s3.S3) ([]*s3.Bucket, error) {
	request := s3.ListBucketsInput{}
	resp, err := svc.ListBuckets(&request)
	if err != nil {
		return nil, err
	}
	return resp.Buckets, nil
}

func RenderBuckets(buckets []*s3.Bucket) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "CreationDate"})
	for _, bucket := range buckets {
		date := *bucket.CreationDate
		table.Append([]string{*bucket.Name, date.String()})
	}
	table.Render()
}

func main() {
	svc, err := CreateS3Client("eu-west-1")
	if err != nil {
		log.Fatal(err)
	}
	buckets, err := ListBuckets(svc)
	if err != nil {
		log.Fatal(err)
	}
	RenderBuckets(buckets)
}




