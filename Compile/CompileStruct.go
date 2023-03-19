package compile

type CompileStruct struct{
    Opcode map[int]map[int]OpRun
    FuncList map[int]FuncStruct
    ClassList map[string]ClassStruct
}

type FuncStruct struct{
    Name string
    ValueVar map[int]string
    Opcode map[int]map[int]OpRun
}

type OpRun struct{
    Type int
    Name string
    Text string
    Line int
    ValueList map[int]ValueListStruct
    Register interface{}
}

type ValueListStruct struct{
    Type int
    Value map[int]map[int]OpRun
}

type ClassStruct struct{
    ClassName string
    Opcode map[int]map[int]OpRun
    ClassUserFunc map[string]FuncStruct
}

type Erun struct{
    Type bool
    OpErun map[int]map[int]OpRun
    RunType int
}