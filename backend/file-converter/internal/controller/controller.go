package controller

import (
	"github.com/ynsgnr/scribo/backend/common/schema/protobuf/generated/file"
	"github.com/ynsgnr/scribo/backend/file-converter/internal/storage"
)

type Interface interface {
	ConvertFile(*file.ConvertFile) (*file.ConvertFile, error)
}

func NewController(storage storage.Interface) Interface {
	return &controller{
		storage: storage,
	}
}

type controller struct {
	storage storage.Interface
}
