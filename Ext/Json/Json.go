package main

import(
  "Wsp/Token"
  "encoding/json"
  "io/ioutil"
  "os"
  "Wsp/WVM/Array"
  "Wsp/WVM"
  "Wsp/Types"
)
func H_Info()(map[int]string){
    info := make(map[int]string)
    info[0] = "Json_encode"
    return info
}
func Parameter_processing(a string)(map[int]string){
    map_snum:=0
    returns :=make(map[int]string)
    tokenser:=token.Wsp_Grammar(token.Wsp_Semantic(token.Wsp_Lexical_func(a)))
    for i:=0;i<=len(tokenser)-1;i++{
        if tokenser[i][1]==","{
            map_snum++
        }else if tokenser[i][1]!="$"{
            returns[map_snum]+=tokenser[i][1]
        }
    }
    return returns
}
func arrays(Arrs string)(string){
    lock := 0
    avrs := make(map[int]string)
    avrs_l := 0
    for i:=0;i<=len(Arrs)-1;i++{
        if string(Arrs[i])=="]"{
            lock--
        }else if string(Arrs[i])=="["{
            lock++
        }
        avrs[avrs_l]+=string(Arrs[i])
        if lock==0{
            avrs_l++
        }
    }
    res :=""
    for i:=1;i<=len(avrs)-1;i++{
        res+="["+vm.Var_so_all(types.Strings_so(avrs[i]))+"]"
    }
    return avrs[0]+res
}
func Json_encode(file string)string{
    date:=Parameter_processing(file)
    date[0]=arrays(date[0])
    file=array.So_Array_Stick(date[0])
    bs,_:=json.Marshal(all(file))
    return string(bs)
}
func Read_File(filepath string) string {
   fi, _ := os.Open(filepath)
   fd, _ := ioutil.ReadAll(fi)
   fi.Close()
   return string(fd)
}
func Name_so(n string)string{
    if len(n)<2{
        return n
    }
    if string(n[0:2])=="[\""{
        return n[2:len(n)-2]
    }else if string(n[0:1])=="["{
        return n[1:len(n)-1]
    }
    return n
}
func all(path string)(map[string]interface{}){
    rd, _ := ioutil.ReadDir(path)
    res := make(map[string]interface{})
    for _, fi := range rd {
        if fi.IsDir() {
            res[Name_so(fi.Name())]=all(path+"/"+fi.Name()+"/")
        } else {
            res[Name_so(fi.Name())]=Read_File(path+"/"+fi.Name())
        }
    }
    return res
}

//go build -buildmode=plugin -o json.so Json.go