package iql

type iqlAction struct {
	ActionType ActionType
	Url        string
	ImageName  string
	Properties map[string]interface{}
}

type initTree struct {
	Actions []iqlAction
}

type generateTree struct {
	Actions []iqlAction
}

type returnTree struct {
	Actions []iqlAction
}

type IQLTree struct {
	Init     initTree
	Generate generateTree
	Return   returnTree
}
