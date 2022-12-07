package vm

import(
  "Wsp/Analysis/Lex"
  "Wsp/Analysis/Ast"
  "Wsp/Compile"
  "strings"
)

func VarAnalysis(Code string, VarDump *FileValue)map[int]string{
    Var:=make(map[int]string)
    Tmpcode:=VarCompile(Code).Body[0]
    for i:=0;i<=len(Tmpcode)-1;i++{
        Var[len(Var)]=CodeBlockRunSingle(Tmpcode[i],VarDump)
    }
    return Var
}

func VarCompile(Code string)compile.Res_Struct{
    OpcodeFuncLock.Lock()
    defer OpcodeFuncLock.Unlock()
    if _,ok:=OpcodeFuncTmp[Code];ok{
        return OpcodeFuncTmp[Code]
    }
    Res := compile.Wsp_Compile(ast.Wsp_Ast(Varlex(Code)))
    OpcodeFuncTmp[Code] = Res
    FuncList = Res.Func //尝试修复BUG位置CALL
    return Res
}

func Varlex(Code string)map[int]lex.Lex_Struct{
    LexOpFuncLock.Lock()
    if _,ok:=LexOpFuncTmp[Code];ok{
        defer LexOpFuncLock.Unlock()
        return LexOpFuncTmp[Code]
    }
    LexOpFuncLock.Unlock()
    Res := lex.Wsp_Lexical(Code+" ")
    LexOpFuncLock.Lock()
    LexOpFuncTmp[Code]=Res
    LexOpFuncLock.Unlock()
    return Res
}

func VarSoAll(Code string,VarDump *FileValue)string{
    Tmpcode:=VarCompile(Code).Body[0]
    Var:=CodeBlockRunSingle(Tmpcode[0],VarDump)
    return Var
}

func VarNx(Value string,VarDump *FileValue)string{
    if Value!=""{
        Vt:=VarAnalysis(Value,VarDump)
        Value = ""
        for z:=0;z<=len(Vt)-1;z++{
            Value+=Vt[z]+","
        }
        Value=Value[0:len(Value)-1]
    }
    return Value
}

func RunCode(Code string,Valse *FileValue){
    CodeRun(VarCompile(Code).Body,Valse)
}

func ForSo(Code string)[]string{
    return strings.Split(Code, ";")
}

func IfvmSo(Code string,Valse *FileValue)bool{
    IfLIst:=StickIfCodea(Code)
    for i:=0;i<=len(IfLIst)-1;i++{
        IflistCd:=StickIfCodeb(IfLIst[i])
        if string(IfLIst[i][0])=="("{
            Codes:=IfLIst[i][1:len(IfLIst[i])-1]
            if IfvmSo(Codes,Valse){
                return true
            }
        }
        Temp:=0
        for z:=0;z<=len(IflistCd)-1;z++{
            if string(IflistCd[z][0])=="("{
                Codes:=IflistCd[z][1:len(IflistCd[z])-1]
                if !IfvmSo(Codes,Valse){
                    Temp=1
                }
            }else{
                if !IfOneVm(IflistCd[z],Valse){
                    Temp=1
                }
            }
        }
        if Temp==0{
            return true
        }
    }
    return false
}

func StickIfCodea(Code string)map[int]string{
    if values,ok:=IfStickTmpA[Code];ok{
        return values
    }
    TmpCodeThis:=Code
    Code += "||"
    strlock := 0
    Res := make(map[int]string)
    lens := 0
    str := ""
    for i:=0;i<=len(Code)-1;i++{
        Text1 := string(Code[i])
        Text2:=""
        if i<len(Code)-1{
            Text2 = string(Code[i+1])
        }
        if strlock==0 && Text1==" "{
            continue
        }
        if Text1 == "("{
            strlock++
            //continue
        }else if Text1 == ")"{
            strlock--
            //continue
        }
        if Text1=="|" && Text2=="|" && strlock==0{
            Res[lens]=str
            str=""
            lens++
            i++
            continue
        }
        str += Text1
    }
    IfStickTmpA[TmpCodeThis]=Res
    return Res
}
func StickIfCodeb(Code string)map[int]string{
    if values,ok:=IfStickTmpB[Code];ok{
        return values
    }
    TmpCodeThis:=Code
    Code += "&&"
    strlock := 0
    Res := make(map[int]string)
    lens := 0
    str := ""
    for i:=0;i<=len(Code)-1;i++{
        Text1 := string(Code[i])
        Text2:=""
        if i<len(Code)-1{
            Text2 = string(Code[i+1])
        }
        if strlock==0 && Text1==" "{
            continue
        }
        if Text1 == "("{
            strlock++
            //continue
        }else if Text1 == ")"{
            strlock--
            //continue
        }
        if Text1=="&" && Text2=="&" && strlock==0{
            Res[lens]=str
            str=""
            lens++
            i++
            continue
        }
        str += Text1
    }
    IfStickTmpB[TmpCodeThis]=Res
    return Res
}
func IfOneVm(Code string,VarDump *FileValue)bool{
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
    One :=VarSoAll(Str[0],VarDump)
    Two :=VarSoAll(Str[1],VarDump)
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

func RuncCrunTmps(CodeRuns string,VarDump *FileValue)string{
    TmpResMap:=make(map[int]string)
    IdLen:=0
    locks:=0
    for i:=0;i<=len(CodeRuns)-1;i++{
        Text := string(CodeRuns[i])
        if Text=="("{
            locks++
        }else if Text==")"{
            locks--
        }
        if Text==" "{
            continue
        }
        if (Text=="+" || Text=="-"|| Text=="*"|| Text=="/"|| Text=="%" )&&locks==0{
            if Text+string(CodeRuns[i+1])=="->"{
                
            }else{
                IdLen++
                TmpResMap[IdLen]+=Text
                if string(CodeRuns[i+1])!="+" && string(CodeRuns[i+1])!="-"{
                    IdLen++
                }
                continue
            }
        } 
        TmpResMap[IdLen]+=Text
    }
    TmpsListMap:=make(map[int]CrunTmpStruct)
    for i:=0;i<=len(TmpResMap)-1;i++{
        Text:=TmpResMap[i]
        if Text=="+" || Text=="-"|| Text=="*"|| Text=="/"|| Text=="%" {
            TmpsListMap[len(TmpsListMap)]=CrunTmpStruct{0,Text}
        }else if len(Text)<1{
            
        }else if string(Text[0])=="("{
            TmpsListMap[len(TmpsListMap)]=CrunTmpStruct{0,VarSoAll(Text[1:len(Text)-1],VarDump)}
        }else{
            TmpsListMap[len(TmpsListMap)]=CrunTmpStruct{0,VarSoAll(Text,VarDump)}
        }
    }
    Str:=""
    //fmt.Println(TmpsListMap)
    for i:=0;i<=len(TmpsListMap)-1;i++{
        Str+=TmpsListMap[i].Text
    }
    return Str
}
func ArrayUseSo(n string)string{
    if len(n)<2{
        return n
    }
    if string(n[0:1])=="["{
        return n[1:len(n)-1]
    }
    return n
}