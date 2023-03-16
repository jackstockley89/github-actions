package localstack

import (
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// ResourceClientS3 returns the S3 client
func ResourceClientS3(local bool) *s3.Client {
	if local {
		awsCfg := Localstack()
	}

	// Create the resource client
	client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	return client
}

// ResourceClientEC2 returns the EC2 client
func ResourceClientEC2(local bool) *ec2.Client {
	if local {
		awsCfg := Localstack()
	}

	// Create the resource client
	client := ec2.NewFromConfig(awsCfg)

	return client
}

// ResourceClientEKS returns the eks client
func ResourceClientEKS(local bool) *eks.Client {
	if local {
		awsCfg := Localstack()
	}

	// Create the resource client
	client := eks.NewFromConfig(awsCfg)

	return client
}
