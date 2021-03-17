# Image Query Language

IQL is a JSON object, reffered to as the "IQL Tree", describing steps to take in order to modify images.

## IQL Tree

The IQL Tree consists of three main parts, which control the actions to take in the lifecycle, and are all required.

### Init Block

The Init block controls what actions to take upon initialization of the script, such as loading images for modification.

#### Examples:

##### Loading an image:

This example uses the [LOAD_IMAGE](actions/LOAD_IMAGE.md) action.

```json
{
	"Init": {
		"Actions": [
			{
				"ActionType": "LOAD_IMAGE",
				"Url": "https://res.discorddungeons.me/icon.png",
				"ImageName": "image1"
			}
		]
	}
	...
}
```

### Generate Block

The Generate block controls what actions to take upon the generation phase of the script, mainly which modifications to apply to images.

#### Examples:

##### Making an image to grayscale

This example uses the [MODIFY_IMAGE](actions/MODIFY_IMAGE.md) action.

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

### Return Block

The Return block controls what actions to take upon the return phase of the script, mainly what images to return.

#### Examples:

##### Returning an image

This example uses the [RETURN_IMAGE](actions/RETURN_IMAGE.md) action.

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
