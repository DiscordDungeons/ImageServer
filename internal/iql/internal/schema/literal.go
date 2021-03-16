package iqlSchema

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
	NO_TYPE         LiteralType
	STRING_LITERAL  LiteralType
	INT_LITERAL     LiteralType
	FLOAT32_LITERAL LiteralType
	FLOAT64_LITERAL LiteralType
}

var LiteralTypes = &literalList{
	NO_TYPE:         0,
	STRING_LITERAL:  1,
	INT_LITERAL:     2,
	FLOAT32_LITERAL: 3,
	FLOAT64_LITERAL: 4,
}

type Literal struct {
	LType LiteralType

	StrVal     string
	IntVal     int
	Float32Val float32
	Float64Val float64
}

func (literal Literal) String() string {
	switch literal.LType {
	case LiteralTypes.STRING_LITERAL:
		return fmt.Sprintf("%s %s", literal.LType, literal.StrVal)
	case LiteralTypes.INT_LITERAL:
		return fmt.Sprintf("%s %d", literal.LType, literal.IntVal)
	case LiteralTypes.FLOAT32_LITERAL:
		return fmt.Sprintf("%s %f", literal.LType, literal.Float32Val)
	case LiteralTypes.FLOAT64_LITERAL:
		return fmt.Sprintf("%s %f", literal.LType, literal.Float64Val)
	}

	return fmt.Sprintf("%s %s", literal.LType, "no value")
}
