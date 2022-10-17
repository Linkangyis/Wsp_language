package vm

import(
    "strings"
)

func VmReturn(From TransmitValue)string{
    From.VarValue.ResLock=true
    Lids := From.OpRunId
    Op := From.Opcode[Lids+1]
    return CodeBlockRunSingle(Op,From.VarValue)
}

func VmWgo(From TransmitValue)string{
    Lids := From.OpRunId
    Op := From.Opcode[Lids]
    From.VarValue.Govm=false
    From.VarValue.SetWgoId(Op.Abrk[0].Text)
    return ""
}

func VmBreak(From TransmitValue)string{
    From.VarValue.LockBreakList="<BREAK>"
    return ""
}

func VmContinue(From TransmitValue)string{
    From.VarValue.AllCodeStop = true
    From.VarValue.LockBreakList="<CONTINUE>"
    return ""
}

func VmFuncInit(From TransmitValue)string{
    Lids := From.OpRunId
    Op := From.Opcode[Lids]
    Name:=Op.Text
    UserFuncInitManual(Name)
    return Name
}

func VmFuncInitBody(From TransmitValue)string{
    Lids := From.OpRunId
    Op := From.Opcode[Lids]
    Name:=Op.Name
    UserFuncInitManual_9C(Name)
    return Name
}

func VmGlobal(From TransmitValue)string{
    Lids := From.OpRunId
    Op := From.Opcode[Lids]
    Text := Op.Text
    FatherList := strings.Split(Text,",")
    for i:=0;i<=len(FatherList)-1;i++{
        pathMain:=From.VarValue.FILE+"Main/Main"+FatherList[i]
        pathFunc:=From.VarValue.paths+From.VarValue.FuncName+FatherList[i]
        Var_Pointer(From.VarValue.FuncName+FatherList[i],From.VarValue)
        CopyVmArray(pathMain,pathFunc)
    }
    
    return ""
}

func VmClassLock(From TransmitValue)string{
    Lids := From.OpRunId
    Op := From.Opcode[Lids]
    Text := Op.Text
    ClassLock[Text]=true
    return ""
}