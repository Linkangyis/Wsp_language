package vm

import(
    "io/ioutil"
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


func NewArrayType(Values map[string]interface{})string{
    RamS:=InitVar("RES",3)
    AddArray("__NEWRES__[__NEWRES__]","<INIT>",&RamS)
    ResAId:=Read_Array("__NEWRES__",&RamS)
    Path:=Pointeref[ResAId]+"/"
    CopyArrayStudio(Values,Path)
    Del_Array("__NEWRES__[__NEWRES__]",&RamS)
    return ResAId
}

