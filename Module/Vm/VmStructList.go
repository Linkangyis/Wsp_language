package vm

import (
	"Wsp/Compile"
	"Wsp/Module/Memory"
)

type vmStruct struct {
	threadName  string                //这个线程的名称
	runCodeLine int                   //执行到第几行
	memory      []memory.MallocSTRING //这个线程创建的所有内存

	*stack   //这个线程的堆栈
	*vmStats //这个线程的状态   （系统内存 会变）
	*opcode  //这个线程的中间码 （系统内存 会变）
}
type opcode struct { //中间码
	opcodeBody      map[int]map[int]compile.OpRun  //必须存在
	opcodeFuncList  map[int]compile.FuncStruct     //部分情况下可以为空，但是会导致函数无法正常被初始化（比如多线程的时候这个值为空，但是里面function了一个Test函数）
	opcodeClassList map[string]compile.ClassStruct //上同 无法new一个类
}

type vmStats struct { //虚拟机状态
	breakStats bool //break状态设置
	*returnStruct
}

type returnStruct struct {
	returnStats bool        //判断当前状态是否为return
	returnValue interface{} //return的值
}

type stack struct {
	classStack    map[string]memory.MallocSTRING //用于判断内存地址是否为类  （栈内存 可变内存）
	variableStack memory.MallocSTRING            //变量模块的内存指向       （堆内存 常变内存）

	*funcStack //函数栈
}
type funcStack struct {
	funcStack   map[int]memory.MallocSTRING //第一次启动的时候初始化 指向为引用类型 拷贝无损失，func指向列表   （栈内存  不变内存）
	funcPointer map[string]int              //第一次启动的时候初始化 指向为int类型 拷贝无损失，获取func具体指向（系统内存 少变内存）
}
