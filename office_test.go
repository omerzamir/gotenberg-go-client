package gotenberg

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/omerzamir/gotenberg-go-client/v7/test"
)

func TestOffice(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	doc, err := NewDocumentFromPath("document.docx", test.OfficeTestFilePath(t, "document.docx"))
	require.Nil(t, err)
	req := NewOfficeRequest(doc)
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

func TestOfficePageRanges(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	doc, err := NewDocumentFromPath("document.docx", test.OfficeTestFilePath(t, "document.docx"))
	require.Nil(t, err)
	req := NewOfficeRequest(doc)
	req.PageRanges("1-1")
	resp, err := c.Post(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestOfficeWebhook(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	doc, err := NewDocumentFromPath("document.docx", test.OfficeTestFilePath(t, "document.docx"))
	require.Nil(t, err)
	req := NewOfficeRequest(doc)
	req.WebhookURL("https://google.com")
	req.WebhookURLTimeout(5.0)
	req.AddWebhookURLHTTPHeader("A-Header", "Foo")
	resp, err := c.Post(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}
