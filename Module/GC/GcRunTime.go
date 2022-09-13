package gc

import(
    "time"
    "os"
    "Wsp/Module/RamDisk"
)

func GC_DEL(file string){  //删除目录
    os.RemoveAll(file)
}

func GC_Runtime(){
    Size:=readDir("./.<Var_Temps>")
    Maps:=Queue.Get()
    if Size>GC_Size{
        for i:=0;i<=len(Maps)-1;i++{
            GC_DEL(Maps[i])
        }
    }
    Size=readDir("./.<Var_Temps>")
    if Size>GC_Size{
        Gc_Panic = true
        ramdisk.Del("./.<Var_Temps>")
        panic("Wsp GC Error: 内存超出最大GC限制")
    }
    time.Sleep(2 * time.Second)
    GC_Runtime()
}