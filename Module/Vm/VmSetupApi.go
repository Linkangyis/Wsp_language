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

func WspCodeFile()(string){
    return CodeFilePath
}

func WspCodeFileSet(File string){
    CodeFilePath = File
}

func UserFuncRun(FuncName string,Value map[int]string)string{
    return VmFuncUser[FuncName](Value)
}