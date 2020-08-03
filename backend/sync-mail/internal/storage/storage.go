package storage

import "io"

type Interface interface {
	DownloadFile(string) (string, error)
	UploadFile(string, io.Reader) (string, error)
}
