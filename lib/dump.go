package lib

import (
	"bytes"
	"fmt"
	"strings"
)

//打印Token流
func (tokens *Tokens) Dump(print bool) string {
	var buffer bytes.Buffer
	for _, token := range []Token(*tokens) {
		buffer.WriteString(fmt.Sprintf("%-16s : %s\n", token.Type, token.Text))
	}
	if print {
		fmt.Print(buffer.String())
	}
	return buffer.String()
}

//打印抽象语法树
func (ast *ASTNode) Dump(print bool) string {
	var buffer bytes.Buffer
	ast.dump(&buffer, 0)
	if print {
		fmt.Print(buffer.String())
	}
	return buffer.String()
}

//抽象语法树打印实现
func (ast *ASTNode) dump(buffer *bytes.Buffer, indent int) {
	buffer.WriteString(fmt.Sprintf("%-16s : %s\n", strings.Repeat(" ", indent)+string(ast.Type), ast.Text))
	for _, child := range ast.Children {
		child.dump(buffer, indent+1)
	}
}
