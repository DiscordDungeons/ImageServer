# Image Query Language

SQL-Like syntax for image modification

Example:

Making image gray:

```
LOAD IMAGE FROM URL <url> AS <image_name>

GENERATE
	WITH image_name DO
		SET GRAYSCALE TO 1
```

Compositing:

```
LOAD IMAGE FROM URL <url> AS floor
LOAD IMAGE FROM URL <url> AS top_image

GENERATE
	WITH floor DO
		GRID 8 8
	WITH top_image DO
		PLACE AT 100, 100

```