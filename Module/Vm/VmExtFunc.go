package vm

import(
  "io/ioutil"
  "os"
  "plugin"
  "strings"
)

func InitFuncUserExt(){
    data, _ := ioutil.ReadFile(os.Getenv("WSPPATH")+"/wsp.ini")
    inis:=strings.Split(string(data),"\n" )
    for i:=0;i<=len(inis)-1;i++{
        iniss:=strings.Split(inis[i],"=" )
        if iniss[0]=="extension"{
            InitFuncUserExtL(iniss[1])
        }
    }
}

func InitFuncUserExtL(file string){
    Tmp, _ := plugin.Open(file)
    AddFunc, _ := Tmp.Lookup("Func_Info")
    Funcmaps:=AddFunc.(func() map[int]string)()
    
    PackageInfo, _ := Tmp.Lookup("Package_Info")
    PackageName:=PackageInfo.(func() string)()
    
    for i:=0;i<=len(Funcmaps)-1;i++{
        Name := Funcmaps[i]
        Names := PackageName+"."+Funcmaps[i]
        if _,ok:=DelFunc[Name];!ok{
            AddFunc, _ = Tmp.Lookup(Name)
            VmFuncUser[Names]=func(Value map[int]string)string{
                paths=TmpPaths
                FuncName=TmpFuncName
                return AddFunc.(func(map[int]string) string)(Value)
            }
        }
    }
}

func Test(){
    InitFuncUserExt()
    Test:=make(map[int]string)
    Test[0]="4"
    VmFuncUser["Tests"](Test)
}
