# Action: MODIFY_IMAGE

The MODIFY_IMAGE action modifies an image.

This action is valid for use in the Generate block.


## Parameters:

| Property   | Type   | Valid values                                           | Description                     |
|------------|--------|--------------------------------------------------------|---------------------------------|
| ImageName  | string | Any loaded image name                                  | The name of the image to modify |
| Properties | Object | An object with [properties](aaaaaaaaaaaaaaaaaaaaaaaaa) | The properties to modify        |

## Example

### Making an image grayscale:

```json
{
	...
	"Generate": {
		"Actions": [
			{
				"ActionType": "MODIFY_IMAGE",
				"ImageName": "image1",
				"Properties": {
					"Grayscale": true
				}
			}
		]
	},
	...
}
```