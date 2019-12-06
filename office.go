package gotenberg

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"strconv"
)

const landscapeOffice string = "landscape"

// OfficeRequest facilitates Office documents
// conversion with the Gotenberg API.
type OfficeRequest struct {
	fileReader  io.Reader
	filename    string

	*request
}

// NewOfficeRequest create OfficeRequest.
func NewOfficeRequest(filename string, fileReader io.Reader) (*OfficeRequest, error) {
	if fileReader == nil {
		return nil, fmt.Errorf("file reader does not exist")
	}

	return &OfficeRequest{fileReader, filename, newRequest()}, nil
}

// Landscape sets landscape form field.
func (req *OfficeRequest) Landscape(isLandscape bool) {
	req.values[landscapeOffice] = strconv.FormatBool(isLandscape)
}

func (req *OfficeRequest) postURL() string {
	return "/convert/office"
}

func (req *OfficeRequest) formFiles() map[string]string {
	files := make(map[string]string)
	files[req.filename] = req.filename

	return files
}

func (req *OfficeRequest) multipartForm() (io.Reader, string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	defer writer.Close() // nolint: errcheck
	part, err := writer.CreateFormFile("files", req.filename)
	if err != nil {
		return nil, "", fmt.Errorf("%s: creating form file: %v", req.filename, err)
	}

	_, err = io.Copy(part, req.fileReader)
	if err != nil {
		return nil, "", fmt.Errorf("%s: copying file: %v", req.filename, err)
	}

	for name, value := range req.formValues() {
		if err := writer.WriteField(name, value); err != nil {
			return nil, "", fmt.Errorf("%s: writing form field: %v", name, err)
		}
	}

	return body, writer.FormDataContentType(), nil
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = Request(new(OfficeRequest))
)
