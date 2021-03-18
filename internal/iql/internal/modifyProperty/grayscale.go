package modifyProperty

import (
	"fmt"
	"image"
	"image/color"
)

// Converts the given loadedImage to grayscale
func HandleGrayscale(loadedImage image.Image, propertyValue interface{}, loadedImages map[string]image.Image) (image.Image, error) {
	switch t := propertyValue.(type) {
	case map[string]interface{}:
		// We've got a json object
		subprops := propertyValue.(map[string]interface{})

		if _, ok := subprops["IncludeTransparency"]; ok {
			includeTransparency, ok := subprops["IncludeTransparency"].(bool)

			if !ok {
				return nil, fmt.Errorf("property IncludeTransparency got value of type %T, but wanted bool", subprops["IncludeTransparency"])
			}

			if includeTransparency {
				return HandleTransparentGrayscale(loadedImage, propertyValue, loadedImages)
			}
		}

	case bool:
		// Grayscale set to a bool
		isTransparent := propertyValue.(bool)
		if isTransparent {
			return HandleTransparentGrayscale(loadedImage, propertyValue, loadedImages)
		}
	default:
		return nil, fmt.Errorf("error applying property %s. value of type %T is invalid", "Grayscale", t)
	}

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

// Converts the given loadedImage to grayscale, while keeping transparency
func HandleTransparentGrayscale(loadedImage image.Image, propertyValue interface{}, loadedImages map[string]image.Image) (image.Image, error) {
	grayImage := image.NewRGBA(loadedImage.Bounds())

	bounds := loadedImage.Bounds()

	w, h := bounds.Max.X, bounds.Max.Y

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			oldColor := loadedImage.At(x, y)

			r, g, b, a := oldColor.RGBA()

			lum := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)

			grayColor := color.RGBA{
				R: uint8(lum / 256),
				G: uint8(lum / 256),
				B: uint8(lum / 256),
				A: uint8(a / 256),
			}

			grayImage.Set(x, y, grayColor)
		}
	}

	return grayImage, nil
}
