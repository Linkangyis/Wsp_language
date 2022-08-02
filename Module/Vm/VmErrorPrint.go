package vm

import(
  "fmt"
  "os"
)

func Error(err string,line int,errs string){
    fmt.Println("语法错误:\n"+"在第",line,"行时"+err+errs)
    os.Exit(0)
}

func Errors(err string){
    fmt.Println("系统异常:\n  "+err,"\n  在第",OverLine,"行")
    os.Exit(0)
}