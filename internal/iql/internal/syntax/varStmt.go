package iqlSyntax

import iqlSchema "discorddungeons.me/imageserver/iql/schema"

type VarStmt struct {
	name        iqlSchema.Token
	initializer Expr
}

func NewVarStmt(name iqlSchema.Token, initializer Expr) *VarStmt {
	return &VarStmt{
		name:        name,
		initializer: initializer,
	}
}

func (varStmt *VarStmt) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitVarStmt(*varStmt)
}
