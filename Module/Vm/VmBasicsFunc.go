package vm

import(
  "fmt"
  "Wsp/Module/Formula"
)

/* PRINT VM*/
func Print(From TransmitValue)string{
    Value := From.Value
    fmt.Println(VarAnalysis(Value,From.VarValue)[0])
    return "<TRUE>"
}

/* ADD VM */
func Add(From TransmitValue)string{
    Value := From.Value
    list := VarAnalysis(Value,From.VarValue)
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
    Ynum:=Read_Array(Op.Text,From.VarValue)
    if Type==0{
        AddArray(Op.Text,TypeStrings(TypeInts(Ynum)-1),From.VarValue)
    }else{
        AddArray(Op.Text,TypeStrings(TypeInts(Ynum)+1),From.VarValue)
    }
    return Ynum
}
/* FOR VM*/
func ForVm(From TransmitValue)string{
    From.VarValue.LockBreakList=""
    Value := From.Value
    //IfOneVm("")
    Lids := From.OpRunId
    Op := From.Opcode[Lids]
    Varlist := ForSo(Value)
    if len(Varlist)>1{
        RunCode(Varlist[0],From.VarValue)
        for{
            if !IfvmSo(Varlist[1],From.VarValue){
                break
            }
            VmFuncUser[Op.Text](make(map[int]string),From.VarValue)
            RunCode(Varlist[2],From.VarValue)
            Res:=From.VarValue.LockBreakList
            if Res=="<BREAK>"{
                break;
            }
            if Res=="<CONTINUE>"{
                From.VarValue.LockBreakList=""
                From.VarValue.AllCodeStop = false
            }
        }
    }else if len(Varlist)==1 && Varlist[0]!=""{
        for{
            if !IfvmSo(Varlist[0],From.VarValue){
                break
            }
            VmFuncUser[Op.Text](make(map[int]string),From.VarValue)
            Res:=From.VarValue.LockBreakList
            if Res=="<BREAK>"{
                break;
            }else if Res=="<CONTINUE>"{
                From.VarValue.LockBreakList=""
                From.VarValue.AllCodeStop = false
            }
        }
    }else if Op.Abrk[0].Type==2&&len(Op.Abrk)==2{
        for{
            VmFuncUser[Op.Text](make(map[int]string),From.VarValue)
            Res:=From.VarValue.LockBreakList
            if Res=="<BREAK>"{
                break;
            }else if Res=="<CONTINUE>"{
                From.VarValue.LockBreakList=""
                From.VarValue.AllCodeStop = false
            }
            if !IfvmSo(Op.Abrk[1].Text,From.VarValue){
                break
            }
        }
    }else{
        for{
            VmFuncUser[Op.Text](make(map[int]string),From.VarValue)
            Res:=From.VarValue.LockBreakList
            if Res=="<BREAK>"{
                break;
            }else if Res=="<CONTINUE>"{
                From.VarValue.LockBreakList=""
                From.VarValue.AllCodeStop = false
            }
        }
    }
    From.VarValue.LockBreakList=""
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
            VmFuncUser[CodeRun](make(map[int]string),From.VarValue)
            return ""
        }
        if IfvmSo(Ifs,From.VarValue){
            VmFuncUser[CodeRun](make(map[int]string),From.VarValue)
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

/* STICK VM*/
func VmStick(From TransmitValue)string{
    Value := From.Value
    list:=VarAnalysis(Value,From.VarValue)
    Res := ""
    for i:=0;i<=len(list)-1;i++{
        Res+=list[i]
    }
    return Res
}

/* 孤儿函数*/
func EvalVm(From TransmitValue)string{
    return VarSoAll(From.Value,From.VarValue)
}

/* 计算虚拟机*/
func CrunVm(From TransmitValue)string{
    Lids := From.OpRunId
    Op := From.Opcode[Lids]
    CodeRuns:="1*"+Op.Text
    Str := RuncCrunTmps(CodeRuns,From.VarValue)
    postfixExp:=crun.PostfixCRun(Str)
    value := crun.RunNums(postfixExp)
    ResValue:=TypeFloatString(value)
    return ResValue
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
            tmp+="["+VarSoAll(BrkList[i].Text,From.VarValue)+"]"
            if BrkList[i+1].Type!=1{
                List[ListLen] = VarSoBrkStruct{1,tmp}
                ListLen++
                tmp = ""
            }
        }else if BrkList[i].Type==0{
            List[ListLen]=VarSoBrkStruct{0,BrkList[i].Text}
            ListLen++
        }else if BrkList[i].Type==3{
            List[ListLen]=VarSoBrkStruct{3,BrkList[i].Text}
            ListLen++
        }
    }
    Init:=Op.Name
    Tmps:=From.VarValue.FuncName
    defer From.VarValue.SetFunc(Tmps)
    for i:=0;i<=len(List)-1;i++{
        if List[i].Type==1{
            if Init[0]!='$'&&string(Init[0:2])!="0x"{
                //Init = Read_Array(Init+List[i].Text)
                Init = string(Init[TypeInts(List[i].Text[1:len(List[i].Text)-1])]);
            }else{
                Init = Read_Array(Init+List[i].Text,From.VarValue)
            }
        }else if List[i].Type==3{
            Name:=List[i].Text+Init
            i++
            if List[i].Type==0&&i<len(List){
                From.VarValue.SetFunc(Tmps)
                Var := VarAnalysis(List[i].Text,From.VarValue)
                Temps:=From.VarValue.AllOverPaths
                From.VarValue.AllOverPaths=From.VarValue.FILE
                From.VarValue.RootCd("Class"+Init)
                From.VarValue.SetFunc(Name)
                defer From.VarValue.SetFunc(Tmps)
                RunCode("$this=\""+Init+"\";",From.VarValue)
                Init = VmClassUser[Init][Name](Var,From.VarValue)
                From.VarValue.AllOverPaths=Temps
                //i++
            }else{
                i--
                Id:=Init
                Temps:=From.VarValue.AllOverPaths
                From.VarValue.AllOverPaths=From.VarValue.FILE
                From.VarValue.RootCd("Class"+Id)
                Tmps:=From.VarValue.FuncName
                defer From.VarValue.SetFunc(Tmps)
                From.VarValue.SetFunc("")
                Init = Read_Array(List[i].Text,From.VarValue)
                From.VarValue.AllOverPaths=Temps
            }
        }else{
            From.VarValue.SetFunc(Tmps)
            Var := VarAnalysis(List[i].Text,From.VarValue)
            From.VarValue.SetFunc(Init)
            if _,ok:=VmFuncUser[Init];!ok{
                ErrorFunc(Init)
            }
            Init = VmFuncUser[Init](Var,From.VarValue)
        }
    }
    
    return Init
}

