package iqlSyntax

import (
	iqlSchema "discorddungeons.me/imageserver/iql/schema"
)

type VariableExpr struct {
	name iqlSchema.Token
}

func NewVariable(name iqlSchema.Token) *VariableExpr {
	return &VariableExpr{
		name: name,
	}
}

func (expr *VariableExpr) Accept(visitor ExprVisitor) interface{} {
	//return visitor.VisitBlockStmt(stmt)
	return visitor.VisitVariableExpr(expr)
}
