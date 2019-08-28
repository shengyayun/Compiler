package main

import (
	"fmt"
	"testing"
)

func TestTokenize(t *testing.T) {
	dict := make(map[string][]Token)

	dict["inta age = 45;"] = []Token{Token{TokenType_Identifier, "inta"}, Token{TokenType_Identifier, "age"}, Token{TokenType_Assignment, "="}, Token{TokenType_IntLiteral, "45"}, Token{TokenType_SemiColon, ";"}}
	dict["int age = 45;"] = []Token{Token{TokenType_Int, "int"}, Token{TokenType_Identifier, "age"}, Token{TokenType_Assignment, "="}, Token{TokenType_IntLiteral, "45"}, Token{TokenType_SemiColon, ";"}}
	dict["age >= 45;"] = []Token{Token{TokenType_Identifier, "age"}, Token{TokenType_GE, ">="}, Token{TokenType_IntLiteral, "45"}, Token{TokenType_SemiColon, ";"}}
	dict["age > 45;"] = []Token{Token{TokenType_Identifier, "age"}, Token{TokenType_GT, ">"}, Token{TokenType_IntLiteral, "45"}, Token{TokenType_SemiColon, ";"}}

	lexer := NewLexer()
	for code, expect := range dict {

		tokens := lexer.Tokenize(code)
		if len(tokens) != len(expect) {
			t.Errorf("'%s' tokenize fail", code)
		}
		for i := 0; i < len(tokens); i++ {
			if tokens[i].Type != expect[i].Type || tokens[i].Text != expect[i].Text {
				fmt.Println(tokens, expect)
				t.Errorf("'%s' tokenize fail", code)
			}
		}
		t.Logf("'%s' tokenize success", code)
	}
}
