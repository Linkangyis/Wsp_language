package ast

type BodyAst_Struct struct{
    Type int
    Name string
    Text string
    Sbrk map[int]string
    Mbrk map[int]string
    Xbrk map[int]string
    Abrk map[int]Brks
    Line int
}

type Brks struct{
    Type int
    Text string
}
type FuncAst_Struct struct{
    FuncList map[string]map[int]BodyAst_Struct
    FuncVars map[string]map[int]string
}
type Ast_Tree struct{
    FuncAst FuncAst_Struct
    BodyAst map[int]BodyAst_Struct
    ClassAst map[string]ClassAstStruct
}
type ClassAstStruct struct{
    ClassFunc FuncAst_Struct
    ClassBody map[int]BodyAst_Struct
}
var classlock int = 0
var funcnewlock int = 0