package main

import(
    "fmt"
    "math/rand"
    "time"
    "strconv"
    "os"
    "Wsp/Module/GC"
    "Wsp/Module/Vm"
)

func STOI(Value string)int{
    Res, _ := strconv.Atoi(Value)
    return Res
}
func ITOS(value int)string{
    return strconv.Itoa(value)
}

func Exit(value map[int]string)(string){
    gc.Gc_Ends()
    vm.VmEnd()
    os.Exit(0)
    return ""
}
func Print(value map[int]string)(string){
    PrintText := ""
    for i:=0;i<=len(value)-1;i++{
        PrintText += value[i] + " "
    }
    PrintText = PrintText[0:len(PrintText)-1]
    fmt.Printf(PrintText)
    return "<TRUE>"
}

func Stick(value map[int]string)(string){
    PrintText := ""
    for i:=0;i<=len(value)-1;i++{
        PrintText += value[i]
    }
    return PrintText
}

func Println(value map[int]string)(string){
    PrintText := ""
    for i:=0;i<=len(value)-1;i++{
        PrintText += value[i] + " "
    }
    PrintText = PrintText[0:len(PrintText)-1]
    fmt.Println(PrintText)
    return "<TRUE>"
}

func Rand(value map[int]string)(string){
    rand.Seed(time.Now().Unix())
    min := STOI(value[0])
    max := STOI(value[1])
    num := rand.Intn(max-min) + min
    return ITOS(num)
}

func Func_Info()(map[int]string){
    info := make(map[int]string)
    info[0] = "Input"
    info[1] = "Rand"
    info[2] = "Print"
    info[3] = "Println"
    info[4] = "Stick"
    info[5] = "Exit"
    return info
}

func Package_Info()(string){
    info := "Sys"
    return info
}

func Input(value map[int]string)(string){
    var text string
    fmt.Printf(value[0])
    fmt.Scanln(&text)
    return text
}
//go build -buildmode=plugin -o system.so System.go