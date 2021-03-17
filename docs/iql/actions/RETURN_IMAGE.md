# Action: RETURN_IMAGE

The RETURN_IMAGE action returns an image.

This action is valid for use in the Return block.


## Parameters:

| Property   | Type   | Valid values                                           | Description                     |
|------------|--------|--------------------------------------------------------|---------------------------------|
| ImageName  | string | Any loaded image name                                  | The name of the image to return |

## Example

### Returning an image

```json
{
	...
	"Return": {
		"Actions": [
			{
				"ActionType": "RETURN_IMAGE",
				"ImageName": "image1"
			}
		]
	}
}
```