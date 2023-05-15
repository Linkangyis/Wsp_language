package vm

import (
	compile "Wsp/Compile"
	"Wsp/Module/Memory"
	"fmt"
	"sync"
)

var threadList memory.MallocSTRING
var rootFuncList = make(map[int]func(map[int]compile.OpRun, int, *vmStruct) interface{})

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

func (this *vmStruct) tmpLoadOpcode(OpcodeTest map[int]map[int]compile.OpRun) interface{} {
	var Res interface{}
	for runPointer := 0; runPointer <= len(OpcodeTest)-1; runPointer++ {
		fmt.Println(OpcodeTest[runPointer])
		for runPointers := 0; runPointers <= len(OpcodeTest[runPointer])-1; runPointers++ {
			a := OpcodeTest[runPointer]
			thiss := a[runPointers]
			Res = rootFuncList[thiss.Type](a, runPointers, this)
		}
	}
	return Res
}

func WspVm(compiles compile.CompileStruct) {
	newThread := vmStruct{}
	(*(*memory.OpenPointer(threadList).Read()).(*sync.Map)).Store("Main", &newThread)
	newThread.threadName = "Main"

	newThread.opcode = &opcode{
		opcodeBody:      compiles.Opcode,
		opcodeClassList: compiles.ClassList,
		opcodeFuncList:  compiles.FuncList,
	}

	varModulePointer := memory.Malloc()
	varModule := make(AnyArray)
	varModulePointer.Open().SetValue(&varModule)
	newThread.stack = &stack{
		classStack:    make(map[string]memory.MallocSTRING),
		variableStack: varModulePointer,
		funcStack: &funcStack{
			funcStack:   make(map[int]memory.MallocSTRING),
			funcPointer: make(map[string]int),
		},
	}

	newThread.vmStats = &vmStats{
		breakStats: false,
		returnStruct: &returnStruct{
			returnStats: false,
		},
	}
	newThread.tmpRootFuncList = make(map[int]func(map[int]compile.OpRun, int, *vmStruct) interface{})
	for k, v := range rootFuncList {
		newThread.tmpRootFuncList[k] = v
	}

	tv, _ := (*(*memory.OpenPointer(threadList).Read()).(*sync.Map)).Load("Main")
	v := tv.(*vmStruct)
	OpcodeTest := (*v.opcode).opcodeBody
	fmt.Println((*v.opcode).opcodeBody)

	newThread.tmpLoadOpcode(OpcodeTest)
}
