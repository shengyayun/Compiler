//语义分析
package parser

import (
	"craft/lib"
	"errors"
	"io"
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
 * primary -> IntLiteral | Id | '(' additive ')'
 */
type Parser struct{}

func NewParser() Parser {
	return Parser{}
}

//将token流解析为抽象语法树
func (parse *Parser) Parse(tokens *([]lib.Token)) (*lib.ASTNode, error) {
	reader := NewTokenReader(tokens)
	//只有一个根节点的抽象语法树
	tree := lib.ASTNode{Type: lib.ASTNodeType_Programm, Text: "pwc"}
	//处理函数数组
	handler := [3]func(reader *TokenReader) (*lib.ASTNode, error){intDeclare, expressionStatement, assignmentStatement}
	var err error
	var child *lib.ASTNode
	for {
		if reader.Peek() == nil {
			break
		}
		for _, method := range handler {
			if child, err = method(&reader); err == nil {
				tree.Append(child)
			} else if err != io.EOF {
				return nil, err
			}
		}
	}
	return &tree, nil
}

/**
 * 整型变量声明，如：
 * int a;
 * int a = 2*3;
 */
func intDeclare(reader *TokenReader) (*lib.ASTNode, error) {
	if token := reader.Peek(); token != nil && token.Type == lib.TokenType_Int { //int
		reader.Read()
		if token = reader.Peek(); token != nil && token.Type == lib.TokenType_Identifier { //a
			reader.Read()
			node := lib.NewASTNode(lib.ASTNodeType_IntDeclaration, token.Text)
			if token = reader.Peek(); token != nil && token.Type == lib.TokenType_Assignment { //=
				//该代码块中不匹配会被判定为语法错误
				reader.Read()
				if child, err := additive(reader); err == nil { //匹配加法表达式
					node.Append(child)
				} else if err == io.EOF { //不匹配
					return nil, errors.New("invalide variable initialization, expecting an expression")
				} else {
					return nil, err //语法错误
				}
			}
		} else {
			return nil, errors.New("variable name expected") //语法错误：找不到变量名
		}
		if token = reader.Peek(); token != nil && token.Type == lib.TokenType_SemiColon {
			reader.Read()
		} else {
			return nil, errors.New("invalid statement, expecting semicolon") //语法错误：找不到分号
		}
	}
	return nil, io.EOF
}

func expressionStatement(reader *TokenReader) (*lib.ASTNode, error) {
	return nil, io.EOF
}

func assignmentStatement(reader *TokenReader) (*lib.ASTNode, error) {
	return nil, io.EOF
}

/**
 * 加法表达式
 * addtive -> multiplicative ( (+ | -) multiplicative)*
 */
func additive(reader *TokenReader) (*lib.ASTNode, error) {
	return nil, io.EOF
}

/**
 * 乘法表达式
 * multiplicative -> primary ( (* | /) primary)*
 */
func multiplicative(reader *TokenReader) (*lib.ASTNode, error) {
	if token := reader.Peek(); token != nil {
		var node, child1 *lib.ASTNode
		var err error
		if child1, err = primary(reader); err != nil { //语法错误或不匹配
			return nil, err
		}
		node = child1
		for {
			if token := reader.Peek(); token.Type == lib.TokenType_Star || token.Type == lib.TokenType_Slash { // * 或者 /
				//该代码块中不匹配会被判定为语法错误
				reader.Read()
				if child2, err := primary(reader); err == nil {
					node = lib.NewASTNode(lib.ASTNodeType_Multiplicative, token.Text)
					node.Append(child1)
					node.Append(child2)
					child1 = node
				} else if err == io.EOF { //匹配失败
					return nil, errors.New("invalid multiplicative expression, expecting the right part")
				} else { //语法错误
					return nil, err
				}
			} else {
				break
			}
		}
		return node, nil
	}
	return nil, io.EOF
}

/**
 * 基础表达式
 * primary -> IntLiteral | Id | '(' additive ')'
 */
func primary(reader *TokenReader) (*lib.ASTNode, error) {
	if token := reader.Peek(); token != nil {
		if token.Type == lib.TokenType_IntLiteral { //IntLiteral
			reader.Read()
			return lib.NewASTNode(lib.ASTNodeType_IntLiteral, token.Text), nil
		} else if token.Type == lib.TokenType_Identifier { //Id
			reader.Read()
			return lib.NewASTNode(lib.ASTNodeType_Identifier, token.Text), nil
		} else if token = reader.Peek(); token.Type == lib.TokenType_LeftParen { // 左括号
			//该代码块中不匹配会被判定为语法错误
			reader.Read()
			if node, err := additive(reader); err == nil { //additive
				if token = reader.Peek(); token != nil && token.Type == lib.TokenType_RightParen { // 右括号
					reader.Read()
					return node, nil
				}
			} else if err == io.EOF { //不匹配加法表达式
				return nil, errors.New("expecting an additive expression inside parenthesis")
			} else { //语法错误
				return nil, err
			}
			return nil, errors.New("expecting right parenthesis") //没有找到右括号
		}
	}
	return nil, io.EOF
}
