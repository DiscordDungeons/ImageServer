package iqlSyntax

type LiteralExpr struct {
	value interface{}
}

func NewLiteralExpr(val interface{}) *LiteralExpr {
	return &LiteralExpr{
		value: val,
	}
}

func (litExpr *LiteralExpr) Accept(visitor StmtVisitor) interface{} {
	return nil
}