# Action: NEW_IMAGE

The NEW_IMAGE action creates a new image

This action is valid for use in the Generate block.


## Parameters:

| Property   | Type   | Valid values                                           | Description                     |
|------------|--------|--------------------------------------------------------|---------------------------------|
| ImageName  | string | Any loaded image name                                  | The name of the image to modify |
| Properties | Object | An object with [properties](aaaaaaaaaaaaaaaaaaaaaaaaa) | The properties to modify        |

### Properties

The NEW_IMAGE action accepts the `Size` property, which describes the size of the image

It is defined as an array of type `[]int`, with the values `[width, height]`.


## Example

### Making a new image

```json
{
	...
	"Generate": {
		"Actions": [
			{
				"ActionType": "NEW_IMAGE",
				"ImageName": "image1",
				"Properties": {
					"Size": [512, 512]
				}
			}
		]
	},
	...
}
```