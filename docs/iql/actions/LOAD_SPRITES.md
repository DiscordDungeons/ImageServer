# Action: LOAD_SPRITES

The LOAD_SPRITES action loads an image and splits it into sprites into the script execution.

This action is valid for use in the Init block.

The images loaded are named according to the format `<imagename>-<row>-<col>`.


## Parameters:

| Property  | Type   | Valid values                  | Description                          |
|-----------|--------|-------------------------------|--------------------------------------|
| Url       | string | Any URL that returns an image | The URL from which to load the image |
| ImageName | string | Any unique string             | The name used to track the image     |
| Properties| Object | The described object below    | The properties for the action        |

### Properties

The LOAD_SPRITES action accepts the `SpriteSize` property, which describes the size of a sprite.

It is defined as an array of type `[]int`, with the values `[width, height]`.

## Example

### Loading a spritesheet:

```json
{
	"Init": {
		"Actions": [
			{
				"ActionType": "LOAD_SPRITES",
				"Url": "https://res.discorddungeons.me/icon.png",
				"ImageName": "image1",
				"Properties": {
					"SpriteSize": [16, 16]
				}
			}
		]
	}
	...
}
```