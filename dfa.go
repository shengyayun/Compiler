//自动机状态枚举
package main

type DfaState string

//有限状态机的各种状态
const (
	DfaState_Initial DfaState = "Initial"

	DfaState_If       = "IF"
	DfaState_Id_if1   = "Id_if1"
	DfaState_Id_if2   = "Id_if2"
	DfaState_Else     = "Else"
	DfaState_Id_else1 = "Id_else1"
	DfaState_Id_else2 = "Id_else2"
	DfaState_Id_else3 = "Id_else3"
	DfaState_Id_else4 = "Id_else4"
	DfaState_Int      = "Int"
	DfaState_Id_int1  = "Id_int1"
	DfaState_Id_int2  = "Id_int2"
	DfaState_Id_int3  = "Id_int3"
	DfaState_Id       = "Id"
	DfaState_GT       = "GT"
	DfaState_GE       = "GE"

	DfaState_Assignment = "Assignment"
	DfaState_Plus       = "Plus"
	DfaState_Minus      = "Minus"
	DfaState_Star       = "Star"
	DfaState_Slash      = "Slash"

	DfaState_SemiColon  = "SemiColon"
	DfaState_LeftParen  = "LeftParen"
	DfaState_RightParen = "RightParen"

	DfaState_IntLiteral = "IntLiteral"
)
