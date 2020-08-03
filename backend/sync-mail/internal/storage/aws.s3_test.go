package storage_test

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/ynsgnr/scribo/backend/sync-mail/internal/config"
	"github.com/ynsgnr/scribo/backend/sync-mail/internal/storage"
)

func TestStorageS3(t *testing.T) {
	cfg, err := config.InitConfig()
	if err != nil {
		t.Fatal(err)
	}
	ses := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	s3 := storage.NewStorageS3(s3manager.NewDownloader(ses), s3manager.NewUploader(ses), cfg.S3Bucket, cfg.TempFolder)

	//Create a file
	tmpFile := path.Join(cfg.TempFolder, "test.md")
	tmpFileContent := "test data"
	err = ioutil.WriteFile(tmpFile, []byte(tmpFileContent), os.ModeExclusive)
	if err != nil {
		t.Fatal(err)
	}

	//Upload file
	file, err := os.Open(tmpFile)
	if err != nil {
		t.Fatal(err)
	}
	tmpKey := path.Base(tmpFile)
	url, err := s3.UploadFile(tmpKey, file)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Uploaded: %s", url)
	file.Close()

	//Delete file
	err = os.Remove(tmpFile)
	if err != nil {
		t.Fatal(err)
	}

	//Download file and check
	newFile, err := s3.DownloadFile(tmpKey)
	if err != nil {
		t.Fatal(err)
	}

	data, err := ioutil.ReadFile(newFile)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != tmpFileContent {
		t.Errorf("Expected file data: %s, got: %s", tmpFileContent, string(data))
	}

	//Delete file
	err = os.Remove(newFile)
	if err != nil {
		t.Fatal(err)
	}
}
