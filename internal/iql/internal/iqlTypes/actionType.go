package iqlTypes

type ActionType = string

type actionType struct {
	LOAD_IMAGE   ActionType
	MODIFY_IMAGE ActionType
	RETURN_IMAGE ActionType
}

var ActionTypes = &actionType{
	LOAD_IMAGE:   "LOAD_IMAGE",
	MODIFY_IMAGE: "MODIFY_IMAGE",
	RETURN_IMAGE: "RETURN_IMAGE",
}
