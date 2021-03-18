package iql

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"net/http"

	"github.com/anthonynsimon/bild/transform"
)

// Max image size in bytes
// Defined as MB * 1024 * 1024
const MAX_IMAGE_SIZE = 5 * 1024 * 1024

// Loads a spritesheet from url as imageName, with the spriteSize [x, y]
func LoadSpritesheet(url string, spriteSize []int, imageName string) (map[string]image.Image, error) {
	sprites := make(map[string]image.Image)

	img, err := LoadImageAction(url)

	if err != nil {
		return nil, err
	}

	imgSize := img.Bounds().Size()

	cols, rows := imgSize.X/spriteSize[0], imgSize.Y/spriteSize[1]

	for currentCol := 0; currentCol < cols; currentCol++ {
		for currentRow := 0; currentRow < rows; currentRow++ {
			spriteName := fmt.Sprintf("%s-%d-%d", imageName, currentRow, currentCol)

			cropRect := image.Rect(
				(spriteSize[0] * (currentCol)), (spriteSize[1] * (currentRow)),
				(spriteSize[0] * (currentCol + 1)), (spriteSize[1] * (currentRow + 1)),
			)

			sprites[spriteName] = transform.Crop(img, cropRect)
		}
	}

	return sprites, nil
}

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
