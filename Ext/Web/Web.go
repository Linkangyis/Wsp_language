package main

import(
    "../../WVM"
    "net/http"
    "fmt"
)

var web_list = make(map[string]string)
func H_Info()(map[int]string){
    info := make(map[int]string)
    info[0] = "New_Web"
    info[1] = "Web_Print"
    info[2] = "Web_Start"
    return info
}
var Tmps_w http.ResponseWriter
var Tmps_r *http.Request

func Web_Print(a string)(string){
    var_arr:=vm.Parameter_processing(a)
    fmt.Fprintf(Tmps_w, var_arr[0])
    return ""
}
func sayhelloName(w http.ResponseWriter, r *http.Request) {
    fs:=vm.Ec_Fs()
    Tmps_r=r
    Tmps_w=w
    vm.Vm_Code_Run(fs[web_list[Tmps_r.URL.Path[1:]]])
}


func New_Web(a string)(string){
    
    var_arr:=vm.Parameter_processing(a)
    file := var_arr[0]
    web_list[file]=var_arr[1]
    http.HandleFunc("/"+file, sayhelloName)
    return "1"
}

func Web_Start(a string)(string){
    var_arr:=vm.Parameter_processing(a)
    port := var_arr[0]
    http.ListenAndServe(":"+port, nil)
    return "1"
}

//go build -buildmode=plugin -o web.so Web.go