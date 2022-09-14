package vm

import(
  "Wsp/Compile"
  "Wsp/Module/Memory"
)

func RootFuncInit(){
    VmFuncRoot[200]=Print
    VmFuncRoot[201]=ForVm
    VmFuncRoot[205]=FuncVm
    VmFuncRoot[301]=VarVm
    VmFuncRoot[302]=VarSo
    VmFuncRoot[303]=VarSetNum
    VmFuncRoot[304]=CrunVm
    VmFuncRoot[209]=Add
    VmFuncRoot[202]=IfVm
    VmFuncRoot[213]=VmSwitch
    VmFuncRoot[216]=EvalVm
    VmFuncRoot[217]=VmStick
    VmFuncRoot[218]=VmClass
    VmFuncRoot[219]=VmLen
    VmFuncRoot[207]=VmReturn
    VmFuncRoot[208]=VmWgo
    VmFuncRoot[210]=VmBreak
    VmFuncRoot[211]=VmContinue
    VmFuncRoot[220]=VmFuncInit
    VmFuncRoot[221]=VmFuncInitBody
    VmFuncRoot[222]=VmGlobal
    VmFuncRoot[0]=StrVm
}

func UserFuncInit(funclist compile.Func_Struct){
    TmpList:=center.R_Memory_FromMap()
    for i:=0;i<=len(TmpList)-1;i++{
        Name:=TmpList[i]
        if Name[0:2]=="0x"{
            VmFuncUser[Name]=func(Null map[int]string,Vales *FileValue)string{
                return CodeRun(funclist.FuncList[Name],Vales)
            }
        }else if Name[0:2]!="0F"&&Name[0:2]!="9C"{
            VmFuncUser[Name]=func(Var map[int]string,Vales *FileValue)string{
                if _,ok:=DelFunc[Name];ok{
                    Errors("函数"+Name+"被禁用")
                }
                VarFuncIs := funclist.FuncVars[Name]
                for i:=0;i<=len(Var)-1;i++{
                    AddArray(VarFuncIs[i],Var[i],Vales)
                }
                Vales.Func.Name = Name
                /*---------------------*/
                Tmp:=ReadVmFuncIs()
                SetVmFuncIs(Name)
                /*---------------------*/
                Del_Dir(Vales.paths)
                New_File(Vales.paths)
                CodeRun(funclist.FuncList[Name],Vales)
                /*---------------------*/
                defer SetVmFuncIs(Tmp)
                defer InitOverAllFuncRes(Vales)
                /*---------------------*/
                return Vales.Func.Res
            }
        }
    }
}

func UserClassInit(Class compile.ClassStruct,Id string,Vales *FileValue){
    Temps:=Vales.AllOverPaths
    Vales.AllOverPaths=Vales.FILE
    Vales.RootCd("Class"+Id)
    Tmps:=Vales.FuncName
    defer Vales.SetFunc(Tmps)
    Vales.SetFunc("")
    CodeRun(Class.ClassBody,Vales)
    Vales.AllOverPaths=Temps
    
    funclist:=Class.ClassFunc
    ListFunc:=make(map[string]func(map[int]string,*FileValue)string)
    for name, _ := range funclist.FuncVars{
        Name:=name+Id
        Fname:=name
        ListFunc[Name]=func(Var map[int]string,Vales *FileValue)string{
            VarFuncIs := funclist.FuncVars[Fname]
            for i:=0;i<=len(Var)-1;i++{
                AddArray(VarFuncIs[i],Var[i],Vales)
            }
            Vales.Func.Name = Name
            /*---------------------*/
            Tmp:=ReadVmFuncIs()
            SetVmFuncIs(Name)
            /*---------------------*/
            CodeRun(funclist.FuncList[Fname],Vales)
            /*---------------------*/
            defer SetVmFuncIs(Tmp)
            defer InitOverAllFuncRes(Vales)
            /*---------------------*/
            return Vales.Func.Res
        }
    }
    VmClassUser[Id]=ListFunc
}


func UserFuncInitManual(Name string){
    funclist := FuncList
    VmFuncUser[Name]=func(Var map[int]string,Vales *FileValue)string{
        if _,ok:=DelFunc[Name];ok{
            Errors("函数"+Name+"被禁用")
        }
        VarFuncIs := funclist.FuncVars[Name]
        for i:=0;i<=len(Var)-1;i++{
            AddArray(VarFuncIs[i],Var[i],Vales)
        }
        Vales.Func.Name = Name
        /*---------------------*/
        Tmp:=ReadVmFuncIs()
        SetVmFuncIs(Name)
        /*---------------------*/
        Del_Dir(Vales.paths)
        New_File(Vales.paths)
        CodeRun(funclist.FuncList[Name],Vales)
        /*---------------------*/
        defer SetVmFuncIs(Tmp)
        defer InitOverAllFuncRes(Vales)
        /*---------------------*/
        return Vales.Func.Res
    }
}

func UserFuncInitManual_9C(Name string){
    funclist := FuncList
    Names:="9C"+Name
    VmFuncUser[Name]=func(Var map[int]string,Vales *FileValue)string{
        if _,ok:=DelFunc[Name];ok{
            Errors("函数"+Name+"被禁用")
        }
        VarFuncIs := funclist.FuncVars[Names]
        for i:=0;i<=len(Var)-1;i++{
            AddArray(VarFuncIs[i],Var[i],Vales)
        }
        Vales.Func.Name = Name
        /*---------------------*/
        Tmp:=ReadVmFuncIs()
        SetVmFuncIs(Name)
        /*---------------------*/
        Del_Dir(Vales.paths)
        New_File(Vales.paths)
        CodeRun(funclist.FuncList[Names],Vales)
        /*---------------------*/
        defer SetVmFuncIs(Tmp)
        defer InitOverAllFuncRes(Vales)
        /*---------------------*/
        return Vales.Func.Res
    }
}