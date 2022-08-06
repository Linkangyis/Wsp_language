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
    VmFuncRoot[0]=StrVm
}

func UserFuncInit(funclist compile.Func_Struct){
    TmpList:=center.R_Memory_FromMap()
    for i:=0;i<=len(TmpList)-1;i++{
        Name:=TmpList[i]
        if Name[0:2]=="0x"{
            VmFuncUser[Name]=func(Null map[int]string)string{
                return CodeRun(funclist.FuncList[Name])
            }
        }else{
            VmFuncUser[Name]=func(Var map[int]string)string{
                if _,ok:=DelFunc[Name];ok{
                    Errors("函数"+Name+"被禁用")
                }
                VarFuncIs := funclist.FuncVars[Name]
                for i:=0;i<=len(Var)-1;i++{
                    AddArray(VarFuncIs[i],Var[i])
                }
                OverAllFuncRes.Name = Name
                /*---------------------*/
                Tmp:=ReadVmFuncIs()
                SetVmFuncIs(Name)
                /*---------------------*/
                CodeRun(funclist.FuncList[Name])
                /*---------------------*/
                defer SetVmFuncIs(Tmp)
                defer InitOverAllFuncRes()
                /*---------------------*/
                return OverAllFuncRes.Res
            }
        }
    }
}