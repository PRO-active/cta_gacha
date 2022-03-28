package s3

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type UploadRequest struct {
	BucketName string
	BucketKey  string
	File       io.Reader
	Config     aws.Config
}

func NewUploadRequest(bucketName, bucketKey string, file io.Reader, conf aws.Config) *UploadRequest {
	return &UploadRequest{
		BucketName: bucketName,
		BucketKey:  bucketKey,
		File:       file,
		Config:     conf,
	}
}

func Upload(ctx context.Context, request *UploadRequest) (string, error) {
	client := s3.NewFromConfig(request.Config)
	uploader := manager.NewUploader(client)

	output, err := uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(request.BucketName),
		Key:    aws.String(request.BucketKey),
		Body:   request.File,
	})
	if err != nil {
		return "", err
	}
	return output.Location, err
}
