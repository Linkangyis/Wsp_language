package main

import(
  "./WVM"
  "./Token"
  "./Build"
  "os"
  "fmt"
  "io/ioutil"
)
func PathExists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil {
        return true, nil
    }
    if os.IsNotExist(err) {
        return false, nil
    }
    return false, err
}
var Files_s string
func main(){
    if len(os.Args)==1{
        fmt.Println("文件或路径不存在")
        os.Exit(0)
    }
    file := ""
    str, _ := os.Getwd()
    if ok,_ := PathExists(str+"/"+os.Args[1]); ok {
        file = str+"/"+os.Args[1]
    }else if ok,_ := PathExists(os.Args[1]); ok {
        file = os.Args[1]
    }else if os.Args[1] == "version"{
        fmt.Println("Version: BETA 1.4\nOpcache V1.0")
        os.Exit(0)
    }else{
        fmt.Println("文件或路径不存在")
        fmt.Println(str+"/"+os.Args[1])
        fmt.Println(os.Args[1])
        os.Exit(0)
    }
    data, _ := ioutil.ReadFile(file)
    files := string(data)
    Inis:=vm.DLS_So_Start()
    cache_file:=Inis["wsp_cache_file"]
    
    if ok,_ := PathExists(cache_file+"/"+build.Md5(files)); ok  && Inis["wsp_cache"]=="1"{
        vm.Wsp_VM(build.Opcaches_Read(cache_file+"/"+build.Md5(files)))
    }else{
        Lex:=token.Wsp_Lexical(file)
        Sem:=token.Wsp_Semantic(Lex)
        Gra:=token.Wsp_Grammar(Sem)
        Buildse:=build.Wsp_Build(Gra)
        if Inis["wsp_cache"]=="1"{
            build.Opcaches_ADD(Buildse,cache_file+"/"+build.Md5(files))
        }
        vm.Wsp_VM(Buildse)
    }
}