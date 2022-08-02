package ini

import(
  "io/ioutil"
  "os"
  "strings"
)

func ReadIni()(map[string]string){
    re := make(map[string]string)
    debugs = 0
    data, _ := ioutil.ReadFile(os.Getenv("WSPPATH")+"/wsp.ini")
    inis:=strings.Split(string(data),"\n" )
    for i:=0;i<=len(inis)-1;i++{
        iniss:=strings.Split(inis[i],"=" )
        if iniss[0]=="wsp_debug" && iniss[1]=="1"{
            debugs = 1
        }else if iniss[0]=="wsp_func_del"{
            wsp_func_del = strings.Split(iniss[1], ",")
        }else if iniss[0] != "" && len(iniss)>1{
            re[iniss[0]]=iniss[1]
        }
    }
    return re
}
func ReadDelFunc()[]string{
    return wsp_func_del
}
func DebugsIf()bool{
    switch debugs{
        case 0:
            return false
        case 1:
            return true
    }
    return false
}