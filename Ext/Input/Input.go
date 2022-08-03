package main

import(
    "fmt"
)

func H_Info()(map[int]string){
    info := make(map[int]string)
    info[0] = "Input"
    return info
}

func Input(value map[int]string)(string){
    var text string
	fmt.Printf(value[0])
	fmt.Scanln(&text)
    return text
}

//go build -buildmode=plugin -o input.so Input.go