package main

import "fmt"

func main() {
	lexer := NewLexer()
	tokens := lexer.Tokenize("int age = 45;")
	for i := 0; i < len(tokens); i++ {
		fmt.Println(string(tokens[i].Type) + " :  " + tokens[i].Text)
	}
}
