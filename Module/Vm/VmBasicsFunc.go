package vm

import(
  "fmt"
)

/* PRINT VM*/
func Print(From TransmitValue)string{
    Value := From.Value
    fmt.Println(VarAnalysis(Value)[0])
    return "<TRUE>"
}

/* ADD VM */
func Add(From TransmitValue)string{
    Value := From.Value
    list := VarAnalysis(Value)
    Res := 0
    for i:=0;i<=len(list)-1;i++{
        Res += TypeInts(list[i])
    }
    return TypeStrings(Res)
}

func VarSetNum(From TransmitValue)string{
    Lids := From.OpRunId
    Op := From.Opcode[Lids]
    Nums:=make(map[string]int)
    Nums["ABB_VAR"]=0
    Nums["ADD_VAR"]=1
    Type := Nums[Op.Name]
    Ynum:=Read_Array(Op.Text)
    if Type==0{
        AddArray(Op.Text,TypeStrings(TypeInts(Ynum)-1))
    }else{
        AddArray(Op.Text,TypeStrings(TypeInts(Ynum)+1))
    }
    return Ynum
}
/* FOR VM*/
func ForVm(From TransmitValue)string{
    LockBreakList=""
    Value := From.Value
    //IfOneVm("")
    Lids := From.OpRunId
    Op := From.Opcode[Lids]
    Varlist := ForSo(Value)
    if len(Varlist)>1{
        RunCode(Varlist[0])
        for{
            if !IfvmSo(Varlist[1]){
                break
            }
            VmFuncUser[Op.Text](make(map[int]string))
            Res:=LockBreakList
            if Res=="<BREAK>"{
                break;
            }
            RunCode(Varlist[2])
            if Res=="<CONTINUE>"{
                LockBreakList=""
                continue
            }
        }
    }else if len(Varlist)==1 && Varlist[0]!=""{
        for{
            if !IfvmSo(Varlist[0]){
                break
            }
            VmFuncUser[Op.Text](make(map[int]string))
            Res:=LockBreakList
            if Res=="<BREAK>"{
                break;
            }else if Res=="<CONTINUE>"{
                LockBreakList=""
                continue
            }
        }
    }else if Op.Abrk[0].Type==2&&len(Op.Abrk)==2{
        for{
            VmFuncUser[Op.Text](make(map[int]string))
            Res:=LockBreakList
            if Res=="<BREAK>"{
                break;
            }else if Res=="<CONTINUE>"{
                LockBreakList=""
                continue
            }
            if !IfvmSo(Op.Abrk[1].Text){
                break
            }
        }
    }else{
        for{
            VmFuncUser[Op.Text](make(map[int]string))
            Res:=LockBreakList
            if Res=="<BREAK>"{
                break;
            }else if Res=="<CONTINUE>"{
                LockBreakList=""
                continue
            }
        }
    }
    LockBreakList=""
    return "<TRUE>"
}

/* IF VM*/
func IfVm(From TransmitValue)string{
    Op := From.Opcode
    for i:=0;i<=len(Op)-1;i++{
        CodeRun := Op[i].Text
        Ifs:=Op[i].Abrk[0].Text
        Type := Op[i].Type
        if Type==203{
            VmFuncUser[CodeRun](make(map[int]string))
            return ""
        }
        if IfvmSo(Ifs){
            VmFuncUser[CodeRun](make(map[int]string))
            return ""
        }
    }
    return ""
    
}
/* STR VM*/
func StrVm(From TransmitValue)string{
    Res:=From.Opcode[From.OpRunId].Text
    if string(Res[0])=="\""{
        Res=TypeStrings_so(Res)
    }
    return Res
}

