package iqlSyntax

type Expr interface {
	Accept(visitor ExprVisitor)
}

type ExprVisitor interface {
	VisitAssignExpr(expr Assign) interface{}
	VisitBinaryExpr(expr Binary) interface{}
	VisitCallExpr(expr Call) interface{}
	VisitGroupingExpr(expr Grouping) interface{}
	VisitLiteralExpr(expr Literal) interface{}
	VisitLogicalExpr(expr Logical) interface{}
	VisitUnaryExpr(expr Unary) interface{}
	VisitVariableExpr(expr *Variable) interface{}
}