/* VAR VM*/
func VarVm(From TransmitValue)string{
    Lids := From.OpRunId
    Op := From.Opcode
    BrkList:=From.Opcode[Lids].Abrk
    
    class := false
    if BrkList[0].Type==3{
        class=true
    }
    VarName := VarNameGenerate(From.Opcode[Lids],From.VarValue)
    for i:=Lids;i<=len(Op)-1;i++{
        if Op[i].Type!=301{
            var Values string
            if _,ok:=TmpCodeRun[i];ok{
                TmpCodeRunLock.Lock()
                Values = TmpCodeRun[i]
                TmpCodeRunLock.Unlock()
            }else{
                Values =CodeBlockRunSingle(Op[i],From.VarValue)
                TmpCodeRunLock.Lock()
                TmpCodeRun[i]=Values
                TmpCodeRunLock.Unlock()
            }
            
            if class{
                VarCdName := Read_Array(VarName,From.VarValue)
                Id := VarCdName
                Temps:=From.VarValue.AllOverPaths
                From.VarValue.AllOverPaths=From.VarValue.FILE
                From.VarValue.RootCd("Class"+Id)
                Tmps:=From.VarValue.FuncName
                defer From.VarValue.SetFunc(Tmps)
                From.VarValue.SetFunc("")
                AddArray(BrkList[0].Text,Values,From.VarValue)
                From.VarValue.AllOverPaths=Temps
            }else{
                AddArray(VarName,Values,From.VarValue)
            }
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
            tmp+="["+VarSoAll(BrkList[i].Text,From.VarValue)+"]"
            if BrkList[i+1].Type!=1{
                List[ListLen] = VarSoBrkStruct{1,tmp}
                ListLen++
                tmp = ""
            }
        }else if BrkList[i].Type==0{
            List[ListLen]=VarSoBrkStruct{0,BrkList[i].Text}
            ListLen++
        }else if BrkList[i].Type==3{
            List[ListLen]=VarSoBrkStruct{3,BrkList[i].Text}
            ListLen++
        }
    }
    Init:=Read_Array(Op.Text,From.VarValue)
    Tmps:=From.VarValue.FuncName
    for i:=0;i<=len(List)-1;i++{
        if List[i].Type==1{
            if Init[0]!='$'&&string(Init[0:2])!="0x"{
                Init = string(Init[TypeInts(List[i].Text[1:len(List[i].Text)-1])]);
            }else{
                Init = Read_Array(Init+List[i].Text,From.VarValue)
            }
        }else if List[i].Type==3{
            Name:=List[i].Text+Init
            i++
            if List[i].Type==0&&i<len(List){
                From.VarValue.SetFunc(Tmps)
                Var := VarAnalysis(List[i].Text,From.VarValue)
                Temps:=From.VarValue.AllOverPaths
                From.VarValue.AllOverPaths=From.VarValue.FILE
                From.VarValue.RootCd("Class"+Init)
                From.VarValue.SetFunc(Name)
                defer From.VarValue.SetFunc(Tmps)
                RunCode("$this=\""+Init+"\";",From.VarValue)
                Init = VmClassUser[Init][Name](Var,From.VarValue)
                From.VarValue.AllOverPaths=Temps
                //i++
            }else{
                i--
                Id:=Init
                Temps:=From.VarValue.AllOverPaths
                From.VarValue.AllOverPaths=From.VarValue.FILE
                From.VarValue.RootCd("Class"+Id)
                Tmps:=From.VarValue.FuncName
                defer From.VarValue.SetFunc(Tmps)
                From.VarValue.SetFunc("")
                Init = Read_Array(List[i].Text,From.VarValue)
                From.VarValue.AllOverPaths=Temps
            }
        }else{
            From.VarValue. SetFunc(Tmps)
            Var := VarAnalysis(List[i].Text,From.VarValue)
             From.VarValue.SetFunc(Init)
            defer  From.VarValue.SetFunc(Tmps)
            if _,ok:=VmFuncUser[Init];!ok{
                ErrorFunc(Init)
            }
            Init = VmFuncUser[Init](Var,From.VarValue)
        }
    }
    return Init
}

