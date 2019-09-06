package main

import (
	"compiler/lexer"
	"compiler/parser"
	"fmt"
)

func main() {
	//代码
	script := "int age = 45+2; age= 20; age+10*2;"

	//词法分析
	l := lexer.NewLexer()
	tokens := l.Tokenize(script)
	tokens.Dump()

	//语义分析
	p := parser.NewParser()
	if tree, err := p.Parse(&tokens); err == nil {
		tree.Dump()
	} else {
		fmt.Println("ex: ", err)
	}
}
