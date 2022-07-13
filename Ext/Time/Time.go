package main

import(
    "time"
    "strconv"
    "Wsp/Types"
    "Wsp/WVM"
)

func H_Info()(map[int]string){
    info := make(map[int]string)
    info[0] = "Time"
    info[1] = "Sleep"
    return info
}

func nulls(a string)(string){
    return a
}
func Time(a string)(string){
    nulls(a)
    return strconv.FormatInt(time.Now().UnixNano(),10)
}

func Sleep(a string)(string){
    str_arr,_:=vm.Parameter_processing(a)
    time.Sleep(time.Duration(types.Ints(str_arr[0]))*time.Second)
    return "1"
}

//go build -buildmode=plugin -o time.so Time.go