package runner

import (
	"Compiler/lib"
	"errors"
	"strconv"
)

type Runner struct {
	variables map[string]int
}

func NewRunner() Runner {
	return Runner{variables: make(map[string]int)}
}

//局部变量
func (runner *Runner) Variables() map[string]int {
	return runner.variables
}

//运算表达式
func (runner *Runner) Evaluate(ast *lib.ASTNode) (ret int, err error) {
	switch ast.Type {
	case lib.ASTNodeType_Programm: //程序入口
		for _, child := range ast.Children {
			if ret, err = runner.Evaluate(child); err != nil {
				return 0, err
			}
		}
	case lib.ASTNodeType_IntDeclaration: //整形变量声明：'int' Id ( = additive) ';'
		runner.variables[ast.Text] = 0 //初始化变量
		if len(ast.Children) > 0 {     //存在对象声明
			if ret, err = runner.Evaluate(ast.Children[0]); err != nil {
				return 0, err
			}
			runner.variables[ast.Text] = ret
		}
	case lib.ASTNodeType_ExpressionStmt: //表达式语句：additive ';'
		if ret, err = runner.Evaluate(ast.Children[0]); err != nil {
			return 0, err
		}
	case lib.ASTNodeType_AssignmentStmt: //赋值语句：assignmentStatement -> id = additive ';'
		var ok bool
		if ret, ok = runner.variables[ast.Text]; !ok {
			return 0, errors.New("变量" + ast.Text + "未定义")
		}
		if ret, err = runner.Evaluate(ast.Children[0]); err != nil {
			return 0, err
		}
		runner.variables[ast.Text] = ret
	case lib.ASTNodeType_Primary: //基础表达式，实际上基本表达式推导后不是基本表达式节点，它直接返回子节点，所以该case不会执行
	case lib.ASTNodeType_Multiplicative: //乘法表达式：
		var l, r int
		if l, err = runner.Evaluate(ast.Children[0]); err != nil {
			return 0, err
		}
		if r, err = runner.Evaluate(ast.Children[1]); err != nil {
			return 0, err
		}
		if ast.Text == "*" { //乘法
			ret = l * r
		} else { //除法
			ret = l / r
		}
	case lib.ASTNodeType_Additive: //加法表达式
		var l, r int
		if l, err = runner.Evaluate(ast.Children[0]); err != nil {
			return 0, err
		}
		if r, err = runner.Evaluate(ast.Children[1]); err != nil {
			return 0, err
		}
		if ast.Text == "+" { //加法
			ret = l + r
		} else { //减法
			ret = l - r
		}
	case lib.ASTNodeType_Identifier: //标识符
		var ok bool
		if ret, ok = runner.variables[ast.Text]; !ok {
			return 0, errors.New("变量" + ast.Text + "未定义")
		}
	case lib.ASTNodeType_IntLiteral: //整型字面量
		if ret, err = strconv.Atoi(ast.Text); err != nil {
			return 0, err
		}
	}
	return
}
