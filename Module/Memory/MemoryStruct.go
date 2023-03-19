package memory

type MemoryStruct struct{
    ReadTime int64
    SetTime int64
    NewTime int64
    Value interface{}
}

var MemoryLoad []map[int]*MemoryStruct