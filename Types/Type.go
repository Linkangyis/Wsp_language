package types

import(
  "strconv"
  "strings"
)

func Ints(text string)(int){
    ints, _ := strconv.Atoi(text)
    return ints
}

func Strings(num int)(string){
    Text := strconv.Itoa(num)
    return Text
}

func Strings_so(text string)(string){
    var texts string
    str_arr := strings.Split(text, "")
    for i:=1;i<=len(str_arr)-2;i++{
        texts+=str_arr[i]
    }
    return texts
}

func Var_so(text string)(string){
    var texts string
    str_arr := strings.Split(text, "")
    for i:=1;i<=len(str_arr)-1;i++{
        texts+=str_arr[i]
    }
    return texts
}

func IsNum(s string) bool {
    _, err := strconv.ParseFloat(s, 64)
    return err == nil
}
func Trims(a string)string{
    return strings.Trim(a," ")
}

type Builds_Parameter struct {
    a  string
    b  string
    c  string
    opcode map[int][6]string
    lens int
    fs map[string]map[int][6]string
    ft map[string]string
}