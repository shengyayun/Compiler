package main

import (
	"compiler/lexer"
	"compiler/lib"
	"compiler/parser"
	"testing"
)

//测试词法分析
func TestTokenize(t *testing.T) {
	dict := make(map[string]lib.Tokens)

	dict["inta age = 45;"] = lib.Tokens([]lib.Token{lib.Token{Type: lib.TokenType_Identifier, Text: "inta"}, lib.Token{Type: lib.TokenType_Identifier, Text: "age"}, lib.Token{Type: lib.TokenType_Assignment, Text: "="}, lib.Token{Type: lib.TokenType_IntLiteral, Text: "45"}, lib.Token{Type: lib.TokenType_SemiColon, Text: ";"}})
	dict["int age = 45;"] = lib.Tokens([]lib.Token{lib.Token{Type: lib.TokenType_Int, Text: "int"}, lib.Token{Type: lib.TokenType_Identifier, Text: "age"}, lib.Token{Type: lib.TokenType_Assignment, Text: "="}, lib.Token{Type: lib.TokenType_IntLiteral, Text: "45"}, lib.Token{Type: lib.TokenType_SemiColon, Text: ";"}})
	dict["age >= 45;"] = lib.Tokens([]lib.Token{lib.Token{Type: lib.TokenType_Identifier, Text: "age"}, lib.Token{Type: lib.TokenType_GE, Text: ">="}, lib.Token{Type: lib.TokenType_IntLiteral, Text: "45"}, lib.Token{Type: lib.TokenType_SemiColon, Text: ";"}})
	dict["age > 45;"] = lib.Tokens([]lib.Token{lib.Token{Type: lib.TokenType_Identifier, Text: "age"}, lib.Token{Type: lib.TokenType_GT, Text: ">"}, lib.Token{Type: lib.TokenType_IntLiteral, Text: "45"}, lib.Token{Type: lib.TokenType_SemiColon, Text: ";"}})

	l := lexer.NewLexer()
	for code, expect := range dict {
		tokens := l.Tokenize(code)
		dump := tokens.Dump(false)
		t.Logf("\n%s\n---------------------------\n%s\n", code, dump)
		if dump != expect.Dump(false) {
			t.Errorf("'%s' tokenize fail", code)
		}
	}
}

//测试语义分析
func TestParse(t *testing.T) {
	dict := make(map[string]*lib.ASTNode)

	dict["int age = 45 + 2;"] = &lib.ASTNode{Type: lib.ASTNodeType_Programm, Text: "pwc", Children: []*lib.ASTNode{&lib.ASTNode{Type: lib.ASTNodeType_IntDeclaration, Text: "age", Children: []*lib.ASTNode{&lib.ASTNode{Type: lib.ASTNodeType_Additive, Text: "+", Children: []*lib.ASTNode{&lib.ASTNode{Type: lib.ASTNodeType_IntLiteral, Text: "45"}, &lib.ASTNode{Type: lib.ASTNodeType_IntLiteral, Text: "2"}}}}}}}
	dict["age = 20;"] = &lib.ASTNode{Type: lib.ASTNodeType_Programm, Text: "pwc", Children: []*lib.ASTNode{&lib.ASTNode{Type: lib.ASTNodeType_AssignmentStmt, Text: "age", Children: []*lib.ASTNode{&lib.ASTNode{Type: lib.ASTNodeType_IntLiteral, Text: "20"}}}}}
	dict["age + 10 * 2;"] = &lib.ASTNode{Type: lib.ASTNodeType_Programm, Text: "pwc", Children: []*lib.ASTNode{&lib.ASTNode{Type: lib.ASTNodeType_ExpressionStmt, Text: "", Children: []*lib.ASTNode{&lib.ASTNode{Type: lib.ASTNodeType_Additive, Text: "+", Children: []*lib.ASTNode{&lib.ASTNode{Type: lib.ASTNodeType_Identifier, Text: "age"}, &lib.ASTNode{Type: lib.ASTNodeType_Multiplicative, Text: "*", Children: []*lib.ASTNode{&lib.ASTNode{Type: lib.ASTNodeType_IntLiteral, Text: "10"}, &lib.ASTNode{Type: lib.ASTNodeType_IntLiteral, Text: "2"}}}}}}}}}

	l := lexer.NewLexer()
	p := parser.NewParser()

	for k, v := range dict {
		tokens := l.Tokenize(k)
		if ast, err := p.Parse(&tokens); err == nil {
			dump := ast.Dump(false)
			t.Logf("\n%s\n---------------------------\n%s\n", k, dump)
			if dump != v.Dump(false) {
				t.Errorf("'%s' tokenize fail", k)
			}
		} else {
			t.Errorf("'%s' parse fail", k)
		}
	}
}
