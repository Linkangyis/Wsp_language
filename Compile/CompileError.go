package compile

import(
    "fmt"
    "os"
)

func ErrorPrint(Line int,ErrorText string){
    fmt.Println("----------------------------------")
    fmt.Println("语法错误：")
    fmt.Println("----------------------------------")
    fmt.Println("    异常行数：",Line)
    fmt.Println("    问题描述：",ErrorText)
    fmt.Println("----------------------------------")
    os.Exit(0)
}