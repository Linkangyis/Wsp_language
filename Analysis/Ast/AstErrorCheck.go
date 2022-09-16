package ast

import(
    "os"
    "fmt"
)

func Error(Text string){
    fmt.Println(Text,"    Line:",Line_Echo())
    os.Exit(0);
}