package iqlSyntax

type LoadStmt struct {
	url       string
	imageName string
	//stmt      Expr
}

func NewLoadStmt(url string, imageName string) *LoadStmt {
	return &LoadStmt{
		url:       url,
		imageName: imageName,
	}
}

func (loadStmt *LoadStmt) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitLoadStmt(loadStmt)
}
