## 简单易懂的编译器

[![LICENSE](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://github.com/shengyayun/Compiler/blob/master/LICENSE)

> 学习资料：[《编译原理之美》](https://time.geekbang.org/column/intro/219) 宫文学 北京物演科技CEO

一个能解析简单的表达式、变量声明和初始化语句、赋值语句的编译器。

#### 功能清单

1. [x] 词法分析：通过有穷自动机将代码转换为token流
2. [x] 语义分析：将token流转换为抽象语法树（AST）
3. [x] 执行程序：运行抽象语法树
4. [x] REPL：Read-Eval-Print Loop

#### 扩展巴科斯范式

```
programm -> intDeclare | expressionStatement | assignmentStatement
intDeclare -> 'int' Id ( = additive) ';'
expressionStatement -> additive ';'
assignmentStatement -> id = additive ';'
additive -> multiplicative ( (+ | -) multiplicative)*
multiplicative -> primary ( (* | /) primary)*
primary -> IntLiteral | Id | '(' additive ')'
```
