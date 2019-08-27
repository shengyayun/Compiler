package main

import "bytes"

type Token struct {
	Type TokenType
	text bytes.Buffer
}

//在文本后添加字符
func (token *Token) Append(ch rune) {
	token.text.WriteRune(ch)
}

//获取token的文本
func (token *Token) Text() string {
	return token.text.String()
}

type TokenType uint8

const (
	TokenType_Plus  TokenType = iota + 1 // +
	TokenType_Minus                      // -
	TokenType_Star                       // *
	TokenType_Slash                      // /

	TokenType_GE // >=
	TokenType_GT // >
	TokenType_EQ // ==
	TokenType_LE // <=
	TokenType_LT // <

	TokenType_SemiColon  // ;
	TokenType_LeftParen  // (
	TokenType_RightParen // )

	TokenType_Assignment // =

	TokenType_If
	TokenType_Else

	TokenType_Int

	TokenType_Identifier //标识符

	TokenType_IntLiteral    //整型字面量
	TokenType_StringLiteral //字符串字面量
)
