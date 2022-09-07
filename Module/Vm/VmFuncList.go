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
        }else{
            VmFuncUser[Name]=func(Var map[int]string,Vales *FileValue)string{
                if _,ok:=DelFunc[Name];ok{
                    Errors("函数"+Name+"被禁用")
                }
                VarFuncIs := funclist.FuncVars[Name]
                for i:=0;i<=len(Var)-1;i++{
                    AddArray(VarFuncIs[i],Var[i],Vales)
                }
                OverAllFuncRes.Name = Name
                /*---------------------*/
                Tmp:=ReadVmFuncIs()
                SetVmFuncIs(Name)
                /*---------------------*/
                CodeRun(funclist.FuncList[Name],Vales)
                /*---------------------*/
                defer SetVmFuncIs(Tmp)
                defer InitOverAllFuncRes()
                /*---------------------*/
                return OverAllFuncRes.Res
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
            OverAllFuncRes.Name = Name
            /*---------------------*/
            Tmp:=ReadVmFuncIs()
            SetVmFuncIs(Name)
            /*---------------------*/
            CodeRun(funclist.FuncList[Fname],Vales)
            /*---------------------*/
            defer SetVmFuncIs(Tmp)
            defer InitOverAllFuncRes()
            /*---------------------*/
            return OverAllFuncRes.Res
        }
    }
    VmClassUser[Id]=ListFunc
}