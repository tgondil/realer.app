package cloudflareR2utils

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	cnfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"strings"
	"sync"
)

const (
	accessKey = "17490c28354435f379f5aa8878953522"                                 //todo
	secretKey = "a0c0c57e76ca9bae3b03862c294c824be52b425ee2fdd522707af8e2be49252d" //todo
	region    = "auto"
)

var bucketName = "boilermake-xi-project"
var s3Config *s3.Client
var uploader *manager.Uploader
var downloader *manager.Downloader

func Init() {
	config, err := cnfg.LoadDefaultConfig(context.TODO(),
		cnfg.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     accessKey,
				SecretAccessKey: secretKey,
				CanExpire:       false,
			},
		}),
		cnfg.WithRegion(region),
	)
	if err != nil {
		fmt.Println("Error creating AWS session:", err)
		return
	}
	cfg := s3.NewFromConfig(config)
	s3Config = cfg
	uploader = manager.NewUploader(cfg)
	downloader = manager.NewDownloader(cfg)

}

var uploadBytePool = sync.Pool{New: func() any {
	return new(bytes.Buffer)
}}

func UploadBytes(b []byte, fileNameOnS3 string, contentType ...string) error {
	byteBuffer := uploadBytePool.Get().(*bytes.Buffer)
	defer func() { byteBuffer.Reset(); uploadBytePool.Put(byteBuffer) }()
	byteBuffer.Reset()
	byteBuffer.Write(b)
	var contentTypeStr string
	if len(contentType) > 0 {
		contentTypeStr = contentType[0]
	} else {
		contentTypeStr = GetContentType(fileNameOnS3)
	}
	_, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:      &bucketName,
		Key:         &fileNameOnS3,
		Body:        byteBuffer,
		ContentType: &contentTypeStr,
	})
	return err
}

func GetContentType(fileNameOnS3 string) (contentTypeStr string) {
	contentTypeStr = ""
	filePathLower := strings.ToLower(fileNameOnS3)
	contentTypeStr = "application/octet-stream"
	if strings.HasSuffix(filePathLower, ".pdf") {
		contentTypeStr = "application/pdf"
	} else if strings.HasSuffix(filePathLower, ".jpg") || strings.HasSuffix(filePathLower, ".jpeg") {
		contentTypeStr = "image/jpeg"
	} else if strings.HasSuffix(filePathLower, ".png") {
		contentTypeStr = "image/png"
	} else if strings.HasSuffix(filePathLower, ".bmp") {
		contentTypeStr = "image/bmp"
	} else if strings.HasSuffix(filePathLower, ".gif") {
		contentTypeStr = "image/gif"
	} else if strings.HasSuffix(filePathLower, ".tif") || strings.HasSuffix(filePathLower, ".tiff") {
		contentTypeStr = "image/tiff"
	} else if strings.HasSuffix(filePathLower, ".svg") || strings.HasSuffix(filePathLower, ".xml") {
		contentTypeStr = "image/svg+xml"
	} else if strings.HasSuffix(filePathLower, ".txt") {
		contentTypeStr = "text/plain"
	} else if strings.HasSuffix(filePathLower, ".rtf") {
		contentTypeStr = "text/rtf"
	} else if strings.HasSuffix(filePathLower, ".doc") || strings.HasSuffix(filePathLower, ".docx") {
		contentTypeStr = "application/msword"
	}
	return
}

func GetFile(fileNameOnS3 string) ([]byte, error) {
	headObject, err := s3Config.HeadObject(context.TODO(), &s3.HeadObjectInput{
		Bucket: &bucketName,
		Key:    &fileNameOnS3,
	})
	if err != nil {
		return nil, err
	}

	buf := make([]byte, *(headObject.ContentLength))
	writeBuf := manager.NewWriteAtBuffer(buf)

	_, err = downloader.Download(context.TODO(), writeBuf, &s3.GetObjectInput{
		Bucket: &bucketName,
		Key:    &fileNameOnS3,
	})
	return writeBuf.Bytes(), err
}
