//Token相关
package lib

type Token struct {
	Type TokenType
	Text string
}

//Token流
type Tokens []Token

type TokenType string

const (
	TokenType_Plus  TokenType = "Plus"  // +
	TokenType_Minus           = "Minus" // -
	TokenType_Star            = "Star"  // *
	TokenType_Slash           = "Slash" // /

	TokenType_GE = "GE" // >=
	TokenType_GT = "GT" // >
	TokenType_EQ = "EQ" // ==
	TokenType_LE = "LE" // <=
	TokenType_LT = "LT" // <

	TokenType_SemiColon  = "SemiColon"  // ;
	TokenType_LeftParen  = "LeftParen"  // (
	TokenType_RightParen = "RightParen" // )

	TokenType_Assignment = "Assignment" // =

	TokenType_If   = "If"
	TokenType_Else = "Else"

	TokenType_Int = "Int"

	TokenType_Identifier = "Identifier" //标识符

	TokenType_IntLiteral    = "IntLiteral"    //整型字面量
	TokenType_StringLiteral = "StringLiteral" //字符串字面量
)
