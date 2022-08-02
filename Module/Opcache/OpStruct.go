package op

import(
  "Wsp/Compile"
)

type Opcodes struct{
    Opcode compile.Res_Struct
    FuncList map[int]string
}