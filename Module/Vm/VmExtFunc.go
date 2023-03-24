package vm

import (
	"Wsp/LoadDll"
	"io/ioutil"
	"strings"
)

func InitFuncUserExt() {
	data, _ := ioutil.ReadFile("./wsp.ini")
	inis := strings.Split(string(data), "\n")
	for i := 0; i <= len(inis)-1; i++ {
		iniss := strings.Split(inis[i], "=")
		if iniss[0] == "extension" {
			InitFuncUserExtL(iniss[1])
		}
	}
}
func InitFuncUserExtL(file string) {
	Tmp := load.OpenDLL(file)
	AddFunc := Tmp.ReadConfig()
	TmpName := Tmp.LoadFunc("LOADINIT", RpcPort)(make(map[int]string)).Text
	for i := 0; i <= len(AddFunc)-1; i++ {
		Name := TmpName + "." + AddFunc[i]
		TNmae := AddFunc[i]
		VmFuncUser[Name] = func(Value map[int]string, Varls *FileValue) string {
			return Tmp.LoadFunc(TNmae, RpcPort)(Value).Text
		}
	}
}

/*
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
            VmFuncUser[Names]=func(Value map[int]string,Varls *FileValue)string{
                AddFunc, _ = Tmp.Lookup(Name)
                Varls.paths=Varls.TmpPaths
                Varls.FuncName=Varls.TmpFuncName
                return AddFunc.(func(map[int]string) string)(Value)
            }
        }
    }
}*/
