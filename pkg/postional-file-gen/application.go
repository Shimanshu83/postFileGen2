package postionalfilegen

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Record struct {
	Value  string
	End    int
	Length int
}

type DataMapperObj struct {
	Header       map[string]Record
	HeaderLength int
	Details      []map[string]Record
	DetailLength int
	FileName     string
	BatchNumber  string
}

func PadValue(value string, length int) string {
	if len(value) < length {
		return value + strings.Repeat(" ", length-len(value))
	}
	return value
}

func (d *DataMapperObj) String() string {
	var builder strings.Builder

	headerString := strings.Repeat(" ", d.HeaderLength)

	for _, record := range d.Header {
		value := PadValue(record.Value, record.Length)
		index := record.End - record.Length
		headerString = headerString[:index] + value + headerString[index+len(value):]
	}

	builder.WriteString(headerString)
	builder.WriteString("\n")

	for _, detail := range d.Details {
		detailString := strings.Repeat(" ", d.DetailLength)
		for _, record := range detail {
			value := PadValue(record.Value, record.Length)
			index := record.End - record.Length
			detailString = detailString[:index] + value + detailString[index+len(value):]

		}

		builder.WriteString(detailString)
		builder.WriteString("\n")
	}

	return builder.String()
}

func (d *DataMapperObj) CreateZipFile() (*bytes.Buffer, error) {
	var buf bytes.Buffer

	dataString := d.String()

	_, err := buf.WriteString(dataString)
	if err != nil {
		return nil, fmt.Errorf("error writing string representation to buffer: %v", err)
	}

	zipBuffer := new(bytes.Buffer)

	zipWriter := zip.NewWriter(zipBuffer)
	defer zipWriter.Close()

	// Create a new file in the ZIP archive
	writer, err := zipWriter.Create(d.FileName + "_" + d.BatchNumber)
	if err != nil {
		return nil, fmt.Errorf("error creating file in ZIP archive: %v", err)
	}

	// Copy the string representation from the buffer to the file in the ZIP archive
	_, err = io.Copy(writer, &buf)
	if err != nil {
		return nil, fmt.Errorf("error copying string representation to ZIP file: %v", err)
	}

	return zipBuffer, nil
}
