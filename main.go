package main

import (
	"craft/lexer"
	"craft/parser"
	"fmt"
)

func main() {
	//代码
	//script := "int age = 45+2; age= 20; age+10*2;"
	script := "int age = 45;"

	//词法分析
	l := lexer.NewLexer()
	tokens := l.Tokenize(script)
	fmt.Println(tokens)

	//语义分析
	p := parser.NewParser()
	tree, err := p.Parse(&tokens)
	fmt.Println(tree, err)
}
