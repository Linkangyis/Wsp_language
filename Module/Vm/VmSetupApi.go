package vm

import(
    "path/filepath"
)

func VarRamStart(){
    VarRam = true
}

func InitOverAllFuncRes(Ls *FileValue){
    Ls.Func=&FuncResTmp{}
    Ls.Func.IfRes=0
}

func SetVmFuncIs(Name string){
    VmFuncIs=Name
}

func ReadVmFuncIs()string{
    return VmFuncIs
}


func WspCodeFile()(string){
    Paths, _ := filepath.Split(CodeFilePath)
    return Paths
}

func WspCodeFileSet(File string){
    CodeFilePath = File
}

func UserFuncRun(FuncName string,Value map[int]string)string{
    Ts:=Mains
    Ts.SetFunc(FuncName)
    return VmFuncUser[FuncName](Value,&Ts)
}

func ReadClassId()string{
    ClassId++
    return "Object-ClassUid<"+TypeStrings(ClassId)+">"
}

func ReadWgoId()string{
    WgoId++
    return "WgoId<"+TypeStrings(WgoId)+">"
}
