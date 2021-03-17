package iqlTypes

type IQLAction struct {
	ActionType ActionType
	Url        string
	ImageName  string
	Properties map[string]interface{}
}

type initTree struct {
	Actions []IQLAction
}

type generateTree struct {
	Actions []IQLAction
}

type returnTree struct {
	Actions []IQLAction
}

type IQLTree struct {
	Init     initTree
	Generate generateTree
	Return   returnTree
}
