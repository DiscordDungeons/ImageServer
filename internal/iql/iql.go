package iql

import (
	"encoding/json"
	"errors"
	"fmt"
	"image"

	"discorddungeons.me/imageserver/iql/iqlTypes"
	"discorddungeons.me/imageserver/iql/modifyProperty"
)

type IQLRunner struct {
	tree iqlTypes.IQLTree

	loadedImages map[string]image.Image
}

// NewIQLRunner creates a new IQLRunner
func NewIQLRunner() *IQLRunner {
	return &IQLRunner{
		loadedImages: make(map[string]image.Image),
	}
}

// RunIQL runs the provided code, and returns a map of images by name, or an error
func (runner *IQLRunner) RunIQL(code string) (map[string]image.Image, error) {
	err := json.Unmarshal([]byte(code), &runner.tree)

	if err != nil {
		return nil, err
	}

	for _, action := range runner.tree.Init.Actions {
		switch action.ActionType {
		case iqlTypes.ActionTypes.LOAD_IMAGE:
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
		case iqlTypes.ActionTypes.MODIFY_IMAGE:
			if action.ImageName == "" {
				return nil, errors.New("no image name provided for action")
			}

			if runner.loadedImages[action.ImageName] == nil {
				return nil, fmt.Errorf("error modifying image %s: it's not loaded", action.ImageName)
			}

			loadedImage := runner.loadedImages[action.ImageName]

			modifiedImage, err := modifyProperty.GetModificationHandler().HandleModification(loadedImage, action)

			if err != nil {
				return nil, err
			}

			runner.loadedImages[action.ImageName] = nil
			runner.loadedImages[action.ImageName] = modifiedImage
		}
		fmt.Println(action)
	}

	returningImages := make(map[string]image.Image)

	for _, action := range runner.tree.Return.Actions {
		switch action.ActionType {
		case iqlTypes.ActionTypes.RETURN_IMAGE:
			if runner.loadedImages[action.ImageName] == nil {
				return nil, fmt.Errorf("error returning image %s: it's not loaded", action.ImageName)
			}

			returningImages[action.ImageName] = runner.loadedImages[action.ImageName]
		}
		fmt.Println(action)
	}

	return returningImages, nil
}
