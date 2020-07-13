package storage

import "io"

type Interface interface {
	GetFile(string) (io.Reader, error)
	WriteFile(string, io.Reader)
}
