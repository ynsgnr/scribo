package storage

import (
	"io"
	"net/url"
	"os"
	"path"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// Creates a new storage object based on S3
func NewStorageS3(downloader *s3manager.Downloader, uploader *s3manager.Uploader, bucket, tmp string) Interface {
	return &storageS3{
		downloader: downloader,
		uploader:   uploader,
		bucket:     bucket,
		tmp:        tmp,
	}
}

type storageS3 struct {
	downloader *s3manager.Downloader
	uploader   *s3manager.Uploader
	bucket     string
	tmp        string
}

func (storage *storageS3) DownloadFile(location string) (string, error) {
	key, err := url.QueryUnescape(path.Base(location))
	if err != nil {
		return "", err
	}
	filePath := path.Join(storage.tmp, key)
	err = os.Mkdir(storage.tmp, os.ModeDir)
	if err != nil && !os.IsExist(err) {
		return "", err
	}
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	_, err = storage.downloader.Download(file, &s3.GetObjectInput{
		Bucket: aws.String(storage.bucket),
		Key:    aws.String(key),
	})
	return filePath, err
}

func (storage *storageS3) UploadFile(location string, data io.Reader) (string, error) {
	key := path.Base(location)
	out, err := storage.uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(storage.bucket),
		Key:    aws.String(key),
		Body:   data,
	})
	if err != nil {
		return "", err
	}
	return out.Location, nil
}
