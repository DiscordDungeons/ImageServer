package iqlSyntax

type Block struct{}

func (stmt *Block) Accept(visitor StmtVisitor) {
	//return visitor.VisitBlockStmt(stmt)
}
