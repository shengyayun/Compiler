package main

import (
	"io"
	"strings"
)

type Lexer struct {
	state  DfaState //Dfa状态
	staged Token    //暂存区
	tokens []Token  //token列表
}

func NewLexer() (lexer Lexer) {
	lexer.state = DfaState_Initial
	lexer.tokens = make([]Token, 0)
	return
}

//通过首字符初始化Token与Dfa状态
func (lexer *Lexer) Checkout(ch rune) {
	if isAlpha(ch) { //字母
		if ch == 'i' {
			lexer.state = DfaState_Id_int1
		} else {
			lexer.state = DfaState_Id //进入Id状态
		}
		lexer.staged.Type = TokenType_Identifier
		lexer.Add(ch)
	} else if isDigit(ch) { //数字
		lexer.state = DfaState_IntLiteral
		lexer.staged.Type = TokenType_IntLiteral
		lexer.Add(ch)
	} else if ch == '>' {
		lexer.state = DfaState_GT
		lexer.staged.Type = TokenType_GT
		lexer.Add(ch)
	} else if ch == '+' {
		lexer.state = DfaState_Plus
		lexer.staged.Type = TokenType_Plus
		lexer.Add(ch)
	} else if ch == '-' {
		lexer.state = DfaState_Minus
		lexer.staged.Type = TokenType_Minus
		lexer.Add(ch)
	} else if ch == '*' {
		lexer.state = DfaState_Star
		lexer.staged.Type = TokenType_Star
		lexer.Add(ch)
	} else if ch == '/' {
		lexer.state = DfaState_Slash
		lexer.staged.Type = TokenType_Slash
		lexer.Add(ch)
	} else if ch == ';' {
		lexer.state = DfaState_SemiColon
		lexer.staged.Type = TokenType_SemiColon
		lexer.Add(ch)
	} else if ch == '(' {
		lexer.state = DfaState_LeftParen
		lexer.staged.Type = TokenType_LeftParen
		lexer.Add(ch)
	} else if ch == ')' {
		lexer.state = DfaState_RightParen
		lexer.staged.Type = TokenType_RightParen
		lexer.Add(ch)
	} else if ch == '=' {
		lexer.state = DfaState_Assignment
		lexer.staged.Type = TokenType_Assignment
		lexer.Add(ch)
	} else {
		lexer.state = DfaState_Initial
	}
}

//当前token的Text内容填充
func (lexer *Lexer) Add(ch rune) {
	lexer.staged.Append(ch)
}

//提交当前token
func (lexer *Lexer) Commit() {
	lexer.tokens = append(lexer.tokens, lexer.staged)
	lexer.staged = Token{}
}

//开始词法分析
func (lexer *Lexer) Tokenize(code string) []Token {
	var ch rune
	var err error
	reader := strings.NewReader(code)
	for {
		if ch, _, err = reader.ReadRune(); err == io.EOF {
			break
		}
		switch lexer.state {
		case DfaState_Initial: //初始状态
			if lexer.staged.Text() != "" { //保存历史token
				lexer.Commit()
			}
			lexer.Checkout(ch)
		case DfaState_Id: //标识名
			if isAlpha(ch) || isDigit(ch) {
				lexer.Add(ch) //保持标识符状态
			} else {
				lexer.Commit()
				lexer.Checkout(ch)
			}
		case DfaState_GT: //>
			if ch == '=' { //>=
				lexer.staged.Type = TokenType_GE
				lexer.state = DfaState_GE
				lexer.Add(ch)
			} else {
				lexer.Commit()
				lexer.Checkout(ch)
			}
		case DfaState_GE: //>=
		case DfaState_Assignment: //=
		case DfaState_Plus: //+
		case DfaState_Minus: //-
		case DfaState_Star: //*
		case DfaState_Slash: //\
		case DfaState_SemiColon: //;
		case DfaState_LeftParen: //(
		case DfaState_RightParen: //)
			lexer.Commit()
			lexer.Checkout(ch)
		case DfaState_IntLiteral: //数字
			if isDigit(ch) {
				lexer.Add(ch)
			} else {
				lexer.Commit()
				lexer.Checkout(ch)
			}
		case DfaState_Id_int1: //i
			if ch == 'n' { //in
				lexer.state = DfaState_Id_int2
				lexer.Add(ch)
			} else if isDigit(ch) || isAlpha(ch) {
				lexer.state = DfaState_Id //切换回Id状态
				lexer.Add(ch)
			} else {
				lexer.Commit()
				lexer.Checkout(ch)
			}
		case DfaState_Id_int2:
			if ch == 't' { //int
				lexer.state = DfaState_Id_int3
				lexer.Add(ch)
			} else if isDigit(ch) || isAlpha(ch) {
				lexer.state = DfaState_Id //切换回id状态
				lexer.Add(ch)
			} else {
				lexer.Commit()
				lexer.Checkout(ch)
			}
		case DfaState_Id_int3: //int
			if isBlank(ch) { //"int"后是空格
				lexer.staged.Type = TokenType_Int
				lexer.Commit()
				lexer.Checkout(ch)
			} else {
				lexer.state = DfaState_Id //切换回Id状态
				lexer.Add(ch)
			}
		}
	}
	if lexer.staged.Text() != "" { //保存历史token
		lexer.Commit()
	}
	return lexer.tokens
}
