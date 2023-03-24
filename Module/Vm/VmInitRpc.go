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

func IfPort(Port int) bool {
	port := TypeStrings(Port)
	conn, err := net.Dial("tcp", ":"+port)
	if err != nil {
		return true
	}
	conn.Close()
	return false
}

func StartRPC(Port int) {
	arith := new(RPCc)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":"+TypeStrings(Port))
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}
