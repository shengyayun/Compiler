//语义分析
package parser

import (
	"Compiler/lib"
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
 * expressionStatement -> additive ';'
 * assignmentStatement -> id = additive ';'
 * additive -> multiplicative ( (+ | -) multiplicative)*
 * multiplicative -> primary ( (* | /) primary)*
 * primary -> IntLiteral | Id | '(' additive ')'
 */
type Parser struct{}

func NewParser() Parser {
	return Parser{}
}

/**
 * 将token流解析为抽象语法树
 *
 * programm -> intDeclare | expressionStatement | assignmentStatement
 */
func (parse *Parser) Parse(tokens *lib.Tokens) (*lib.ASTNode, error) {
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
		child = nil
		for _, method := range handler {
			if child, err = method(&reader); err == nil {
				break
			} else if err != io.EOF {
				return nil, err
			}
		}
		if child != nil {
			tree.Append(child)
		} else {
			return nil, errors.New("未知的语法")
		}
	}
	return &tree, nil
}

/**
 * 整型变量声明语句
 *
 * intDeclare -> 'int' Id ( = additive) ';'
 */
func intDeclare(reader *TokenReader) (*lib.ASTNode, error) {
	if token := reader.Peek(); token.Type == lib.TokenType_Int { //int
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
					return nil, errors.New("语法错误：整形变量声明中缺少表达式语句")
				} else {
					return nil, err //语法错误
				}
			}
			if token = reader.Peek(); token != nil && token.Type == lib.TokenType_SemiColon { // ;
				reader.Read()
				return node, nil
			}
			return nil, errors.New("语法错误：整形变量声明中缺少分号") //语法错误：找不到分号
		}
		return nil, errors.New("语法错误：整形变量声明中缺少变量名") //语法错误：找不到变量名
	}
	return nil, io.EOF
}

/**
 * 表达式语句
 *
 * expressionStatement -> additive ';'
 */
func expressionStatement(reader *TokenReader) (*lib.ASTNode, error) {
	pos := reader.Position()
	if child, err := additive(reader); err == nil {
		if token := reader.Peek(); token != nil && token.Type == lib.TokenType_SemiColon {
			reader.Read()
			node := lib.NewASTNode(lib.ASTNodeType_ExpressionStmt, "")
			node.Append(child)
			return node, nil
		} //回溯
		reader.SetPosition(pos)
	}
	return nil, io.EOF
}

/**
 * 赋值语句
 *
 * assignmentStatement -> id = additive ';'
 */
func assignmentStatement(reader *TokenReader) (*lib.ASTNode, error) {
	if token := reader.Peek(); token.Type == lib.TokenType_Identifier { //id
		node := lib.NewASTNode(lib.ASTNodeType_AssignmentStmt, token.Text)
		reader.Read()
		if token := reader.Peek(); token != nil && token.Type == lib.TokenType_Assignment { // =
			reader.Read()
			if child, err := additive(reader); err == nil { //additive
				node.Append(child)
				if token := reader.Peek(); token != nil && token.Type == lib.TokenType_SemiColon { //;
					reader.Read()
					return node, nil
				}
				return nil, errors.New("语法错误：赋值语句中缺少分号")
			}
			return nil, errors.New("语法错误：赋值语句中缺少表达式语句")
		}
		reader.Unread()
	}
	return nil, io.EOF
}

/**
 * 加法表达式
 * additive -> multiplicative ( (+ | -) multiplicative)*
 */
func additive(reader *TokenReader) (*lib.ASTNode, error) {
	var node, child1 *lib.ASTNode
	var err error
	if child1, err = multiplicative(reader); err != nil {
		return nil, err
	}
	node = child1
	for {
		if token := reader.Peek(); token != nil && (token.Type == lib.TokenType_Plus || token.Type == lib.TokenType_Minus) { // + 或者 -
			//该代码块中不匹配会被判定为语法错误
			reader.Read()
			if child2, err := multiplicative(reader); err == nil {
				node = lib.NewASTNode(lib.ASTNodeType_Additive, token.Text)
				node.Append(child1)
				node.Append(child2)
				child1 = node
			} else if err == io.EOF { //匹配失败
				return nil, errors.New("语法错误：加法表达式中缺少右部分")
			} else { //语法错误
				return nil, err
			}
		} else {
			break
		}
	}
	return node, nil
}

/**
 * 乘法表达式
 * multiplicative -> primary ( (* | /) primary)*
 */
func multiplicative(reader *TokenReader) (*lib.ASTNode, error) {
	var node, child1 *lib.ASTNode
	var err error
	if child1, err = primary(reader); err != nil { //语法错误或不匹配
		return nil, err
	}
	node = child1
	for {
		if token := reader.Peek(); token != nil && (token.Type == lib.TokenType_Star || token.Type == lib.TokenType_Slash) { // * 或者 /
			//该代码块中不匹配会被判定为语法错误
			reader.Read()
			if child2, err := primary(reader); err == nil {
				node = lib.NewASTNode(lib.ASTNodeType_Multiplicative, token.Text)
				node.Append(child1)
				node.Append(child2)
				child1 = node
			} else if err == io.EOF { //匹配失败
				return nil, errors.New("语法错误：乘法表达式中缺少右部分")
			} else { //语法错误
				return nil, err
			}
		} else {
			break
		}
	}
	return node, nil
}

/**
 * 基础表达式（为了简化语法树，直接返回子节点）
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
				return nil, errors.New("语法错误：基础表达式中括号内不是一个加法表达式")
			} else { //语法错误
				return nil, err
			}
			return nil, errors.New("语法错误：基础表达式中缺少右括号") //没有找到右括号
		}
	}
	return nil, io.EOF
}
