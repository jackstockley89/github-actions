package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	resource_client "github.com/github-actions/aws/aws-api/resource-client"
)

// getBucket returns the bucket name
func GetBucket(name, key string) (string, error) {

	client := resource_client.ResourceClientS3()

	params := &s3.GetObjectInput{
		Bucket: aws.String(name),
		Key:    aws.String(key),
	}

	bucket, err := client.GetObject(context.TODO(), params, func(o *s3.Options) {
		o.UsePathStyle = true
	})
	if err != nil {
		panic(err)
	}

	return bucket, err
}

// GetEKS returns the EKS cluster name
func GetEKS(name, key string) (string, error) {

	client := resource_client.ResourceClientEKS()

	params := &eks.DescribeClusterInput{
		Name: aws.String(name),
	}

	eks, err := client.DescribeCluster(context.TODO(), params)
	if err != nil {
		panic(err)
	}

	return eks, err
}
