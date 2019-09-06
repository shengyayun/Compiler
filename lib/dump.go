package lib

import (
	"fmt"
	"strings"
)

//打印Token流
func (tokens *Tokens) Dump() {
	split("tokens")
	for _, token := range []Token(*tokens) {
		fmt.Printf("%-16s : %s\n", token.Type, token.Text)
	}
}

//打印抽象语法树
func (ast *ASTNode) Dump() {
	split("ast")
	ast.dump(0)
}

//打印分割线
func split(title string) {
	fmt.Println(strings.Repeat("-", 12), title, strings.Repeat("-", 12))
}

//抽象语法树打印实现
func (ast *ASTNode) dump(indent int) {
	fmt.Printf("%-16s : %s\n", strings.Repeat(" ", indent)+string(ast.Type), ast.Text)
	for _, child := range ast.Children {
		child.dump(indent + 1)
	}
}
