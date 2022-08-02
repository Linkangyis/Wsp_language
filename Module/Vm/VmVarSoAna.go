package vm

import(
  "Wsp/Analysis/Lex"
  "Wsp/Analysis/Ast"
  "Wsp/Compile"
  "strings"
)

func VarAnalysis(Code string)map[int]string{
    Var:=make(map[int]string)
    Tmpcode:=VarCompile(Code+" ").Body[0]
    for i:=0;i<=len(Tmpcode)-1;i++{
        Var[len(Var)]=CodeBlockRunSingle(Tmpcode[i])
    }
    return Var
}

func VarCompile(Code string)compile.Res_Struct{
    if _,ok:=OpcodeFuncTmp[Code];ok{
        return OpcodeFuncTmp[Code]
    }
    Res := compile.Wsp_Compile(ast.Wsp_Ast(Varlex(Code)))
    OpcodeFuncTmp[Code] = Res
    return Res
}

func Varlex(Code string)map[int]lex.Lex_Struct{
    if _,ok:=LexOpFuncTmp[Code];ok{
        return LexOpFuncTmp[Code]
    }
    Res := lex.Wsp_Lexical(Code+" ")
    LexOpFuncTmp[Code]=Res
    return Res
}

func VarSoAll(Code string)string{
    Tmpcode:=VarCompile(Code).Body[0]
    Var:=CodeBlockRunSingle(Tmpcode[0])
    return Var
}

func RunCode(Code string){
    CodeRun(VarCompile(Code).Body)
}

func ForSo(Code string)[]string{
    return strings.Split(Code, ";")
}

func IfvmSo(Code string)bool{
    IfLIst:=strings.Split(Code, "||")
    for i:=0;i<=len(IfLIst)-1;i++{
        IflistCd:=strings.Split(Code, "&&")
        Temp:=0
        for z:=0;z<=len(IflistCd)-1;z++{
            if !IfOneVm(IflistCd[i]){
                Temp=1
            }
        }
        if Temp==0{
            return true
        }
    }
    return false
}

func IfOneVm(Code string)bool{
    Tmp:=Varlex(Code)
    Str := []string{}
    DlLock := 0
    Type := -1
    for i:=0;i<=len(Tmp)-1;i++{
        if Tmp[i].Type==97&&Tmp[i+1].Type==95{
            Type=1
            Str=strings.Split(Code, "<=")
        }else if Tmp[i].Type==97{
            Type=5
            Str=strings.Split(Code, "<")
        }else if Tmp[i].Type==98&&Tmp[i+1].Type==95{
            Type=2
            Str=strings.Split(Code, ">=")
        }else if Tmp[i].Type==98{
            Type=4
            Str=strings.Split(Code, ">")
        }else if Tmp[i].Type==99{
            Type=3
            Str=strings.Split(Code, "!=")
        }else if Tmp[i].Type==95{
            DlLock++ 
            if DlLock==2{
                Type=0
                Str=strings.Split(Code, "==")
            }
        }
    }
    One :=VarSoAll(Str[0])
    Two :=VarSoAll(Str[1])
    switch Type{
        case 0:
            if One==Two{
                return true
            }
        case 1:
            if TypeInts(One)<=TypeInts(Two){
                return true
            }
        case 2:
            if TypeInts(One)>=TypeInts(Two){
                return true
            }
        case 3:
            if One!=Two{
                return true
            }
        case 4:
            if TypeInts(One)>TypeInts(Two){
                return true
            }
        case 5:
            if TypeInts(One)<TypeInts(Two){
                return true
            }
    }
    return false
}