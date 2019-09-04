//抽象语法树相关
package main

//抽象语法树的节点
type ASTNode struct {
	Parent   *ASTNode    //父节点
	Children []*ASTNode  //子节点
	Type     ASTNodeType //节点类型
	Text     string      //节点内容
}

//抽象语法树类型
type ASTNodeType string

//抽象语法树的类型常量
const (
	ASTNodeType_Programm ASTNodeType = "Programm" //程序入口，根节点

	ASTNodeType_IntDeclaration = "IntDeclaration" //整型变量声明
	ASTNodeType_ExpressionStmt = "ExpressionStmt" //表达式语句，即表达式后面跟个分号
	ASTNodeType_AssignmentStmt = "AssignmentStmt" //赋值语句

	ASTNodeType_Primary        = "Primary"        //基础表达式
	ASTNodeType_Multiplicative = "Multiplicative" //乘法表达式
	ASTNodeType_Additive       = "Additive"       //加法表达式

	ASTNodeType_Identifier = "Identifier" //标识符
	ASTNodeType_IntLiteral = "IntLiteral" //整型字面量
)
