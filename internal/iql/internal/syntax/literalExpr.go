package iqlSyntax

type LiteralExpr struct {
	Value interface{}
}

func NewLiteralExpr(val interface{}) *LiteralExpr {
	return &LiteralExpr{
		Value: val,
	}
}

func (litExpr *LiteralExpr) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitLiteralExpr(litExpr)
}
