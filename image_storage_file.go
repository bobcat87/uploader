package uploader

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type imageStorageFile struct {
	directory string
}

func NewImageStorageFile(config map[string]string) ImageStorage {
	if config["directory"] == "" {
		panic("File storage configuration needs to have a directory")
	}
	return &imageStorageFile{
		directory: strings.TrimRight(config["directory"], "/"),
	}
}

func (is *imageStorageFile) makeFilename(filename string) string {
	return fmt.Sprintf("%s/%s", is.directory, filename)
}

func (is *imageStorageFile) Put(filename string, imageData []byte) error {
	fd, err := os.Create(is.makeFilename(filename))
	if err != nil {
		return err
	}
	defer fd.Close()

	_, err = fd.Write(imageData)
	return err
}

func (is *imageStorageFile) Delete(filename string) error {
	return os.Remove(is.makeFilename(filename))
}

func (is *imageStorageFile) Get(filename string) ([]byte, error) {
	fd, err := os.Open(is.makeFilename(filename))
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	imageData, err := ioutil.ReadAll(fd)
	if err != nil {
		return nil, err
	}
	return imageData, nil
}

func (is *imageStorageFile) Exists(filename string) bool {
	_, err := os.Stat(is.makeFilename(filename))
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}
