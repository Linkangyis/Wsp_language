package lib

import(
  "strconv"
  "strings"
)

func TypeInts(text string)(int){
    ints, _ := strconv.Atoi(text)
    return ints
}

func TypeStrings(num int)(string){
    Text := strconv.Itoa(num)
    return Text
}

func TypeFloatString(num float64)string{
    return strconv.FormatFloat(num,'f',-1,64)
}
func TypeStrings_so(text string)(string){
    var texts string
    str_arr := strings.Split(text, "")
    for i:=1;i<=len(str_arr)-2;i++{
        texts+=str_arr[i]
    }
    return texts
}

func TypeVar_so(text string)(string){
    var texts string
    str_arr := strings.Split(text, "")
    for i:=1;i<=len(str_arr)-1;i++{
        texts+=str_arr[i]
    }
    return texts
}

func TypeIsNum(s string) bool {
    _, err := strconv.ParseFloat(s, 64)
    return err == nil
}
func TypeTrims(a string)string{
    return strings.Trim(a," ")
}
