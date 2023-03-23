package vm

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type UserFuncRunVm struct {
	FuncName string
	Value    map[int]string
}

type FilePathReadVm struct {
	Value string
}

type Struct struct{}

type RPCc int

func (t *RPCc) UserFuncRun(args *UserFuncRunVm, res *string) error {
	*res = UserFuncRun(args.FuncName, args.Value)
	return nil
}
func (t *RPCc) WspCodeFile(args *Struct, res *string) error {
	*res = WspCodeFile()
	return nil
}
func (t *RPCc) FilePathRead(args *FilePathReadVm, res *string) error {
	*res = FilePathRead(args.Value)
	return nil
}

func StartRPC() {
	arith := new(RPCc)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":25000")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}
