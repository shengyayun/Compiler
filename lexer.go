//词法分析
package main

import (
	"bytes"
	"io"
	"strings"
)

type Lexer struct {
	DfaState DfaState //Dfa状态

	TokenType   TokenType    //Token类型
	TokenBuffer bytes.Buffer //Token内容

	Result []Token //token结果列表
}

func NewLexer() (lexer Lexer) {
	lexer.DfaState = DfaState_Initial
	lexer.Result = make([]Token, 0)
	return
}

func (lexer *Lexer) Reset() {
	lexer.DfaState = DfaState_Initial
	lexer.TokenType = ""
	lexer.TokenBuffer.Reset()
	lexer.Result = lexer.Result[0:0]
}

//通过首字符初始化Token与Dfa状态
func (lexer *Lexer) Checkout(ch rune) {
	if lexer.TokenBuffer.String() != "" { //保存历史token
		lexer.Commit()
	}
	if isAlpha(ch) { //字母
		if ch == 'i' {
			lexer.DfaState = DfaState_Id_int1
		} else {
			lexer.DfaState = DfaState_Id //进入Id状态
		}
		lexer.TokenType = TokenType_Identifier
		lexer.TokenBuffer.WriteRune(ch)
	} else if isDigit(ch) { //数字
		lexer.DfaState = DfaState_IntLiteral
		lexer.TokenType = TokenType_IntLiteral
		lexer.TokenBuffer.WriteRune(ch)
	} else if ch == '>' {
		lexer.DfaState = DfaState_GT
		lexer.TokenType = TokenType_GT
		lexer.TokenBuffer.WriteRune(ch)
	} else if ch == '+' {
		lexer.DfaState = DfaState_Plus
		lexer.TokenType = TokenType_Plus
		lexer.TokenBuffer.WriteRune(ch)
	} else if ch == '-' {
		lexer.DfaState = DfaState_Minus
		lexer.TokenType = TokenType_Minus
		lexer.TokenBuffer.WriteRune(ch)
	} else if ch == '*' {
		lexer.DfaState = DfaState_Star
		lexer.TokenType = TokenType_Star
		lexer.TokenBuffer.WriteRune(ch)
	} else if ch == '/' {
		lexer.DfaState = DfaState_Slash
		lexer.TokenType = TokenType_Slash
		lexer.TokenBuffer.WriteRune(ch)
	} else if ch == ';' {
		lexer.DfaState = DfaState_SemiColon
		lexer.TokenType = TokenType_SemiColon
		lexer.TokenBuffer.WriteRune(ch)
	} else if ch == '(' {
		lexer.DfaState = DfaState_LeftParen
		lexer.TokenType = TokenType_LeftParen
		lexer.TokenBuffer.WriteRune(ch)
	} else if ch == ')' {
		lexer.DfaState = DfaState_RightParen
		lexer.TokenType = TokenType_RightParen
		lexer.TokenBuffer.WriteRune(ch)
	} else if ch == '=' {
		lexer.DfaState = DfaState_Assignment
		lexer.TokenType = TokenType_Assignment
		lexer.TokenBuffer.WriteRune(ch)
	} else {
		lexer.DfaState = DfaState_Initial
	}
}

//提交当前token
func (lexer *Lexer) Commit() {
	lexer.Result = append(lexer.Result, Token{lexer.TokenType, lexer.TokenBuffer.String()})

	lexer.TokenType = ""
	lexer.TokenBuffer.Reset()
}

//开始词法分析
func (lexer *Lexer) Tokenize(code string) []Token {
	lexer.Reset()
	var ch rune
	var err error
	reader := strings.NewReader(code)
	for {
		if ch, _, err = reader.ReadRune(); err == io.EOF {
			break
		}
		switch lexer.DfaState {
		case DfaState_Initial: //初始状态
			lexer.Checkout(ch)
		case DfaState_Id: //标识名
			if isAlpha(ch) || isDigit(ch) {
				lexer.TokenBuffer.WriteRune(ch) //保持标识符状态
			} else {
				lexer.Checkout(ch)
			}
		case DfaState_GT: //>
			if ch == '=' { //>=
				lexer.TokenType = TokenType_GE
				lexer.DfaState = DfaState_GE
				lexer.TokenBuffer.WriteRune(ch)
			} else {
				lexer.Checkout(ch)
			}
		case DfaState_IntLiteral: //数字
			if isDigit(ch) {
				lexer.TokenBuffer.WriteRune(ch)
			} else {
				lexer.Checkout(ch)
			}
		case DfaState_Id_int1: //i
			if ch == 'n' { //in
				lexer.DfaState = DfaState_Id_int2
				lexer.TokenBuffer.WriteRune(ch)
			} else if isDigit(ch) || isAlpha(ch) {
				lexer.DfaState = DfaState_Id //切换回Id状态
				lexer.TokenBuffer.WriteRune(ch)
			} else {
				lexer.Checkout(ch)
			}
		case DfaState_Id_int2:
			if ch == 't' { //int
				lexer.DfaState = DfaState_Id_int3
				lexer.TokenBuffer.WriteRune(ch)
			} else if isDigit(ch) || isAlpha(ch) {
				lexer.DfaState = DfaState_Id //切换回id状态
				lexer.TokenBuffer.WriteRune(ch)
			} else {
				lexer.Checkout(ch)
			}
		case DfaState_Id_int3: //int
			if isBlank(ch) { //"int"后是空格
				lexer.TokenType = TokenType_Int
				lexer.Checkout(ch)
			} else {
				lexer.DfaState = DfaState_Id //切换回Id状态
				lexer.TokenBuffer.WriteRune(ch)
			}
		case DfaState_GE, DfaState_Assignment, DfaState_Plus, DfaState_Minus, DfaState_Star, DfaState_Slash, DfaState_SemiColon, DfaState_LeftParen, DfaState_RightParen:
			lexer.Checkout(ch)
		}
	}
	if lexer.TokenBuffer.String() != "" { //保存历史token
		lexer.Commit()
	}
	return lexer.Result
}

//是否是字母
func isAlpha(ch rune) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z'
}

//是否是数字
func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

//是否是空白字符
func isBlank(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}
