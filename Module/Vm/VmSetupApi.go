package vm

import(
    "path/filepath"
)

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
    Paths, _ := filepath.Split(CodeFilePath)
    return Paths
}

func WspCodeFileSet(File string){
    CodeFilePath = File
}

func UserFuncRun(FuncName string,Value map[int]string)string{
    return VmFuncUser[FuncName](Value,&Mains)
}

func ReadClassId()string{
    ClassId++
    return "Object-ClassUid<"+TypeStrings(ClassId)+">"
}

func ReadWgoId()string{
    WgoId++
    return "WgoId<"+TypeStrings(WgoId)+">"
}