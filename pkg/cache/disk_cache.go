package cache

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sync"
)

type diskCache struct {
	directory string
	lock      sync.Mutex
}

// NewDiskCache returns a new Cache
func NewDiskCache(directory string) (Cache, error) {
	// Create directory
	err := os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("unable to create directory '%s': %v", directory, err)
	}

	return &diskCache{
		directory: directory,
	}, nil
}

func (d *diskCache) Get(key string) ([]byte, error) {
	filename := d.buildFileName(key)

	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %v", err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("unable to read file: %v", err)
	}

	return data, nil
}

func (d *diskCache) Set(key string, data []byte) error {
	filename := d.buildFileName(key)

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("unable to create file: %v", err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("unable to write items to file: %v", err)
	}

	return nil
}

func (d *diskCache) Clean(key string) error {
	return nil
}

func (d *diskCache) buildFileName(key string) string {
	hasher := sha1.New()
	hasher.Write([]byte(key))

	return path.Join(d.directory, hex.EncodeToString(hasher.Sum(nil)))
}
