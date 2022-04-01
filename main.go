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
    str, _ := os.Getwd()
    if ok,_ := PathExists(str+"/"+os.Args[1]); !ok {
        fmt.Println("文件或路径不存在",str+"/"+os.Args[1])
        os.Exit(0)
    }
    Lex:=token.Wsp_Lexical(str+"/"+os.Args[1])
    Sem:=token.Wsp_Semantic(Lex)
    Gra:=token.Wsp_Grammar(Sem)
    Buildse:=build.Wsp_Build(Gra)
    vm.Wsp_VM(Buildse)
}