package main

import(
  "fmt"
)
func Func_Info()(map[int]string){
    info := make(map[int]string)
    info[0] = "Testb"
    info[1] = "Tests"
    return info
}
func Package_Info()(string){
    info := "Test"
    return info
}


func Testb(Value map[int]string)(string){
    fmt.Println(Value[0]+Value[1])
    return "TRUE"
}

func Tests(Value map[int]string)(string){
    fmt.Println(Value[0])
    return "TRUE"
}

//go build -buildmode=plugin -o test.so Test.go