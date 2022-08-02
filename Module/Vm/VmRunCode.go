package vm

import(
  "Wsp/Compile"
  "Wsp/Module/Ini"
)

func Wsp_Vm(OpcodeStruct compile.Res_Struct){
    Temps:=ini.ReadDelFunc()
    for i:=0;i<=len(Temps)-1;i++{
        DelFunc[Temps[i]]=1
    }
    RootFuncInit()
    UserFuncInit(OpcodeStruct.Func)
    InitFuncUserExt()
    Opcode := OpcodeStruct.Body
    CodeRun(Opcode)
    VmEnd()
}

func CodeRun(Opcode map[int]map[int]compile.Body_Struct_Run)string{
    for i:=0;i<=len(Opcode)-1;i++{
        Value := CodeBlockRun(Opcode[i])
        if Value!="<FNV>"{
            return Value
        }
    }
    return "<FNV>"
}

func CodeBlockRun(OpcodeStructCd map[int]compile.Body_Struct_Run)string{
    var ResLock int = 0
    var ResValue string
    var Govm int = 0
    TmpCodeRun = make(map[int]string)
    for i:=0;i<=len(OpcodeStructCd)-1;i++{
        OverLine = OpcodeStructCd[i].Line
        var lock int = 0
        //Del_Dirl()
        /*func return生效*/
        if OverAllFuncRes.IfRes==1{
            break
        }
        
        if _,ok:=TmpCodeRun[i];ok{
            ResValue = TmpCodeRun[i]
            lock = 1
        }
        
        /*括号值涵盖*/
        Type:=OpcodeStructCd[i].Type
        Value := ""
        if OpcodeStructCd[i].Abrk[0].Type==0{
            Value=OpcodeStructCd[i].Abrk[0].Text
        }
        
        /*return 处理*/
        if Type==207{
            ResLock = 1
            continue
        }
        /*多线程处理*/
        if Type==208{
            Govm = 1
            continue
        }
        /*FORBREAK*/
        if Type==210{
            LockBreakList="<BREAK>"
            return "<BREAK>"
        }
        /*FORCONTINUE*/
        if Type==211{
            LockBreakList="<CONTINUE>"
            return "<CONTINUE>"
        }
        /*运行ROOT函数*/
        if _,ok:=VmFuncRoot[Type];ok && lock!=1{
            if Govm==0{ //多线程关闭状态
                ResValue = VmFuncRoot[Type](TransmitValue{
                    Value : Value,
                    Opcode : OpcodeStructCd,
                    OpRunId : i,
                })
            }else{    //多线程启动状态
                go VmFuncRoot[Type](TransmitValue{
                    Value : Value,
                    Opcode : OpcodeStructCd,
                    OpRunId : i,
                })
                Govm=0   //关闭多线程
            }
        }
        /*输出return内容*/
        if ResLock==1{
            OverAllFuncRes.Res = ResValue
            OverAllFuncRes.IfRes = 1
            return ResValue
        }
    }
    return "<FNV>"
}

func CodeBlockRunSingle(OpcodeStructCd compile.Body_Struct_Run)string{
    /*括号值涵盖*/
    Type:=OpcodeStructCd.Type
    Value := ""
    if OpcodeStructCd.Abrk[0].Type==0{
        Value=OpcodeStructCd.Abrk[0].Text
    }
    /*运行ROOT函数*/
    var Res string
    OpcodeValue:=make(map[int]compile.Body_Struct_Run)
    OpcodeValue[0]=OpcodeStructCd
    if _,ok:=VmFuncRoot[Type];ok{
        Res = VmFuncRoot[Type](TransmitValue{
            Value : Value,
            Opcode : OpcodeValue,
            OpRunId : 0,
        })
    }
    return Res
}