package vm

func VmReturn(From TransmitValue)string{
    From.VarValue.ResLock=true
    return ""
}

func VmWgo(From TransmitValue)string{
    From.VarValue.Govm=false
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