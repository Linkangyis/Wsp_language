package main

import(
    "strings"
    "Wsp/Module/Vm"
    "net/http"
    "fmt"
    "path/filepath"
)

var list = make(map[string]string)
func Func_Info()(map[int]string){
    info := make(map[int]string)
    info[0] = "New_Web"
    info[1] = "Print"
    info[2] = "Start"
    info[3] = "Start_Ssl"
    info[4] = "Header_Set"
    info[5] = "New_WebFiles"
    info[6] = "SslGoto"
    return info
}

func Package_Info()(string){
    info := "Web"
    return info
}

var Tmps_w http.ResponseWriter
var Tmps_r *http.Request


func Print(var_arr map[int]string)(string){
    fmt.Fprintf(Tmps_w, var_arr[0])
    return ""
}
func Header_Set(var_arr map[int]string)(string){
    Tmps_w.Header().Set(var_arr[0], var_arr[1])
    return ""
}
func webTmp(w http.ResponseWriter, r *http.Request) {
    Tmps_r=r
    Tmps_w=w
    if value,ok:=list[Tmps_r.URL.Path[1:]];ok{
        vm.UserFuncRun(value,make(map[int]string))
    }else{
        vm.UserFuncRun(list["/"],make(map[int]string))
    }
}

func New_WebFiles(var_arr map[int]string)(string){
    dir, _ := filepath.Split(vm.WspCodeFile())
    http.Handle("/"+var_arr[0]+"/",http.StripPrefix("/"+var_arr[0], http.FileServer(http.Dir(dir+var_arr[1]))))
    return ""
}
 
func New_Web(var_arr map[int]string)(string){
    if var_arr[0]=="" || var_arr[0]=="/"{
        list["/"]=var_arr[1]
        http.HandleFunc("/", webTmp)
    }else{
        file := var_arr[0]
        list[file]=var_arr[1]
        http.HandleFunc("/"+file, webTmp)
    }
    return "1"
}

func Start(var_arr map[int]string)(string){
    port := var_arr[0]
    http.ListenAndServe(":"+port, nil)
    return "1"
}
func Start_Ssl(var_arr map[int]string)(string){
    port := var_arr[0]
    certpem :=var_arr[1]
    keypem :=var_arr[2]
    http.ListenAndServeTLS(":"+port,vm.FilePathRead(certpem),vm.FilePathRead(keypem), nil)
    return "1"
}

func SslGoto(var_arr map[int]string)(string){
    go http.ListenAndServe(":"+var_arr[0], http.HandlerFunc(SslGotoCode))
    return "1"
}

func SslGotoCode(w http.ResponseWriter, req *http.Request) { 
    _host := strings.Split(req.Host, ":")
    target := "https://" + strings.Join(_host, ":") + req.URL.Path
    if len(req.URL.RawQuery) > 0 {
        target += "?" + req.URL.RawQuery
    }
    http.Redirect(w, req, target,http.StatusTemporaryRedirect)
}

//go build -buildmode=plugin -o web.so Web.go