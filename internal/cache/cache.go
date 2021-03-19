package cache

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"image"
	"image/png"
	_ "image/png"
	"os"

	_ "golang.org/x/image/webp"
)

type Cache struct {
	directory string
}

func NewCache(directory string) *Cache {
	_ = os.Mkdir(directory, os.ModeDir)

	return &Cache{
		directory: directory,
	}
}

// Computes the hash of the IQLTree
func (cache *Cache) ComputeHash(data []byte) string {
	return fmt.Sprintf("%x", md5.Sum(data))
}

// Saves a png file as fileName
// The filename should ideally be a MD5 sum, with the file extension attached.
func (cache *Cache) SavePngFile(fileName string, file image.Image) error {

	cacheFile, err := os.Create(cache.directory + "/" + fileName)

	if err != nil {
		return err
	}

	w := bufio.NewWriter(cacheFile)

	err = png.Encode(w, file)

	if err != nil {
		return err
	}

	return nil
}

// Checks if a file exists in the cache directory
func (cache *Cache) HasFile(fileName string) bool {
	_, err := os.Stat(cache.directory + "/" + fileName)
	return err == nil
}
