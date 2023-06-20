package vm

import (
	compile "Wsp/Compile"
	"Wsp/Module/Memory"
	"fmt"
	"sync"
)

func (this *vmStruct) tmpLoadOpcode(OpcodeTest map[int]map[int]compile.OpRun) interface{} {
	var Res interface{}
	for runPointer := 0; runPointer <= len(OpcodeTest)-1; runPointer++ {
		fmt.Println(OpcodeTest[runPointer])
		for runPointers := 0; runPointers <= len(OpcodeTest[runPointer])-1; runPointers++ {
			a := OpcodeTest[runPointer]
			thisOpcode := a[runPointers]
			Res = rootFuncList[thisOpcode.Type](a, runPointers, this)
		}
	}
	return Res
}

func WspVm(compiles compile.CompileStruct) {
	newThreadVar := vmStruct{}
	newThread("Main", &newThreadVar)
	newThreadVar.threadName = "Main"

	newThreadVar.opcode = &opcode{
		opcodeBody:      compiles.Opcode,
		opcodeClassList: compiles.ClassList,
		opcodeFuncList:  compiles.FuncList,
	}

	varModulePointer := memory.Malloc()
	varModule := make(AnyArray)
	varModulePointer.Open().SetValue(&varModule)
	newThreadVar.stack = &stack{
		classStack:    make(map[string]memory.MallocSTRING),
		variableStack: varModulePointer,
		funcStack: &funcStack{
			funcStack:   make(map[int]memory.MallocSTRING),
			funcPointer: make(map[string]int),
		},
	}

	newThreadVar.vmStats = &vmStats{
		breakStats: false,
		returnStruct: &returnStruct{
			returnStats: false,
		},
	}
	newThreadVar.tmpRootFuncList = make(map[int]func(map[int]compile.OpRun, int, *vmStruct) interface{})
	for k, v := range rootFuncList {
		newThreadVar.tmpRootFuncList[k] = v
	}

	tv, _ := (*(*memory.OpenPointer(threadList).Read()).(*sync.Map)).Load("Main")
	v := tv.(*vmStruct)
	OpcodeTest := (*v.opcode).opcodeBody
	fmt.Println((*v.opcode).opcodeBody)

	newThreadVar.tmpLoadOpcode(OpcodeTest)
}
