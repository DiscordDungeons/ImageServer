package iqlSyntax

type Stmt interface {
	Accept(visitor StmtVisitor) interface{}
}

type StmtVisitor interface {
	//VisitBlockStmt(stmt Block) interface{}
	VisitExpressionStmt(stmt *ExpressionStmt) interface{}
	//VisitFunctionStmt(stmt Function) interface{}
	//VisitIfStmt(stmt If) interface{}
	//VisitPrintStmt(stmt Print) interface{}
	//VisitReturnStmt(stmt Return) interface{}
	VisitVarStmt(stmt VarStmt) interface{}
	//VisitWhileStmt(stmt While) interface{}
}
