package main

import(
  "fmt"
)
func H_Info()(map[int]string){
    info := make(map[int]string)
    info[0] = "Test"
    info[1] = "Tests"
    return info
}

func Test(parameter string)(string){
    fmt.Println(parameter)
    return "TRUE"
}

func Tests(parameter string)(string){
    fmt.Println("Hello world")
    return "TRUE"
}

//go build -buildmode=plugin -o test.so Test.go