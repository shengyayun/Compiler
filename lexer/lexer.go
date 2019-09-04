//词法分析
package lexer

import (
	"bytes"
	"craft/lib"
	"io"
	"strings"
)

type Lexer struct {
	dfaState DfaState //Dfa状态

	tokenType   lib.TokenType //Token类型
	tokenBuffer bytes.Buffer  //Token内容

	result []lib.Token //token结果列表
}

func NewLexer() (lexer Lexer) {
	lexer.dfaState = DfaState_Initial
	lexer.result = make([]lib.Token, 0)
	return
}

func (lexer *Lexer) Reset() {
	lexer.dfaState = DfaState_Initial
	lexer.tokenType = ""
	lexer.tokenBuffer.Reset()
	lexer.result = lexer.result[0:0]
}

//通过首字符初始化Token与Dfa状态
func (lexer *Lexer) Checkout(ch rune) {
	if lexer.tokenBuffer.String() != "" { //保存历史token
		lexer.Commit()
	}
	if isAlpha(ch) { //字母
		if ch == 'i' {
			lexer.dfaState = DfaState_Id_int1
		} else {
			lexer.dfaState = DfaState_Id //进入Id状态
		}
		lexer.tokenType = lib.TokenType_Identifier
		lexer.tokenBuffer.WriteRune(ch)
	} else if isDigit(ch) { //数字
		lexer.dfaState = DfaState_IntLiteral
		lexer.tokenType = lib.TokenType_IntLiteral
		lexer.tokenBuffer.WriteRune(ch)
	} else if ch == '>' {
		lexer.dfaState = DfaState_GT
		lexer.tokenType = lib.TokenType_GT
		lexer.tokenBuffer.WriteRune(ch)
	} else if ch == '+' {
		lexer.dfaState = DfaState_Plus
		lexer.tokenType = lib.TokenType_Plus
		lexer.tokenBuffer.WriteRune(ch)
	} else if ch == '-' {
		lexer.dfaState = DfaState_Minus
		lexer.tokenType = lib.TokenType_Minus
		lexer.tokenBuffer.WriteRune(ch)
	} else if ch == '*' {
		lexer.dfaState = DfaState_Star
		lexer.tokenType = lib.TokenType_Star
		lexer.tokenBuffer.WriteRune(ch)
	} else if ch == '/' {
		lexer.dfaState = DfaState_Slash
		lexer.tokenType = lib.TokenType_Slash
		lexer.tokenBuffer.WriteRune(ch)
	} else if ch == ';' {
		lexer.dfaState = DfaState_SemiColon
		lexer.tokenType = lib.TokenType_SemiColon
		lexer.tokenBuffer.WriteRune(ch)
	} else if ch == '(' {
		lexer.dfaState = DfaState_LeftParen
		lexer.tokenType = lib.TokenType_LeftParen
		lexer.tokenBuffer.WriteRune(ch)
	} else if ch == ')' {
		lexer.dfaState = DfaState_RightParen
		lexer.tokenType = lib.TokenType_RightParen
		lexer.tokenBuffer.WriteRune(ch)
	} else if ch == '=' {
		lexer.dfaState = DfaState_Assignment
		lexer.tokenType = lib.TokenType_Assignment
		lexer.tokenBuffer.WriteRune(ch)
	} else {
		lexer.dfaState = DfaState_Initial
	}
}

//提交当前token
func (lexer *Lexer) Commit() {
	lexer.result = append(lexer.result, lib.Token{Type: lexer.tokenType, Text: lexer.tokenBuffer.String()})

	lexer.tokenType = ""
	lexer.tokenBuffer.Reset()
}

//开始词法分析
func (lexer *Lexer) Tokenize(code string) []lib.Token {
	lexer.Reset()
	var ch rune
	var err error
	reader := strings.NewReader(code)
	for {
		if ch, _, err = reader.ReadRune(); err == io.EOF {
			break
		}
		switch lexer.dfaState {
		case DfaState_Initial: //初始状态
			lexer.Checkout(ch)
		case DfaState_Id: //标识名
			if isAlpha(ch) || isDigit(ch) {
				lexer.tokenBuffer.WriteRune(ch) //保持标识符状态
			} else {
				lexer.Checkout(ch)
			}
		case DfaState_GT: //>
			if ch == '=' { //>=
				lexer.tokenType = lib.TokenType_GE
				lexer.dfaState = DfaState_GE
				lexer.tokenBuffer.WriteRune(ch)
			} else {
				lexer.Checkout(ch)
			}
		case DfaState_IntLiteral: //数字
			if isDigit(ch) {
				lexer.tokenBuffer.WriteRune(ch)
			} else {
				lexer.Checkout(ch)
			}
		case DfaState_Id_int1: //i
			if ch == 'n' { //in
				lexer.dfaState = DfaState_Id_int2
				lexer.tokenBuffer.WriteRune(ch)
			} else if isDigit(ch) || isAlpha(ch) {
				lexer.dfaState = DfaState_Id //切换回Id状态
				lexer.tokenBuffer.WriteRune(ch)
			} else {
				lexer.Checkout(ch)
			}
		case DfaState_Id_int2:
			if ch == 't' { //int
				lexer.dfaState = DfaState_Id_int3
				lexer.tokenBuffer.WriteRune(ch)
			} else if isDigit(ch) || isAlpha(ch) {
				lexer.dfaState = DfaState_Id //切换回id状态
				lexer.tokenBuffer.WriteRune(ch)
			} else {
				lexer.Checkout(ch)
			}
		case DfaState_Id_int3: //int
			if isBlank(ch) { //"int"后是空格
				lexer.tokenType = lib.TokenType_Int
				lexer.Checkout(ch)
			} else {
				lexer.dfaState = DfaState_Id //切换回Id状态
				lexer.tokenBuffer.WriteRune(ch)
			}
		case DfaState_GE, DfaState_Assignment, DfaState_Plus, DfaState_Minus, DfaState_Star, DfaState_Slash, DfaState_SemiColon, DfaState_LeftParen, DfaState_RightParen:
			lexer.Checkout(ch)
		}
	}
	if lexer.tokenBuffer.String() != "" { //保存历史token
		lexer.Commit()
	}
	return lexer.result
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
