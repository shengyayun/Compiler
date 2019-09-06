package main

import (
	"compiler/lexer"
	"compiler/lib"
	"testing"
)

func TestTokenize(t *testing.T) {
	dict := make(map[string][]lib.Token)

	dict["inta age = 45;"] = []lib.Token{lib.Token{Type: lib.TokenType_Identifier, Text: "inta"}, lib.Token{Type: lib.TokenType_Identifier, Text: "age"}, lib.Token{Type: lib.TokenType_Assignment, Text: "="}, lib.Token{Type: lib.TokenType_IntLiteral, Text: "45"}, lib.Token{Type: lib.TokenType_SemiColon, Text: ";"}}
	dict["int age = 45;"] = []lib.Token{lib.Token{Type: lib.TokenType_Int, Text: "int"}, lib.Token{Type: lib.TokenType_Identifier, Text: "age"}, lib.Token{Type: lib.TokenType_Assignment, Text: "="}, lib.Token{Type: lib.TokenType_IntLiteral, Text: "45"}, lib.Token{Type: lib.TokenType_SemiColon, Text: ";"}}
	dict["age >= 45;"] = []lib.Token{lib.Token{Type: lib.TokenType_Identifier, Text: "age"}, lib.Token{Type: lib.TokenType_GE, Text: ">="}, lib.Token{Type: lib.TokenType_IntLiteral, Text: "45"}, lib.Token{Type: lib.TokenType_SemiColon, Text: ";"}}
	dict["age > 45;"] = []lib.Token{lib.Token{Type: lib.TokenType_Identifier, Text: "age"}, lib.Token{Type: lib.TokenType_GT, Text: ">"}, lib.Token{Type: lib.TokenType_IntLiteral, Text: "45"}, lib.Token{Type: lib.TokenType_SemiColon, Text: ";"}}

	l := lexer.NewLexer()
	for code, expect := range dict {
		tokens := l.Tokenize(code)
		if len(tokens) != len(expect) {
			t.Errorf("'%s' tokenize fail", code)
		}
		for i := 0; i < len(tokens); i++ {
			if tokens[i].Type != expect[i].Type || tokens[i].Text != expect[i].Text {
				t.Errorf("'%s' tokenize fail", code)
			}
		}
		t.Logf("'%s' tokenize success", code)
	}
}
