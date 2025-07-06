package initialize

import (
	"context"
	"show-calendar/config"
	"sync"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var s3Client *s3.Client
var s3Once sync.Once

func NewS3() *s3.Client {
	if s3Client == nil {
		s3Once.Do(func() {
			cfg, err := awsConfig.LoadDefaultConfig(context.TODO(), awsConfig.WithRegion(config.Region))
			if err != nil {
				logger := NewLogger()
				logger.Error("unable to load SDK config", err)
			}
			s3Client = s3.NewFromConfig(cfg)
		})
	}
	return s3Client
}
