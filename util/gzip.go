package util

import (
	"bytes"
	gzip2 "compress/gzip"
	"github.com/pkg/errors"
	"io"
)

func Gzip(b []byte) ([]byte, error) {
	res := bytes.Buffer{}
	writer := gzip2.NewWriter(&res)
	_, err := writer.Write(b)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	err = writer.Close()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return res.Bytes(), nil
}

func UnGzip(b []byte) ([]byte, error) {
	reader, err := gzip2.NewReader(bytes.NewReader(b))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer reader.Close()
	return io.ReadAll(reader)
}
