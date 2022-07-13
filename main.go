package main

import(
  "Wsp/WVM"
  "Wsp/WVM/Array"
  "Wsp/Token"
  "Wsp/Build"
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
        fmt.Println("Version: V2.0\nOpcache: V1.0\nWsp_To_Go: BETA 1.1.1")
        os.Exit(0)
    }else if os.Args[1] == "WTG"{
        if ok,_ := PathExists(str+"/"+os.Args[2]); ok {
            file = str+"/"+os.Args[2]
        }else if ok,_ := PathExists(os.Args[2]); ok {
            file = os.Args[2]
        }else{
            fmt.Println("文件或路径不存在")
            os.Exit(0)
        }
        Lex:=token.Wsp_Lexical(file)
        Sem:=token.Wsp_Semantic(Lex)
        Gra:=token.Wsp_Grammar(Sem)
        Buildse:=build.Wsp_Build(Gra)
        ioutil.WriteFile(file+"_c.go",[]byte(build.GO_BUILD(Buildse)), 0666)
        os.Exit(0)
    }else if os.Args[1] == "help"{
        if len(os.Args)==2{
            fmt.Println("Wsp 是用来运行Wsp语言代码的解释器\n")
            fmt.Println("使用方法：\n\n        wsp <路径>\n")
            fmt.Println("扩展方法：\n\n        wsp WTG <路径>           WTG是将Wsp编译为Golang代码的一种测试方法，极其不稳定，对于大多数功能无法实现，可以使用基本功能\n")
            fmt.Println("有关该主题的更多信息，请使用“go help OR wsp help ini”。\n")
            os.Exit(0)
        }else if os.Args[2] == "ini"{
            fmt.Println("wsp_func_del   用来禁用函数，用“,”隔开")
            fmt.Println("wsp_debug      用来显示OPCODE数据 默认0关闭 1为开启")
            fmt.Println("wsp_cache      用来开启OPCODE缓存 默认1开启 0关闭")
            fmt.Println("wsp_cache_file 用来设置OPCODE缓存存储路径")
            fmt.Println("extension      用来载入SO动态库扩展，需使用绝对路径")
        }
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
    array.Del_Dir("./Var_Temps")
}