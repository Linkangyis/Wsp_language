package main

import(
    "../../WVM"
    "encoding/base64"
)

func H_Info()(map[int]string){
    info := make(map[int]string)
    info[0] = "Base64_Encode"
    info[1] = "Base64_Decode"
    return info
}

func Base64_Encode(a string)(string){
    str_arr:=vm.Parameter_processing(a)
    data := str_arr[0]
    uEnc := base64.URLEncoding.EncodeToString([]byte(data))
    return string(uEnc)
}
func Base64_Decode(a string)(string){
    str_arr:=vm.Parameter_processing(a)
    data := str_arr[0]
    sDec, err := base64.StdEncoding.DecodeString(data)
    if err != nil {
        return "NULL"
    }
    return string(sDec)
}
//go build -buildmode=plugin -o base64.so Base64.go