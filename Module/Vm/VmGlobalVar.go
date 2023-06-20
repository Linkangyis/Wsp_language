package vm

import (
	compile "Wsp/Compile"
	memory "Wsp/Module/Memory"
)

var threadList memory.MallocSTRING
var rootFuncList = make(map[int]func(map[int]compile.OpRun, int, *vmStruct) interface{})
