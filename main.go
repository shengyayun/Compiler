package main

import (
	"fmt"
)

func main() {
	lexer := NewLexer()
	tokens := lexer.Tokenize("int age = 45;")
	fmt.Println(tokens)
}
