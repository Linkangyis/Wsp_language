package vm

import (
	memory "Wsp/Module/Memory"
	"sync"
)

func newThread(ThreadName string, Thread *vmStruct) {
	(*(*memory.OpenPointer(threadList).Read()).(*sync.Map)).Store(ThreadName, Thread)
}
