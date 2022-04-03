package main

import(
  "./WVM"
  "./Token"
  "./Build"
  "os"
  "fmt"
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
func main(){
    vm.DLS_So_Start()
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
    }else{
        fmt.Println("文件或路径不存在")
        fmt.Println(str+"/"+os.Args[1])
        fmt.Println(os.Args[1])
        os.Exit(0)
    }
    Lex:=token.Wsp_Lexical(file)
    Sem:=token.Wsp_Semantic(Lex)
    Gra:=token.Wsp_Grammar(Sem)
    Buildse:=build.Wsp_Build(Gra)
    vm.Wsp_VM(Buildse)
}