package vm

import(
  "Wsp/Compile"
  "Wsp/Analysis/Lex"
  "github.com/gorilla/websocket"
  "sync"
)

var VmFuncRoot = make(map[int]func(TransmitValue)string)

var VmFuncUser = make(map[string]func(map[int]string,*FileValue)string)

var VmClassUser = make(map[string]map[string]func(map[int]string,*FileValue)string)

var VmFuncIs string = "Main"

var TmpCodeRunLock sync.Mutex

var TmpCodeRun map[int]string

var FuncPr int

var OpcodeFuncTmp = make(map[string]compile.Res_Struct)

var LexOpFuncTmp = make(map[string]map[int]lex.Lex_Struct)

var LexOpFuncLock sync.Mutex
var OpcodeFuncLock sync.Mutex

var IfStickTmpA  = make(map[string]map[int]string)
var IfStickTmpB  = make(map[string]map[int]string)

var DelFunc = make(map[string]int)

var OverLine int

var FuncList compile.Func_Struct

var CodeFilePath string

var ClassId int = -1

var WgoId int = -1

var Mains FileValue

var VarRam bool = false

var WgoList = make(map[string]*FileValue)

var OverClassAll map[string]compile.ClassStruct

var ClassLock = make(map[string]bool)

var EnvList = make(map[string]FileValue)

var EnvListLock sync.Mutex

type OpStruct map[int]map[int]compile.Body_Struct_Run

type TransmitValue struct{
    Value string
    Opcode map[int]compile.Body_Struct_Run
    OpRunId int
    VarValue *FileValue
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

type FileValue struct{
    FILE string
    AllOverPaths string
    paths string
    TmpPaths string
    FuncName string
    TmpFuncName  string
    LockBreakList string
    AllCodeStop bool
    ResLock bool
    Govm bool
    WgoIdName string
    Func *FuncResTmp
}

type FileWebSocket struct{
    Name string
    Md5 string
    File string
}

type ListenSocket struct{
    WebSokcet *websocket.Conn
}
