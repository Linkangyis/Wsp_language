package main

import(
    "../../WVM"
)

func H_Info()(map[int]string){
    info := make(map[int]string)
    info[0] = "Multithreading"
    return info
}

func Multithreading(a string)(string){
    str_arr:=vm.Parameter_processing(a)
    fs:=vm.Ec_Fs()
    go vm.Vm_Code_Run(fs[str_arr[0]])
    return ""
}

//go build -buildmode=plugin -o multithreading.so Multithreading.go