package main

import(
    "../../Types"
    "../../WVM"
    "strings"
)

func H_Info()(map[int]string){
    info := make(map[int]string)
    info[0] = "MUL"
    info[1] = "SUB"
    info[2] = "DIV"
    return info
}

func MUL(a string)(string){
    str_arr := strings.Split(a, ",")
    add_num:=types.Ints(str_arr[0])
    for i:=1;i<=len(str_arr)-1;i++{
        if string(str_arr[i][0]) == string("$"){
            add_num=types.Ints(vm.Vars[strings.Replace(str_arr[i],"$","",-1)])*add_num
        }else{
            add_num=types.Ints(str_arr[i])*add_num
        }
    }
    return types.Strings(add_num)
}

func SUB(a string)(string){
    str_arr:=vm.Parameter_processing(a)
    add_num:=types.Ints(str_arr[0])
    for i:=1;i<=len(str_arr)-1;i++{
        add_num=add_num-types.Ints(str_arr[i])
    }
    return types.Strings(add_num)
}

func DIV(a string)(string){
    str_arr := strings.Split(a, ",")
    add_num:=types.Ints(str_arr[0])
    for i:=1;i<=len(str_arr)-1;i++{
        if string(str_arr[i][0]) == string("$"){
            add_num=add_num/types.Ints(vm.Vars[strings.Replace(str_arr[i],"$","",-1)])
        }else{
            add_num=add_num/types.Ints(str_arr[i])
        }
    }
    return types.Strings(add_num)
}


//go build -buildmode=plugin -o calculation.so Calculation.go