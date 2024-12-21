package user

import (
	"context"
	"derma/detect/config"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type UserService struct {
	ctx    context.Context
	bucket *oss.Bucket
}

// NewUserService new UserService
func NewUserService(ctx context.Context) *UserService {
	// bucket init
	if config.OSS == nil {
		return &UserService{ctx: ctx, bucket: nil}
	}

	client, err := oss.New("https://"+config.OSS.BucketName+"."+config.OSS.Endpoint, config.OSS.AccessKeyID, config.OSS.AccessKeySecret, oss.UseCname(true))
	if err != nil {
		hlog.Fatal(err)
	}

	bucket, err := client.Bucket(config.OSS.BucketName)
	if err != nil {
		hlog.Fatal(err)
	}
	return &UserService{ctx: ctx, bucket: bucket}
}
