package utils

import (
	"context"
	"io"
	"show-calendar/config"
	"show-calendar/initialize"

	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"sync"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
)

var s3Uploader *Uploader
var s3Once sync.Once

type Uploader struct {
	s3 *s3.Client
}

func (u *Uploader) Upload(body io.Reader, ext string) (string, error) {
	uploader := manager.NewUploader(u.s3)
	uuid := uuid.New()
	fileName := uuid.String() + "." + ext
	output, err := uploader.Upload(context.Background(), &s3.PutObjectInput{
		Bucket: &config.Bucket,
		Key:    &fileName,
		Body:   body,
	})
	if err != nil {
		logger := initialize.NewLogger()
		logger.Error("File upload error", err)
		return "", err
	}
	return *output.Key, err
}

func (u *Uploader) Delete(key string) error {
	_, err := u.s3.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: &config.Bucket,
		Key:    &key,
	})
	if err != nil {
		logger := initialize.NewLogger()
		logger.Error("File delete error", err)
	}
	return err
}

func NewUploader() *Uploader {
	if s3Uploader == nil {
		s3Once.Do(func() {
			cfg, err := awsConfig.LoadDefaultConfig(context.TODO(), awsConfig.WithRegion(config.Region))
			if err != nil {
				logger := initialize.NewLogger()
				logger.Error("unable to load SDK config", err)
			}
			s3Uploader = &Uploader{s3.NewFromConfig(cfg)}
		})
	}
	return s3Uploader
}
