package main

import (
	"io"
	"strings"
)

type Lexer struct {
	state  DfaState
	staged Token   //暂存区
	tokens []Token //token列表
}

func NewLexer() (lexer Lexer) {
	lexer.state = DfaState_Initial
	lexer.tokens = make([]Token, 0)
	return
}

func (lexer *Lexer) Next(ch rune) {
	if isAlpha(ch) {
		if ch == 'i' {
			lexer.state = DfaState_Id_int1
		} else {
			lexer.state = DfaState_Id //进入Id状态
		}
		lexer.staged.Type = TokenType_Identifier
		lexer.staged.Append(ch)
	} else if isDigit(ch) { //第一个字符是数字
		lexer.state = DfaState_IntLiteral
		lexer.staged.Type = TokenType_IntLiteral
		lexer.staged.Append(ch)
	} else if ch == '>' { //第一个字符是>
		lexer.state = DfaState_GT
		lexer.staged.Type = TokenType_GT
		lexer.staged.Append(ch)
	} else if ch == '+' {
		lexer.state = DfaState_Plus
		lexer.staged.Type = TokenType_Plus
		lexer.staged.Append(ch)
	} else if ch == '-' {
		lexer.state = DfaState_Minus
		lexer.staged.Type = TokenType_Minus
		lexer.staged.Append(ch)
	} else if ch == '*' {
		lexer.state = DfaState_Star
		lexer.staged.Type = TokenType_Star
		lexer.staged.Append(ch)
	} else if ch == '/' {
		lexer.state = DfaState_Slash
		lexer.staged.Type = TokenType_Slash
		lexer.staged.Append(ch)
	} else if ch == ';' {
		lexer.state = DfaState_SemiColon
		lexer.staged.Type = TokenType_SemiColon
		lexer.staged.Append(ch)
	} else if ch == '(' {
		lexer.state = DfaState_LeftParen
		lexer.staged.Type = TokenType_LeftParen
		lexer.staged.Append(ch)
	} else if ch == ')' {
		lexer.state = DfaState_RightParen
		lexer.staged.Type = TokenType_RightParen
		lexer.staged.Append(ch)
	} else if ch == '=' {
		lexer.state = DfaState_Assignment
		lexer.staged.Type = TokenType_Assignment
		lexer.staged.Append(ch)
	} else {
		lexer.state = DfaState_Initial
	}
}

func (lexer *Lexer) Add(ch rune) {
	lexer.staged.Append(ch)
}

func (lexer *Lexer) Commit() {
	lexer.tokens = append(lexer.tokens, lexer.staged)
	lexer.staged = Token{}
}

func (lexer *Lexer) Tokenize(code string) []Token {
	var ch rune
	var err error
	reader := strings.NewReader(code)
	for {
		if ch, _, err = reader.ReadRune(); err == io.EOF {
			break
		}
		switch lexer.state {
		case DfaState_Initial:
			if lexer.staged.Text() != "" {
				lexer.Commit()
			}
			lexer.Next(ch)
		case DfaState_Id:
			if isAlpha(ch) || isDigit(ch) {
				lexer.Add(ch) //保持标识符状态
			} else {
				lexer.Commit()
				lexer.Next(ch)
			}
			break
		}
	}
	return lexer.tokens
}
