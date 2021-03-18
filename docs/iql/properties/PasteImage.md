# Property: PasteImage

The PasteImage action pastes a loaded image at a given position

This action is valid for use in a MODIFY_IMAGE block.


## Values

The PasteImage property accepts the value type `Object`.

### Parameters:

| Property  | Type     | Valid values                  | Description                                   |
|-----------|----------|-------------------------------|-----------------------------------------------|
| ImageName | `string` | Any unique string             | The name of the image to paste                |
| PasteAt   | `int[]`  | Two ints                      | The X and Y coordinates to paste the image at |

## Examples

### Pasting an image:

```json
{
	...
	"Generate": {
		"Actions": [
			{
				"ActionType": "MODIFY_IMAGE",
				"ImageName": "image1",
				"Properties": {
					"PasteImage": {
						"ImageName": "image2",
						"PasteAt": [0,0]
					}
				}
			}
		]
	},
	...
}
```
