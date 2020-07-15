package controller

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/ynsgnr/scribo/backend/common/schema/protobuf/generated/file"
)

func (c *controller) ConvertFile(file2convert *file.ConvertFile) (*file.ConvertFile, error) {
	origFile, err := c.storage.DownloadFile(path.Base(file2convert.FileLocation))
	if err != nil {
		return nil, err
	}
	//Get output path
	fileName, _ := filepath.Split(path.Base(origFile))
	convertedFile := path.Join(path.Dir(origFile), fmt.Sprintf("%s.%s", fileName, file2convert.Target))
	//Convert file
	calibreConverter, err := exec.LookPath("ebook-convert")
	if err != nil {
		return nil, err
	}
	cmd := &exec.Cmd{
		Path:   calibreConverter,
		Args:   []string{origFile, convertedFile},
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	err = cmd.Run()
	if err != nil {
		return nil, err
	}
	upload, err := os.Open(convertedFile)
	if err != nil {
		return nil, err
	}
	convertedURL, err := c.storage.UploadFile(path.Base(convertedFile), upload)
	if err != nil {
		return nil, err
	}
	return &file.ConvertFile{
		FileID:       file2convert.FileID,
		FileLocation: convertedURL,
		Target:       file2convert.Target,
	}, nil
}
