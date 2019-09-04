//语义分析
package parser

import (
	"craft/lib"
	"fmt"
)

/**
 * 一个简单的语法解析器。
 * 能够解析简单的表达式、变量声明和初始化语句、赋值语句。
 * 它支持的语法规则为：
 *
 * programm -> intDeclare | expressionStatement | assignmentStatement
 * intDeclare -> 'int' Id ( = additive) ';'
 * expressionStatement -> addtive ';'
 * addtive -> multiplicative ( (+ | -) multiplicative)*
 * multiplicative -> primary ( (* | /) primary)*
 * primary -> IntLiteral | Id | (additive)
 */
type Parser struct{}

func NewParser() Parser {
	return Parser{}
}

func (parse *Parser) Parse(tokens *([]lib.Token)) lib.ASTNode {
	reader := NewTokenReader(tokens)
	fmt.Println(reader)
	return lib.ASTNode{}
}
