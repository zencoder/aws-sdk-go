package s3

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/protocol/restxml"
	"github.com/awslabs/aws-sdk-go/aws/signer/v4"
)

// S3 is a client for Amazon Simple Storage Service.
type S3 struct {
	*aws.Service
}

type S3Config struct {
	*aws.Config
}

// New returns a new S3 client.
func New(config *S3Config) *S3 {
	if config == nil {
		config = &S3Config{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config.Config),
		ServiceName: "s3",
		APIVersion:  "2006-03-01",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Build.PushBack(restxml.Build)
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Unmarshal.PushBack(restxml.Unmarshal)

	return &S3{service}
}
