package iql

type ModifyProperty = string

type modifyProperty struct {
	GRAYSCALE ModifyProperty
}

var ModifyProperties = &modifyProperty{
	GRAYSCALE: "Grayscale",
}
