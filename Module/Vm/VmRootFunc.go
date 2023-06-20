package vm

import (
	compile "Wsp/Compile"
	memory "Wsp/Module/Memory"
	"fmt"
	"sync"
)

func init() {
	threadList = memory.Malloc()
	threadList.Open().SetValue(&sync.Map{})
	rootFuncList[200] = func(op map[int]compile.OpRun, self int, this *vmStruct) interface{} {
		a := this.tmpLoadOpcode(op[self].ValueList[0].Value)
		fmt.Println(a)
		return "<NIL>"
	}
	rootFuncList[0] = func(op map[int]compile.OpRun, self int, this *vmStruct) interface{} {
		uString := op[self]
		return uString.Text
	}
}
