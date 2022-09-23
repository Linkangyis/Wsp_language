package gc

import(
    "strconv"
    "sync"
)

var GC_Size int
var Queue *ConcurMap
var lock bool = true
var Gc_Panic bool
var Gc_End bool = false


func SetGcSize(Size string){
    Sizes, _ := strconv.Atoi(Size[0:len(Size)-1])
    GC_Size = Sizes*1024*1024
    Queue = NewConcurMap()
    Gc_Panic = false
}

type ConcurMap struct {
    Data map[int]string
    Lock *sync.RWMutex
}

func NewConcurMap()  *ConcurMap{
    return &ConcurMap{
        Data: make(map[int]string),
        Lock:&sync.RWMutex{},
    }
}
func (d ConcurMap) Set(v string) {
    d.Lock.Lock()
    defer d.Lock.Unlock()
    d.Data[len(d.Data)]=v
}
func (d ConcurMap) Get() map[int]string{
    d.Lock.RLock()
    defer d.Lock.RUnlock()
    Res:=make(map[int]string)
    for i:=0;i<=len(d.Data)-1;i++{
        Res[i]=d.Data[i]
    }
    return Res
}

func GC_Queue(file string){
    Queue.Set(file)
}

func Gc_Ends(){
    Gc_End = true
}