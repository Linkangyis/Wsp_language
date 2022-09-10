package main

import(
    "Wsp/Module/Vm"
    "encoding/json"
)

func Func_Info()(map[int]string){
    info := make(map[int]string)
    info[0] = "Json_encode"
    info[1] = "Json_decode"
    return info
}

func Package_Info()(string){
    info := "Seq"
    return info
}

func Json_encode(value map[int]string)(string){
    Array:=vm.ArrayRead(value[0])
    Res,_:=json.Marshal(Array)
    return string(Res)
}

func Json_decode(value map[int]string)string{
    TmpArray := make(map[string]interface{})
    Json := []byte(value[0])
    json.Unmarshal(Json,&TmpArray)
    return vm.NewArrayType(TmpArray)
}

//go build -buildmode=plugin -o seq.so Seq.go