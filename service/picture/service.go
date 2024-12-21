package picture

import (
	"context"
	"derma/detect/config"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type PictureService struct {
	ctx    context.Context
	bucket *oss.Bucket
}

// NewPictureService new PictureService
func NewPictureService(ctx context.Context) *PictureService {
	// bucket init
	if config.OSS == nil {
		return &PictureService{ctx: ctx, bucket: nil}
	}

	client, err := oss.New("https://"+config.OSS.BucketName+"."+config.OSS.Endpoint, config.OSS.AccessKeyID, config.OSS.AccessKeySecret, oss.UseCname(true))
	if err != nil {
		hlog.Fatal(err)
	}

	bucket, err := client.Bucket(config.OSS.BucketName)
	if err != nil {
		hlog.Fatal(err)
	}
	return &PictureService{ctx: ctx, bucket: bucket}
}
