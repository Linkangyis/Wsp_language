package echo

import(
  "fmt"
  "strings"
  "../Types"
  "unicode/utf8"
)
func Arr_Echo(Gra_Comp map[int][4]string){
    for i:=0;i<=len(Gra_Comp)-1;i++{    //TOKEN格式化输出
        Gra_Comp[i]=[4]string{Gra_Comp[i][0],string(strings.Replace(string(Gra_Comp[i][1]),"\n","",-1)),Gra_Comp[i][2],Gra_Comp[i][3]}
        fmt.Println(i,Gra_Comp[i])
    }
}

func Arr_Echo_Build(Gra_Comp map[int][5]string){
    for i:=0;i<=len(Gra_Comp)-1;i++{    //OPCODE格式化输出 开发者
        fmt.Println(i,Gra_Comp[i])
    }
}

func Arr_Echo_Opcode_View(Gra_Comp map[int][6]string,VARS map[string]string,funcname string){
    fmt.Println("\n"+Arr_Echo_Opcode_View_r(125))
    fmt.Printf("变量列表:  ")
    for key,value := range VARS {
        fmt.Printf(key+"="+value+"  ")
    }
    fmt.Printf("\n函数名称:  "+funcname)
    fmt.Println("\n"+Arr_Echo_Opcode_View_r(125))
    fmt.Println("ID"+Arr_Echo_Opcode_View_k(17)+"TOKEN"+Arr_Echo_Opcode_View_k(15)+"参数标识"+Arr_Echo_Opcode_View_k(12)+"形参A"+Arr_Echo_Opcode_View_k(15)+"形参B"+Arr_Echo_Opcode_View_k(15)+"形参C/指向跳转/内存地址")
    fmt.Println(Arr_Echo_Opcode_View_r(125))
    for i:=0;i<=len(Gra_Comp)-1;i++{    //OPCODE格式化输出 用户
    fmt.Printf(types.Strings(i))
    Arr_Echo_Opcode_View_kss(20-len(types.Strings(i)))
        for z:=0;z<=len(Gra_Comp[i])-2;z++{
            fmt.Printf(Gra_Comp[i][z])
            Arr_Echo_Opcode_View_kss(20-utf8.RuneCountInString(Gra_Comp[i][z])-zw(Gra_Comp[i][z]))
        }
        fmt.Println("")
    }
    fmt.Println(Arr_Echo_Opcode_View_r(125))
    fmt.Println(Arr_Echo_Opcode_View_k(125))
}

func Arr_Echo_Opcode_View_kss(Gra_Comp int){
    for i:=0;i<=Gra_Comp-1;i++{
        fmt.Printf(" ")
    }
    fmt.Printf("|")
}
func Arr_Echo_Opcode_View_k(Gra_Comp int)(string){
    re := ""
    for i:=0;i<=Gra_Comp;i++{
        re += " "
    }
    return re
}
func Arr_Echo_Opcode_View_r(Gra_Comp int)(string){
    re := ""
    for i:=0;i<=Gra_Comp;i++{
        re += "-"
    }
    return re
}

func zw(s1 string)(int){
    a := 0
    b := 0
    for _, v := range s1 {
        a++
        if v < 'z' {
    b++
        }
    }
    return a-b
}


func Zero_int(Gra_Comp int)(string){
    re := ""
    for i:=0;i<=Gra_Comp;i++{
        re += "0"
    }
    return re
}

func Debugs(i int){
    fmt.Println("<--->",i)
}