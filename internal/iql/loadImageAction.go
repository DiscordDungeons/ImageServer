package iql

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
)

// Max image size in bytes
// Defined as MB * 1024 * 1024
const MAX_IMAGE_SIZE = 5 * 1024 * 1024

func LoadImageAction(url string) (image.Image, error) {
	resp, err := http.Head(url)

	if err != nil {
		return nil, err
	}

	if resp.ContentLength > MAX_IMAGE_SIZE {
		return nil, fmt.Errorf("image size too big. Got %d/%d bytes", resp.ContentLength, MAX_IMAGE_SIZE)
	}

	resp, err = http.Get(url)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error loading image from %s: %s", url, http.StatusText(resp.StatusCode))
	}

	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("error decoding image from %s: %s", url, err)
	}

	return img, nil
}
