package main

import "C"
import (
	public "Wsp/Public"
	"fmt"
	"net/rpc"
	"unsafe"
)

var inittype = false

//export INITGODLL
func INITGODLL(Map uintptr, LoadFuncName uintptr, RpcPtr uintptr) uintptr {
	INITS()
	TmpMap := PtrMap(Map)
	FuncName := uintptrToString(LoadFuncName)
	RpcPort := uintptrToString(RpcPtr)
	if !inittype {
		initVm(RpcPort)
		inittype = true
	}
	if _, ok := ConfigMap[FuncName]; ok {
		TmpRes := ConfigMap[FuncName](TmpMap)
		fmt.Scanln(TmpRes)
		return NewRes(&TmpRes)
	}

	NewResT := public.TypeDLL{
		Is_Array: false,
		Text:     "<ERROR>",
	}
	fmt.Scanln(NewRes(&NewResT)) //进入GC不让其消失
	panic("ERROR NOT FUNC" + FuncName)
	return NewRes(&NewResT)
}

//export INITREADALL
func INITREADALL() uintptr {
	INITS()
	TmpMap := make(map[int][]byte)
	for k, _ := range ConfigMap {
		TmpMap[len(TmpMap)] = []byte(k)
	}
	fmt.Scanln(MapPtr(&TmpMap))
	return MapPtr(&TmpMap)
}

func NewRes(Structs *public.TypeDLL) uintptr {

	Structss := public.Type{
		Is_Array: Structs.Is_Array,
		Text:     []byte(Structs.Text),
	}

	return uintptr(unsafe.Pointer(&Structss))
}

func uintptrToString(ptr uintptr) string {
	return *(*string)(unsafe.Pointer(ptr))
}

func StrPtr(s *string) uintptr {
	return uintptr(unsafe.Pointer(s))
}

func PtrMap(ptr uintptr) map[int]string {
	res := make(map[int]string)
	tmp := *(*map[int][]byte)(unsafe.Pointer(ptr))
	for k, v := range tmp {
		res[k] = string(v)
	}
	return res
}

func MapPtr(m *map[int][]byte) uintptr {
	return uintptr(unsafe.Pointer(m))
}

type UserFuncRunVm struct {
	FuncName string
	Value    map[int]string
}

type FilePathReadVm struct {
	Value string
}

type Struct struct{}

type vmStruct struct {
	UserFuncRun  func(string, map[int]string) string
	WspCodeFile  func() string
	FilePathRead func(string) string
}

var vm = vmStruct{}

func initVm(RpcPort string) {
	client, _ := rpc.DialHTTP("tcp", "127.0.0.1:"+RpcPort)
	vm.UserFuncRun = func(Funcname string, Value map[int]string) string {
		args := &UserFuncRunVm{Funcname, Value}
		var reply string
		client.Call("RPCc.UserFuncRun", args, &reply)
		return reply
	}
	vm.WspCodeFile = func() string {
		args := &vmStruct{}
		var reply string
		client.Call("RPCc.WspCodeFile", args, &reply)
		return reply
	}
	vm.FilePathRead = func(Value string) string {
		args := &FilePathReadVm{Value}
		var reply string
		client.Call("RPCc.FilePathRead", args, &reply)
		return reply
	}
}

func main() {

}
