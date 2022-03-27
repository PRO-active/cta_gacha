package s3

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type UploadRequest struct {
	BucketName string
	BucketKey  string
	File       io.Reader
}

func NewUploadRequest(bucketName, bucketKey string, file io.Reader) *UploadRequest {
	return &UploadRequest{
		BucketName: bucketName,
		BucketKey:  bucketKey,
		File:       file,
	}
}

func Upload(ctx context.Context, request *UploadRequest) (string, error) {
	conf, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return "", err
	}
	client := s3.NewFromConfig(conf)
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
