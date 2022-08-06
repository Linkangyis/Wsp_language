package main

import(
    "time"
    "strconv"
)

func Func_Info()(map[int]string){
    info := make(map[int]string)
    info[0] = "Time"
    info[1] = "Sleep"
    return info
}

func Package_Info()(string){
    info := "Time"
    return info
}


func Ints(text string)(int){
    ints, _ := strconv.Atoi(text)
    return ints
}

func Time(a map[int]string)(string){
    return strconv.FormatInt(time.Now().UnixNano(),10)
}

func Sleep(value map[int]string)(string){
    time.Sleep(time.Duration(Ints(value[0]))*time.Second)
    return "1"
}

//go build -buildmode=plugin -o time.so Time.go