/* FUNC VM*/
/*
func FuncVm(From TransmitValue)string{
    /*
    括号内容解析
    Value := From.Value
    fmt.Println(Value)
    
    Lids := From.OpRunId
    Op := From.Opcode[Lids]
    return VmFuncUser[Op.Name]()
}*/
/* FUNC VM 2.0 */
func FuncVm(From TransmitValue)string{
    Lids := From.OpRunId
    Op := From.Opcode[Lids]
    BrkList:=Op.Abrk
    List:=make(map[int]VarSoBrkStruct)
    ListLen:=0
    tmp:=""
    for i:=0;i<=len(BrkList)-1;i++{
        if BrkList[i].Type==1{
            tmp+="["+VarSoAll(BrkList[i].Text)+"]"
            if BrkList[i+1].Type!=1{
                List[ListLen] = VarSoBrkStruct{1,tmp}
                ListLen++
                tmp = ""
            }
        }else{
            List[ListLen]=VarSoBrkStruct{0,BrkList[i].Text}
            ListLen++
        }
    }
    Init:=Op.Name
    Tmps:=FuncName
    defer SetFunc(Tmps)
    for i:=0;i<=len(List)-1;i++{
        if List[i].Type==1{
            Init = Read_Array(Init+List[i].Text)
        }else{
            SetFunc(Tmps)
            Var := VarAnalysis(List[i].Text)
            SetFunc(Init)
            Init = VmFuncUser[Init](Var)
        }
    }
    return Init
}

/* VAR VM*/
func VarVm(From TransmitValue)string{
    Lids := From.OpRunId
    Op := From.Opcode
    VarName := VarNameGenerate(From.Opcode[Lids])
    for i:=Lids;i<=len(Op)-1;i++{
        if Op[i].Type!=301{
            var Values string
            if _,ok:=TmpCodeRun[i];ok{
                Values = TmpCodeRun[i]
            }else{
                Values =CodeBlockRunSingle(Op[i])
                TmpCodeRun[i]=Values
            }
            AddArray(VarName,Values)
            break
        }
    }
    return "<TRUE>"
}

/* VARFUNC VM*/
func VarSo(From TransmitValue)string{
    Lids := From.OpRunId
    Op := From.Opcode[Lids]
    BrkList:=Op.Abrk
    List:=make(map[int]VarSoBrkStruct)
    ListLen:=0
    tmp:=""
    for i:=0;i<=len(BrkList)-1;i++{
        if BrkList[i].Type==1{
            tmp+="["+VarSoAll(BrkList[i].Text)+"]"
            if BrkList[i+1].Type!=1{
                List[ListLen] = VarSoBrkStruct{1,tmp}
                ListLen++
                tmp = ""
            }
        }else{
            List[ListLen]=VarSoBrkStruct{0,BrkList[i].Text}
            ListLen++
        }
    }
    Init:=Read_Array(Op.Text)
    Tmps:=FuncName
    for i:=0;i<=len(List)-1;i++{
        if List[i].Type==1{
            Init = Read_Array(Init+List[i].Text)
        }else{
            SetFunc(Tmps)
            Var := VarAnalysis(List[i].Text)
            SetFunc(Init)
            defer SetFunc(Tmps)
            Init = VmFuncUser[Init](Var)
        }
    }
    return Init
}

func VmSwitch(From TransmitValue)string{
    LockBreakList=""
    Lids := From.OpRunId
    Op := From.Opcode[Lids]
    Condition:=VarSoAll(Op.Name)
    Id:=Op.Text
    Opcode := FuncList.FuncList[Id]
    CodeRunsOpcode := make(OpStruct)
    tmpName := ""
    var Else OpStruct
    ResCodeOp := make(map[string]OpStruct)
    Type := 0
    for i:=0;i<=len(Opcode)-1;i++{
        if Opcode[i][0].Type==214 || Opcode[i][0].Type==215{
            if Opcode[i][0].Type==214{
                Type=1
            }else{
                Type=2
            }
            if tmpName!=""{
                ResCodeOp[tmpName]=CodeRunsOpcode
            }
            tmpName =CodeBlockRunSingle(Opcode[i][1])
            CodeRunsOpcode = make(OpStruct)
            continue
        }
        CodeRunsOpcode[len(CodeRunsOpcode)]= Opcode[i]
    }
    if Type==1{
        ResCodeOp[tmpName]=CodeRunsOpcode
    }else{
        Else=CodeRunsOpcode
    }
    if _,ok:=ResCodeOp[Condition];ok{
        Code:=ResCodeOp[Condition]
        for i:=0;i<=len(Code)-1;i++{
            CodeBlockRun(Code[i])
            if LockBreakList=="<BREAK>"{
                LockBreakList=""
                break;
            }
        }
    }else{
        Code:=Else
        for i:=0;i<=len(Code)-1;i++{
            CodeBlockRun(Code[i])
            if LockBreakList=="<BREAK>"{
                LockBreakList=""
                break;
            }
        }
    }
    return ""
    
}