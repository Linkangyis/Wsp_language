package memory

import "sync"

type MemoryStruct struct {
	ReadTime int64
	SetTime  int64
	NewTime  int64
	Value    interface{}
}

var MemoryLoad []sync.Map

var Lock sync.Mutex
