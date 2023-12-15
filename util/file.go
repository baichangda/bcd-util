package util

import (
	"bufio"
	"github.com/pkg/errors"
	"io"
	"os"
)

func ReadSplitAll_file(file string, split byte) ([][]byte, error) {
	open, err := os.Open(file)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer open.Close()
	return ReadSplitAll_reader(open, split)
}

func ReadSplitAll_reader(reader io.Reader, split byte) ([][]byte, error) {
	bufReader, ok := reader.(*bufio.Reader)
	if !ok {
		bufReader = bufio.NewReader(reader)
	}
	var res [][]byte
	for {
		last := false
		bytes, err := bufReader.ReadBytes(split)
		if err == nil {
			bytes = bytes[:len(bytes)-1]
		} else {
			if err == io.EOF {
				last = true
			} else {
				return nil, errors.WithStack(err)
			}
		}

		res = append(res, bytes)
		if last {
			break
		}
	}
	return res, nil
}

func ReadSplitOneByOne_file(file string, split byte, fn func(data []byte) (bool, error)) error {
	open, err := os.Open(file)
	if err != nil {
		return errors.WithStack(err)
	}
	defer open.Close()
	return ReadSplitOneByOne_reader(open, split, fn)
}

func ReadSplitOneByOne_reader(reader io.Reader, split byte, fn func(data []byte) (bool, error)) error {
	bufReader, ok := reader.(*bufio.Reader)
	if !ok {
		bufReader = bufio.NewReader(reader)
	}
	for {
		last := false
		bytes, err := bufReader.ReadBytes(split)
		if err == nil {
			bytes = bytes[:len(bytes)-1]
		} else {
			if err == io.EOF {
				last = true
			} else {
				return errors.WithStack(err)
			}
		}

		goon, err := fn(bytes)
		if err != nil {
			return err
		}
		if !goon {
			return nil
		}
		if last {
			break
		}
	}
	return nil
}
