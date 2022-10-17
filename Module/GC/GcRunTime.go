package gc

import(
    "time"
    "os"
    "Wsp/Module/RamDisk"
)

func GC_DEL(file string){  //删除目录
    num:=0
    for{
        err:=os.RemoveAll(file)
        if err==nil{
            break
        }
        if num==20{
            break
        }
        num++
    }
}

func GC_Runtime(){
    Maps:=Queue.Get()
    for i:=0;i<=len(Maps)-1;i++{
        GC_DEL(Maps[i])
        delete(Maps,i)
    }
    if !Gc_End{
        time.Sleep(2 * time.Second)
    }
    Size:=readDir("./.<Var_Temps>")
    if Size>GC_Size{
        Gc_Panic = true
        ramdisk.Del("./.<Var_Temps>")
        panic("Wsp GC Error: 内存超出最大GC限制")
    }
    if !Gc_End{
        time.Sleep(2 * time.Second)
        GC_Runtime()
    }
}