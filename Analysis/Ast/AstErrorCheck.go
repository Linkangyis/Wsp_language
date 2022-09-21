package ast

import(
    "os"
    "fmt"
    "Wsp/Module/Const"
)

func Error(Text string){
    fmt.Println(Text,"    Line:",Line_Echo(),"\n目录：",consts.WspConst.WspFile)
    os.Exit(0);
}