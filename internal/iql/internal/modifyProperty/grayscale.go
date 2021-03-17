package modifyProperty

import (
	"image"
	"image/color"
)

func HandleGrayscale(loadedImage image.Image) (image.Image, error) {
	grayImage := image.NewGray(loadedImage.Bounds())

	bounds := loadedImage.Bounds()

	w, h := bounds.Max.X, bounds.Max.Y

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			oldColor := loadedImage.At(x, y)
			grayColor := color.GrayModel.Convert(oldColor)

			grayImage.Set(x, y, grayColor)
		}
	}

	return grayImage, nil
}
