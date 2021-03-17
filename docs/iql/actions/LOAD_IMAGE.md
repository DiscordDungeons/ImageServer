# Action: LOAD_IMAGE

The LOAD_IMAGE action loads an image into the script execution.

This action is valid for use in the Init block.


## Parameters:

| Property  | Type   | Valid values                  | Description                          |
|-----------|--------|-------------------------------|--------------------------------------|
| Url       | string | Any URL that returns an image | The URL from which to load the image |
| ImageName | string | Any unique string             | The name used to track the image     |

## Example

### Loading an image:

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