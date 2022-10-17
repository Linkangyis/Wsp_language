package vm

import(
  "Wsp/Compile"
  "Wsp/Module/Ini"
  "Wsp/Module/GC"
)

func Wsp_Vm(OpcodeStruct compile.Res_Struct){
    ClassLock = OpcodeStruct.ClassLock
    Mains = InitVar("Main",0)
    Temps:=ini.ReadDelFunc()
    for i:=0;i<=len(Temps)-1;i++{
        DelFunc[Temps[i]]=1
    }
    RootFuncInit()
    FuncList = OpcodeStruct.Func
    UserFuncInit(OpcodeStruct.Func)
    InitFuncUserExt()
    Opcode := OpcodeStruct.Body
    OverClassAll = OpcodeStruct.Class
    CodeRun(Opcode,&Mains)
}

func CodeRun(Opcode map[int]map[int]compile.Body_Struct_Run,Vales*FileValue)string{
    for i:=0;i<=len(Opcode)-1;i++{
        if gc.Gc_Panic{
            for{}
        }
        if Vales.AllCodeStop{
            continue;
        }
        Value := CodeBlockRun(Opcode[i],Vales)
        if Value!="<FNV>"{
            return Value
        }
    }
    return "<FNV>"
}

func CodeBlockRun(OpcodeStructCd map[int]compile.Body_Struct_Run,Vales *FileValue)string{
    var ResValue string
    TmpCodeRun = make(map[int]string)
    for i:=0;i<=len(OpcodeStructCd)-1;i++{
        if gc.Gc_Panic{
            for{}
        }
        OverLine = OpcodeStructCd[i].Line
        var lock int = 0
        //Del_Dirl()
        
        /*func return生效*/
        if Vales.Func.IfRes==1{
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
        
        /* FOR END IF*/
        if Vales.LockBreakList!=""{
            break
        }
        
        /*运行ROOT函数*/
        if _,ok:=VmFuncRoot[Type];ok && lock!=1{
            if Vales.Govm{ //多线程关闭状态
                ResValue = VmFuncRoot[Type](TransmitValue{
                    Value : Value,
                    Opcode : OpcodeStructCd,
                    OpRunId : i,
                    VarValue : Vales,
                })
            }else{    //多线程启动状态
                Id := "Null"
                if Vales.WgoIdName!=""{
                    Id = Vales.WgoIdName
                }else{
                    Id = ReadWgoId()
                }
                Tmp:=InitVar(Id,1)
                CopyVmArray(Mains.FILE,Tmp.FILE)
                go func(i int){
                    VmFuncRoot[Type](TransmitValue{
                        Value : Value,
                        Opcode : OpcodeStructCd,
                        OpRunId : i,
                        VarValue : &Tmp,
                    })
                    gc.GC_Queue(Tmp.FILE)
                }(i)
                Vales.Govm=true   //关闭多线程
            }
        }
        /*输出return内容*/
        if Vales.ResLock{
            Vales.Func.Res = ResValue
            Vales.Func.IfRes = 1
            Vales.ResLock=false
            return ResValue
        }
    }
    return "<FNV>"
}

func CodeBlockRunSingle(OpcodeStructCd compile.Body_Struct_Run,VarTmpP *FileValue)string{
    /*括号值涵盖*/
    if gc.Gc_Panic{
        for{}
    }
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
            VarValue : VarTmpP,
        })
    }
    return Res
}