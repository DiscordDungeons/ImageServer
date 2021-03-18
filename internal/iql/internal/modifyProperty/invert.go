package modifyProperty

import (
	"image"

	"github.com/anthonynsimon/bild/effect"
)

// Inverts the given loadedImage
func HandleInvert(loadedImage image.Image, propertyValue interface{}, loadedImages map[string]image.Image) (image.Image, error) {

	inverted := effect.Invert(loadedImage)

	return inverted, nil
}
