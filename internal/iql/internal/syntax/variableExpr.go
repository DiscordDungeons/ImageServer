package iqlSyntax

import (
	iqlSchema "discorddungeons.me/imageserver/iql/schema"
)

type Variable struct {
	name iqlSchema.Token
}

func NewVariable(name iqlSchema.Token) *Variable {
	return &Variable{
		name: name,
	}
}

func (expr *Variable) Accept(visitor ExprVisitor) interface{} {
	//return visitor.VisitBlockStmt(stmt)
	return visitor.VisitVariableExpr(expr)
}
