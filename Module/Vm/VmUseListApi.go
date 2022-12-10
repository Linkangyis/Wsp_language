package vm

import(
    "io/ioutil"
    "os"
)

func ArrayRead(Value string)map[string]interface{}{
    res := make(map[string]interface{})
    path,ok:=Pointeref[Value]
    if Value[0:2] == "./"{
        path = Value
    }
    if ok || path == Value {
        rd, _ := ioutil.ReadDir(path)
        for _, fi := range rd {
            if fi.IsDir() {
                res[ArrayUseSo(fi.Name())]=ArrayRead(path+"/"+fi.Name()+"/")
            } else {
                res[ArrayUseSo(fi.Name())]=Read_File(path+"/"+fi.Name())
            }
        }
        return res
    }
    return res
}

func VarFree(varName string,VarValue *FileValue){
    File:= So_Array_Stick(VarValue.FuncName+varName,VarValue)
    Del_Dir(File+"/")
    Del_Dir(File)
    Del_File(File)
    Del_Files(File)
}

func NewArrayType(Values map[string]interface{})string{
    RamS:=InitVar("RES",3,FileValue{})
    VarFree("__NEWRES__",&RamS)
    AddArray("__NEWRES__[__NEWRES__]","<INIT>",&RamS)
    ResAId:=Read_Array("__NEWRES__",&RamS)
    Path:=Pointeref[ResAId]+"/"
    CopyArrayStudio(Values,Path)
    Del_Array("__NEWRES__[__NEWRES__]",&RamS)
    return ResAId
}

func FilePathRead(Path string)string{
    file := Path
    str, _ := os.Getwd()
    if Exists(str+"/"+Path){
        file = str+"/"+Path
    }else if Exists(Path){
        file = Path
    }else if Exists(WspCodeFile()+Path){
        file = WspCodeFile()+Path
    }
    return file
}