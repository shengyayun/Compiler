package parser

import (
	"craft/lib"
	"errors"
)

type TokenReader struct {
	tokens *([]lib.Token)
	pos    int
}

func NewTokenReader(tokens *([]lib.Token)) TokenReader {
	return TokenReader{tokens: tokens, pos: 0}
}

//读取下一个token，并移动游标
func (reader *TokenReader) Read() (token *lib.Token) {
	token = reader.Peek()
	reader.pos++
	return
}

//读取下一个token，但不移动游标
func (reader *TokenReader) Peek() *lib.Token {
	if reader.pos > len(*reader.tokens)-1 {
		return nil
	}
	return &(*reader.tokens)[reader.pos]
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
