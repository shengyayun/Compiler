//Token相关
package main

import "errors"

type Token struct {
	Type TokenType
	Text string
}

type TokenType string

const (
	TokenType_Plus  TokenType = "Plus"  // +
	TokenType_Minus           = "Minus" // -
	TokenType_Star            = "Star"  // *
	TokenType_Slash           = "Slash" // /

	TokenType_GE = "GE" // >=
	TokenType_GT = "GT" // >
	TokenType_EQ = "EQ" // ==
	TokenType_LE = "LE" // <=
	TokenType_LT = "LT" // <

	TokenType_SemiColon  = "SemiColon"  // ;
	TokenType_LeftParen  = "LeftParen"  // (
	TokenType_RightParen = "RightParen" // )

	TokenType_Assignment = "Assignment" // =

	TokenType_If   = "If"
	TokenType_Else = "Else"

	TokenType_Int = "Int"

	TokenType_Identifier = "Identifier" //标识符

	TokenType_IntLiteral    = "IntLiteral"    //整型字面量
	TokenType_StringLiteral = "StringLiteral" //字符串字面量
)

type TokenReader struct {
	tokens []*Token
	pos    int
}

func NewTokenReader(tokens []*Token) TokenReader {
	return TokenReader{tokens: tokens, pos: 0}
}

//读取下一个token，并移动游标
func (reader *TokenReader) Read() *Token {
	token := reader.Peek()
	reader.pos++
	return token
}

//读取下一个token，但不移动游标
func (reader *TokenReader) Peek() *Token {
	if reader.pos > len(reader.tokens)-1 {
		return nil
	}
	return reader.tokens[reader.pos]
}

//游标回退一步
func (reader *TokenReader) Unread() error {
	if reader.pos > 0 {
		reader.pos--
		return nil
	}
	return errors.New("Position Out of Range")
}

//返回当前游标位置
func (reader *TokenReader) Position() int {
	return reader.pos
}

//设置游标位置
func (reader *TokenReader) SetPosition(pos int) error {
	if pos >= 0 {
		reader.pos = pos
		return nil
	}
	return errors.New("Position Out of Range")
}
