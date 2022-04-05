package main

import(
    "../../WVM"
    "fmt"
)

func H_Info()(map[int]string){
    info := make(map[int]string)
    info[0] = "Input"
    return info
}

func Input(a string)(string){
    str_arr:=vm.Parameter_processing(a)
    var text string
	fmt.Printf(str_arr[0])
	fmt.Scanln(&text)
    return text
}

//go build -buildmode=plugin -o scanlns.so Scanlns.go