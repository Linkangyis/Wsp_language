package main

import(
    "Wsp/Module/Vm"
    "strconv"
)

func Func_Info()(map[int]string){
    info := make(map[int]string)
    info[0] = "OpenCont"
    info[1] = "ReadCont"
    return info
}

func Package_Info()(string){
    info := "Raflect"
    return info
}

var ContList = make(map[string]string)
var ContId int = 0

func OpenCont(value map[int]string)(string){
    Id := strconv.Itoa(ContId)
    ContId++
    Name := value[0]
    if _,ok:=vm.EnvList[Name];!ok{
        return "NULL"
    }
    ContList[Id] = Name
    return Id
}


func ReadCont(value map[int]string)(string){
    Tmp := vm.EnvList[ContList[value[0]]]
    return vm.Read_Array(value[1],&Tmp)
}

//go build -buildmode=plugin -o reflect.so Reflect.go