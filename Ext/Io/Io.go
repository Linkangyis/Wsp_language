package main

import(
    "Wsp/Module/Vm"
    "io/ioutil"
)

func Func_Info()(map[int]string){
    info := make(map[int]string)
    info[0] = "ReadFile"
    return info
}

func Package_Info()(string){
    info := "Io"
    return info
}

func ReadFile(value map[int]string)(string){
    data, _ := ioutil.ReadFile(vm.FilePathRead(value[0]))
    return string(data)
}

//go build -buildmode=plugin -o io.so Io.go