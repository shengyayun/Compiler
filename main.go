package main

import (
	"Compiler/lexer"
	"Compiler/parser"
	"Compiler/runner"
	"bufio"
	"fmt"
	"os"
)

var lx lexer.Lexer
var ps parser.Parser
var rr runner.Runner

func init() {
	lx = lexer.NewLexer()
	ps = parser.NewParser()
	rr = runner.NewRunner()
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
	if len(tokens) > 0 {
		//语义分析
		fmt.Println("-------------", "语义分析", "-------------")
		if tree, err := ps.Parse(&tokens); err == nil {
			tree.Dump(true)
			//执行程序
			fmt.Println("-------------", "执行程序", "-------------")
			if ret, err := rr.Evaluate(tree); err == nil {
				fmt.Println(ret)
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	}
}
