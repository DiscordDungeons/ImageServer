package iqlSyntax

type ExpressionStmt struct {
	expression Expr
}

func NewExpressionStmt(expression Expr) *ExpressionStmt {
	return &ExpressionStmt{
		expression: expression,
	}
}

func (expressionStmt *ExpressionStmt) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitExpressionStmt(expressionStmt)
}
