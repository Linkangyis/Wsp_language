package vm

import(
    "fmt"
    "os"
    "Wsp/Module/Const"
)

func Error(err string,line int,errs string){
    fmt.Println("语法错误:\n"+"在第",line,"行时"+err+errs,"\n  目录：",consts.WspConst.WspFile)
    os.Exit(0)
}

func Errors(err string){
    fmt.Println("系统异常:\n  "+err,"\n  在第",OverLine,"行","\n  目录：",consts.WspConst.WspFile)
    os.Exit(0)
}

func ErrorFunc(err string){ 
    fmt.Println("-------------------------\n内核异常:\n-------------------------\n  函数 "+err+" 不存在","\n  在第",OverLine,"行\n  目录：",consts.WspConst.WspFile,"\n-------------------------")
    os.Exit(0)
}

func ErrorClass(err string){ 
    fmt.Println("-------------------------\n内核异常:\n-------------------------\n  Class "+err+" 不存在","\n  在第",OverLine,"行\n  目录：",consts.WspConst.WspFile,"\n-------------------------")
    os.Exit(0)
}