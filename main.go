package main

import "fmt"

func main() {
	script := "int age = 45+2; age= 20; age+10*2;"
	parser := NewParser()
	tree := parser.parse(script)
	fmt.Println(tree)
}
