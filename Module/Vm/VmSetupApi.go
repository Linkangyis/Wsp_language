package vm

func InitOverAllFuncRes(){
    OverAllFuncRes=FuncResTmp{}
}

func SetVmFuncIs(Name string){
    VmFuncIs=Name
}

func ReadVmFuncIs()string{
    return VmFuncIs
}

func ReadFuncOver()FuncResTmp{
    return OverAllFuncRes
}