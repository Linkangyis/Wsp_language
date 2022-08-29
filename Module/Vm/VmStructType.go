package vm

import(
  "Wsp/Compile"
  "Wsp/Analysis/Lex"
)

var VmFuncRoot = make(map[int]func(TransmitValue)string)

var VmFuncUser = make(map[string]func(map[int]string)string)

var VmClassUser = make(map[string]map[string]func(map[int]string)string)

var VmFuncIs string = "Main"

var OverAllFuncRes FuncResTmp

var TmpCodeRun map[int]string

var FuncPr int

var OpcodeFuncTmp = make(map[string]compile.Res_Struct)

var LexOpFuncTmp = make(map[string]map[int]lex.Lex_Struct)

var LockBreakList string

var DelFunc = make(map[string]int)

var OverLine int

var FuncList compile.Func_Struct

var CodeFilePath string

var ClassId int = -1

var OverClassAll map[string]compile.ClassStruct

var classcdlock int=0

type OpStruct map[int]map[int]compile.Body_Struct_Run

type TransmitValue struct{
    Value string
    Opcode map[int]compile.Body_Struct_Run
    OpRunId int
}

type FuncResTmp struct{
    Name string
    Res string
    IfRes int
}

type VarSoBrkStruct struct{
    Type int
    Text string
}

type CrunTmpStruct struct{
    Type int
    Text string
}