/*VM SWITCH*/
func VmSwitch(From TransmitValue)string{
    From.VarValue.LockBreakList=""
    Lids := From.OpRunId
    Op := From.Opcode[Lids]
    Condition:=VarSoAll(Op.Name,From.VarValue)
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
            tmpName =CodeBlockRunSingle(Opcode[i][1],From.VarValue)
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
            CodeBlockRun(Code[i],From.VarValue)
            if From.VarValue.LockBreakList=="<BREAK>"{
                From.VarValue.LockBreakList=""
                break;
            }
        }
    }else{
        Code:=Else
        for i:=0;i<=len(Code)-1;i++{
            CodeBlockRun(Code[i],From.VarValue)
            if From.VarValue.LockBreakList=="<BREAK>"{
                From.VarValue.LockBreakList=""
                break;
            }
        }
    }
    return ""
    
}

/*VM CLASS*/
func VmClass(From TransmitValue)string{
    Lids := From.OpRunId
    Op := From.Opcode[Lids]
    Name := Op.Name
    if _,ok:=ClassLock[Name];!ok{
        ErrorClass(Name)
    }
    if !ClassLock[Name]{
        ErrorClass(Name)
    }
    IdRes:=ReadClassId()
    UserClassInit(OverClassAll[Name],IdRes,From.VarValue)
    BrkList:=Op.Abrk
    if len(BrkList)>0&&(BrkList[0].Type==0){
        Values:=VarAnalysis(BrkList[0].Text,From.VarValue)
        Tmps:=From.VarValue.FuncName
        defer From.VarValue.SetFunc(Tmps)
        Temps:=From.VarValue.AllOverPaths
        From.VarValue.AllOverPaths=From.VarValue.FILE
        From.VarValue.RootCd("Class"+IdRes)
        From.VarValue.SetFunc("_init_"+IdRes)
        RunCode("$this=\""+IdRes+"\";",From.VarValue)
        if _,ok:=VmClassUser[IdRes]["_init_"+IdRes];ok{
            VmClassUser[IdRes]["_init_"+IdRes](Values,From.VarValue)
        }
        From.VarValue.AllOverPaths=Temps
    }
    return IdRes
}

/*VM LEN*/
func VmLen(From TransmitValue)string{
    Value := From.Value
    Text := VarAnalysis(Value,From.VarValue)[0]
    if len(Text)>2{
        if Text[0:2]=="0x"{
            return TypeStrings(len(ArrayRead(Text)))
        }
    }
    return TypeStrings(len(Text))
}