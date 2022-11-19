package main

import(
    "fmt"
    "math/rand"
    "time"
    "strconv"
    "os"
    "os/exec"
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
    info[6] = "Free"
    info[7] = "Panic"
    info[8] = "Eval"
    info[9] = "Exec"
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

func Free(value map[int]string)(string){
    varName := value[0]
    Copy := vm.Mains
    File:= vm.So_Array_Stick(Copy.FuncName+varName,&Copy)
    vm.Del_Dir(File+"/")
    vm.Del_Dir(File)
    vm.Del_File(File)
    vm.Del_Files(File)
    return "<TRUE>"
}

func Panic(value map[int]string)(string){
    fmt.Println("WspVm Panic :",value[0])
    Exit(value)
    return ""
}

func Eval(value map[int]string)(string){
    Copy := vm.Mains
    vm.RunCode(value[0],&Copy)
    return ""
}

func Exec(value map[int]string)(string){
    out, _ := exec.Command(value[0]).Output()
    return string(out)
}

//go build -buildmode=plugin -o system.so System.go