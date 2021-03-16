package iql

import (
	"fmt"

	iqlSyntax "discorddungeons.me/imageserver/iql/syntax"
)

type ASTPrinter struct{}

func NewASTPrinter() *ASTPrinter {
	return &ASTPrinter{}
}

func (astPrinter *ASTPrinter) Print(expr iqlSyntax.Expr) interface{} {
	return expr.Accept(astPrinter)
}

func (astPrinter *ASTPrinter) parenthesize(name string, exprs []iqlSyntax.Expr) string {
	str := fmt.Sprintf("(%s)", name)

	for _, expr := range exprs {
		str = fmt.Sprintf("%s %s", str, expr.Accept(astPrinter))
	}

	str = fmt.Sprintf("%s)", str)

	return str
}

func (astPrinter *ASTPrinter) VisitLiteralExpr(expr *iqlSyntax.LiteralExpr) interface{} {
	if expr.Value == nil {
		return "null"
	}

	return expr.Value
}

func (astPrinter *ASTPrinter) VisitVariableExpr(expr *iqlSyntax.VariableExpr) interface{} {
	return "nil"
}
