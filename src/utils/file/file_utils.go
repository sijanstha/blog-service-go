package fileutils

import (
	"errors"
	"io/ioutil"
	"strings"
)

var (
	ErrFileNotFound = errors.New("file not found")
)

func LoadResourceAsString(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", ErrFileNotFound
	}

	return strings.TrimSpace(string(data)), nil
}
