package main

import(
    "Wsp/WVM"
    "Wsp/Token"
    "Wsp/Build"
    "Wsp/Maps"
)

func H_Info()(map[int]string){
    info := make(map[int]string)
    info[0] = "Multithreading"
    info[1] = "Eval"
    return info
}

func Eval(a string)(string){
    code_ok_f:=maps.MAP_COPY_codeok(vm.CodesOkre())
    vm.CodesOk(make(map[int]string))
    str_arr,_:=vm.Parameter_processing(a)
    t:=build.Wsp_Build(token.Wsp_Semantic(token.Wsp_Grammar(token.Wsp_Lexical_func(str_arr[0]))))
    vm.Wsp_VM(t)
    vm.CodesOk(code_ok_f)
    return "1"
}
func Multithreading(a string)(string){
    str_arr,_:=vm.Parameter_processing(a)
    fs:=vm.Ec_Fs()
    go vm.Vm_Code_Run(fs[str_arr[0]])
    return ""
}

//go build -buildmode=plugin -o multithreading.so Multithreading.go