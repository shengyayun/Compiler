package main

type DfaState uint8

//有限状态机的各种状态
const (
	DfaState_Initial DfaState = iota

	DfaState_If
	DfaState_Id_if1
	DfaState_Id_if2
	DfaState_Else
	DfaState_Id_else1
	DfaState_Id_else2
	DfaState_Id_else3
	DfaState_Id_else4
	DfaState_Int
	DfaState_Id_int1
	DfaState_Id_int2
	DfaState_Id_int3
	DfaState_Id
	DfaState_GT
	DfaState_GE

	DfaState_Assignment
	DfaState_Plus
	DfaState_Minus
	DfaState_Star
	DfaState_Slash

	DfaState_SemiColon
	DfaState_LeftParen
	DfaState_RightParen

	DfaState_IntLiteral
)
