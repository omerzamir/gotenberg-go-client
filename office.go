package gotenberg

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"strconv"
)

const (
	landscapeOffice  string = "landscape"
	pageRangesOffice string = "pageRanges"
)

// OfficeRequest facilitates Office documents
// conversion with the Gotenberg API.
type OfficeRequest struct {
	docs []Document
	readerCopyBuffer []byte
	*request
}

// NewOfficeRequest create OfficeRequest.
func NewOfficeRequest(docs ...Document) *OfficeRequest {
	return &OfficeRequest{docs, readerCopyBuffer:nil, newRequest()}
}

// NewOfficeRequestWithBuffer create OfficeRequest.
func NewOfficeRequestWithBuffer(docs ...Document,  readerCopyBuffer []byte) *OfficeRequest {
	return &OfficeRequest{docs, readerCopyBuffer: readerCopyBuffer, newRequest()}
}

// Landscape sets landscape form field.
func (req *OfficeRequest) Landscape(isLandscape bool) {
	req.values[landscapeOffice] = strconv.FormatBool(isLandscape)
}

// PageRanges sets pageRanges form field.
func (req *OfficeRequest) PageRanges(ranges string) {
	req.values[pageRangesOffice] = ranges
}

func (req *OfficeRequest) postURL() string {
	return "/convert/office"
}

func (req *OfficeRequest) formFiles() map[string]Document {
	files := make(map[string]Document)
	for _, doc := range req.docs {
		files[doc.Filename()] = doc
	}

	return files
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = Request(new(OfficeRequest))
)
