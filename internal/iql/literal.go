package iql

import "fmt"

type LiteralType int

func (ltype LiteralType) String() string {
	switch ltype {
	case LiteralTypes.STRING_LITERAL:
		return "STRING_LITERAL"
	case LiteralTypes.INT_LITERAL:
		return "INT_LITERAL"
	case LiteralTypes.FLOAT32_LITERAL:
		return "FLOAT32_LITERAL"
	case LiteralTypes.FLOAT64_LITERAL:
		return "FLOAT64_LITERAL"

	}

	return "UNKNOWN_LITERAL"
}

type literalList struct {
	STRING_LITERAL  LiteralType
	INT_LITERAL     LiteralType
	FLOAT32_LITERAL LiteralType
	FLOAT64_LITERAL LiteralType
}

var LiteralTypes = &literalList{
	STRING_LITERAL:  0,
	INT_LITERAL:     1,
	FLOAT32_LITERAL: 2,
	FLOAT64_LITERAL: 3,
}

type Literal struct {
	lType LiteralType

	strVal     string
	intVal     int
	float32Val float32
	float64Val float64
}

func (literal Literal) String() string {
	switch literal.lType {
	case LiteralTypes.STRING_LITERAL:
		return fmt.Sprintf("%s %s", literal.lType, literal.strVal)
	case LiteralTypes.INT_LITERAL:
		return fmt.Sprintf("%s %d", literal.lType, literal.intVal)
	case LiteralTypes.FLOAT32_LITERAL:
		return fmt.Sprintf("%s %f", literal.lType, literal.float32Val)
	case LiteralTypes.FLOAT64_LITERAL:
		return fmt.Sprintf("%s %f", literal.lType, literal.float64Val)
	}

	return fmt.Sprintf("%s %s", literal.lType, "no value")
}
