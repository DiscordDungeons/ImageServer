package iqlTypes

type ActionType = string

type actionType struct {
	LOAD_IMAGE   ActionType
	LOAD_SPRITES ActionType
	MODIFY_IMAGE ActionType
	RETURN_IMAGE ActionType
	NEW_IMAGE    ActionType
}

var ActionTypes = &actionType{
	LOAD_IMAGE:   "LOAD_IMAGE",
	LOAD_SPRITES: "LOAD_SPRITES",
	MODIFY_IMAGE: "MODIFY_IMAGE",
	RETURN_IMAGE: "RETURN_IMAGE",
	NEW_IMAGE:    "NEW_IMAGE",
}
