package modifyProperty

import (
	"fmt"
	"image"
	"sync"

	"discorddungeons.me/imageserver/iql/iqlTypes"
)

type ModifyProperty = string

type modifyProperty struct {
	GRAYSCALE ModifyProperty
}

var ModifyProperties = &modifyProperty{
	GRAYSCALE: "Grayscale",
}

type modificationHandler struct {
	registeredProperties map[string]func(loadedImage image.Image, propertyValue interface{}, loadedImages map[string]image.Image) (image.Image, error)
}

var modificationHandlerInstance *modificationHandler
var once sync.Once

// Gets the modificationHandler instance
func GetModificationHandler() *modificationHandler {
	once.Do(func() {
		modificationHandlerInstance = &modificationHandler{
			registeredProperties: make(map[string]func(loadedImage image.Image, propertyValue interface{}, loadedImages map[string]image.Image) (image.Image, error)),
		}

		modificationHandlerInstance.RegisterProperty("Grayscale", HandleGrayscale)
		modificationHandlerInstance.RegisterProperty("Invert", HandleInvert)
		modificationHandlerInstance.RegisterProperty("PasteImage", HandlePasteImage)
	})

	return modificationHandlerInstance
}

// Registers the given property to be handled using the given action function
func (handler *modificationHandler) RegisterProperty(property string, action func(loadedImage image.Image, propertyValue interface{}, loadedImages map[string]image.Image) (image.Image, error)) {
	handler.registeredProperties[property] = action
}

// Handles modification
func (handler *modificationHandler) HandleModification(modImage image.Image, action iqlTypes.IQLAction, loadedImages map[string]image.Image) (image.Image, error) {
	modifiedImage := modImage

	for property, value := range action.Properties {
		if handler.registeredProperties[property] == nil {
			return nil, fmt.Errorf("error modifying image %s: the property %s isn't registered", action.ImageName, property)
		}

		fmt.Println(value)

		newImage, err := handler.registeredProperties[property](modifiedImage, value, loadedImages)

		if err != nil {
			return nil, err
		}

		modifiedImage = newImage
	}

	return modifiedImage, nil
}
