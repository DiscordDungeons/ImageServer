# Property: Grayscale

The Grayscale action makes an image grayscale.

This action is valid for use in a MODIFY_IMAGE block.


## Values

The Grayscale property accepts two value types: `boolean` and `Object`.

If set to a boolean, no transparency is included.

The object type only accepts one subproperty, `IncludeTransparency`, which accepts a value type of `boolean`.

If `IncludeTransparency` is set to `true`, the grayscale image will include transparency.

## Examples

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

### Making an image grayscale, with transparency:

```json
{
	...
	"Generate": {
		"Actions": [
			{
				"ActionType": "MODIFY_IMAGE",
				"ImageName": "image1",
				"Properties": {
					"Grayscale": {
						"IncludeTransparency": true
					}
				}
			}
		]
	},
	...
}
```