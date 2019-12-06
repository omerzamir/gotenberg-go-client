package gotenberg

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/meateam/gotenberg-go-client/v6/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOffice(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	filename := "document.docx"
	file, err := os.Open(test.OfficeTestFilePath(t, filename))
	assert.Nil(t, err)
	defer file.Close()
	req, err := NewOfficeRequest(filename, file)
	require.Nil(t, err)
	req.ResultFilename("foo.pdf")
	req.WaitTimeout(5)
	req.Landscape(false)
	dirPath, err := test.Rand()
	require.Nil(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = os.MkdirAll(filepath.Dir(dest), 0755)
	assert.Nil(t, err)
	newFile, err := os.Create(dest)
	assert.Nil(t, err)
	defer newFile.Close()
	err = c.StoreWriter(req, newFile)
	assert.Nil(t, err)
	assert.FileExists(t, dest)
	err = os.RemoveAll(dirPath)
	assert.Nil(t, err)
}
