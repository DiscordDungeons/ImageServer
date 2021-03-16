package iql

import (
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"image/color"
)

type IQLRunner struct {
	tree IQLTree

	loadedImages map[string]image.Image
}

func NewIQLRunner() *IQLRunner {
	return &IQLRunner{
		loadedImages: make(map[string]image.Image),
	}
}

func (runner *IQLRunner) RunIQL(code string) (map[string]image.Image, error) {
	err := json.Unmarshal([]byte(code), &runner.tree)

	if err != nil {
		return nil, err
	}

	for _, action := range runner.tree.Init.Actions {
		switch action.ActionType {
		case ActionTypes.LOAD_IMAGE:
			if action.Url == "" {
				return nil, errors.New("no url provided for image")
			}

			if action.ImageName == "" {
				return nil, errors.New("no name provided for image")
			}

			image, err := LoadImageAction(action.Url)

			if err != nil {
				return nil, err
			}

			runner.loadedImages[action.ImageName] = image
		}
		fmt.Println(action)
	}

	for _, action := range runner.tree.Generate.Actions {
		switch action.ActionType {
		case ActionTypes.MODIFY_IMAGE:
			if action.ImageName == "" {
				return nil, errors.New("no image name provided for action")
			}

			if runner.loadedImages[action.ImageName] == nil {
				return nil, fmt.Errorf("error modifying image %s: it's not loaded", action.ImageName)
			}

			loadedImage := runner.loadedImages[action.ImageName]

			for property, _ := range action.Properties {
				fmt.Printf("property: %s\n", property)

				switch property {
				case ModifyProperties.GRAYSCALE:
					fmt.Println("GRAYSCALE!")
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

					runner.loadedImages[action.ImageName] = nil

					runner.loadedImages[action.ImageName] = grayImage
				}

			}
		}
		fmt.Println(action)
	}

	returningImages := make(map[string]image.Image)

	for _, action := range runner.tree.Return.Actions {
		switch action.ActionType {
		case ActionTypes.RETURN_IMAGE:
			if runner.loadedImages[action.ImageName] == nil {
				return nil, fmt.Errorf("error returning image %s: it's not loaded", action.ImageName)
			}

			returningImages[action.ImageName] = runner.loadedImages[action.ImageName]
		}
		fmt.Println(action)
	}

	return returningImages, nil
}
