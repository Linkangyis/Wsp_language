package compile

import(
  "Wsp/Analysis/Ast"
)

type Res_Struct struct{
    Func Func_Struct
    Body map[int]map[int]Body_Struct_Run
    Class map[string]ClassStruct
}
type Func_Struct struct{
    FuncList map[string]map[int]map[int]Body_Struct_Run
    FuncVars map[string]map[int]string
}

type ClassStruct struct{
    ClassFunc Func_Struct
    ClassBody map[int]map[int]Body_Struct_Run
}

type Body_Struct_Run struct{
    Type int
    Abrk map[int]ast.Brks
    Name string
    Text string
    Movs string
    Line int
}
var Code ast.Ast_Tree