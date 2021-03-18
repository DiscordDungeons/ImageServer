package modifyProperty

import (
	"fmt"
	"image"
	"image/color"

	"github.com/disintegration/imaging"
)

// Pastes image at a point
func HandlePasteImage(loadedImage image.Image, propertyValue interface{}, loadedImages map[string]image.Image) (image.Image, error) {
	pasteImageData, ok := propertyValue.(map[string]interface{})

	if !ok {
		return nil, fmt.Errorf("property PasteImage got value of type %T, but wanted object", propertyValue)
	}

	imageName, ok := pasteImageData["ImageName"].(string)

	if !ok {
		return nil, fmt.Errorf("property PasteImage.ImageName got value of type %T, but wanted string", pasteImageData["ImageName"])
	}

	if loadedImages[imageName] == nil {
		return nil, fmt.Errorf("can't paste image %s as it's not loaded", imageName)
	}

	pasteAtData, ok := pasteImageData["PasteAt"].([]interface{})

	if !ok {
		return nil, fmt.Errorf("property PasteImage.PasteAt got value of type %T, but wanted []float", pasteImageData["PasteAt"])
	}

	pasteAt := make([]float64, len(pasteAtData))

	for i := range pasteAtData {
		coord, ok := pasteAtData[i].(float64)

		if !ok {
			return nil, fmt.Errorf("property PasteImage.PasteAt[%d] got value of type %T, but wanted float", i, pasteAtData[i])
		}

		pasteAt[i] = coord
	}

	size := loadedImage.Bounds().Size()

	dst := imaging.New(size.X, size.Y, color.NRGBA{0, 0, 0, 0})

	dst = imaging.Paste(dst, loadedImage, image.Pt(0, 0))

	dst = imaging.Paste(dst, loadedImages[imageName], image.Pt(int(pasteAt[0]), int(pasteAt[1])))

	return dst, nil
}
