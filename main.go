package main

import (
	"Compiler/lexer"
	"Compiler/parser"
	"bufio"
	"fmt"
	"os"
)

var lx lexer.Lexer
var ps parser.Parser

func init() {
	lx = lexer.NewLexer()
	ps = parser.NewParser()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for scanner.Scan() {
		script := scanner.Text()
		compile(script)
		fmt.Print("\n> ")
	}
}

func compile(script string) {
	//词法分析
	fmt.Println("-------------", "词法分析", "-------------")
	tokens := lx.Tokenize(script)
	tokens.Dump(true)

	//语义分析
	fmt.Println("-------------", "语义分析", "-------------")
	if tree, err := ps.Parse(&tokens); err == nil {
		tree.Dump(true)
	} else {
		fmt.Println(err)
	}
}
