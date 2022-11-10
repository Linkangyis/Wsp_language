package main

import(
    "fmt"
    "Wsp/Module/Vm"
)
func Func_Info()(map[int]string){
    info := make(map[int]string)
    info[0] = "ReturnArray"
    info[1] = "ReadArray"
    return info
}
func Package_Info()(string){
    info := "Test"
    return info
}


func ReturnArray(Value map[int]string)(string){
    ResMap := make(map[string]interface{})   //多维数组 1层
    ResMap["1"] = make(map[string]interface{})    //多维数组二层
    ResMap["0"] = "TestArray[0]"
    ResMap["1"].(map[string]interface{})["2"] = "TestArray[1][2]"
    return vm.NewArrayType(ResMap)    //写入RAM并获取引用指针
}

func ReadArray(Value map[int]string)(string){
    Array:=vm.ArrayRead(Value[0])   //输入引用指针，将指针转化为 map[string]interface{} 多维数组类型
    fmt.Println(Array)
    return "TRUE"
}

//go build -buildmode=plugin -o test.so Test.go
