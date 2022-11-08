package vm

import (
    "net/http"
    "github.com/gorilla/websocket"
    "io/ioutil"
    "crypto/md5"
    "encoding/hex"
    "os"
    "fmt"
)

func SververSocketClient(Port string,Files string,ContName string){
    if _,ok:=WebSocketLock[ContName];ok{
        return
    }else{
        WebSocketLock[ContName] = true
    }
    var upgrader = websocket.Upgrader{
        ReadBufferSize:   1024,
        WriteBufferSize:  1024,
        HandshakeTimeout: 5 * 1000,
    }
    http.HandleFunc("/Server_"+ContName, func(w http.ResponseWriter, r *http.Request) {
        conn, _ := upgrader.Upgrade(w, r, nil) 
        Socket:=ServerSocket{
            WebSokcet : conn,
        }
        for{
            Maps := ReadAllDir(Files)
            for _,Struct := range Maps {
                Socket.Read()
                
                
                Socket.Send(Struct.Name)
                Socket.Send(Struct.Md5)
                
                Type:= Socket.Read()
                if string(Type)=="No"{
                    Times:= Socket.Read()
                    if GetFileModTime(Struct.Name)>int64(TypeInts(string(Times))){
                        Socket.Send("<BTB>")
                        Socket.Send(Struct.File)
                    }else{
                        Socket.Send("<ATB>")
                        Text := Socket.Read()
                        New_File(Struct.Name)
                        New_File_Var(Struct.Name,string(Text))
                    }
                    
                }
            }
        }
    })
    http.ListenAndServe(":"+Port, nil)
}


func SververSocketClientUser(Port string,Files string,ContName string){
    if _,ok:=WebSocketLock[ContName+"_User"];ok{
        return
    }else{
        WebSocketLock[ContName+"_User"] = true
    }
    var upgrader = websocket.Upgrader{
        ReadBufferSize:   1024,
        WriteBufferSize:  1024,
        HandshakeTimeout: 5 * 1000,
    }
    http.HandleFunc("/Server_"+ContName+"_User", func(w http.ResponseWriter, r *http.Request) {
        conn, _ := upgrader.Upgrade(w, r, nil) 
        Socket:=ServerSocket{
            WebSokcet : conn,
        }
        for{
            Socket.Read()
            File:=Socket.Read()
            if(File=="<YC>"){
                break
            }
            md5:=Socket.Read()
            
            if Exists(File){
                data, _ := ioutil.ReadFile(File)
                FileData := string(data)
                if md5!=Md5(FileData){
                    Socket.Send("No")
                    Socket.Send(TypeStrings(int(GetFileModTime(File))))
                    IfType := Socket.Read()
                    if IfType=="<BTB>"{
                        FileText := Socket.Read()
                        New_File(File)
                        New_File_Var(File,FileText)
                    }else{
                        Socket.Send(FileData)
                    }
                }else{
                    Socket.Send("Yes")
                }
            }else{
                Socket.Send("No")
                Socket.Send("0")
                IfType := Socket.Read()
                if IfType=="<BTB>"{
                    FileText := Socket.Read()
                    New_File(File)
                    New_File_Var(File,FileText)
                }
            }
        }
    })
    http.ListenAndServe(":"+Port, nil)
}

func SyncVarSever(Ip string,Port string,ContName string,Files string){
    Socket := ListenSocket{}
    types:=Socket.Client("ws://"+Ip+":"+Port+"/Server_"+ContName+"_User")
    if types!=nil{
        fmt.Println("连接到堆时出现异常")
        os.Exit(0);
    }
    for{
        Maps := ReadAllDir(Files)
        for _,Struct := range Maps {
            Socket.Send("Run")
            Socket.Send(Struct.Name)
            Socket.Send(Struct.Md5)
            
            
            Type:= Socket.Read()
            if string(Type)=="No"{
                Times:= Socket.Read()
                if GetFileModTime(Struct.Name)>int64(TypeInts(string(Times))){
                    Socket.Send("<BTB>")
                    Socket.Send(Struct.File)
                }else{
                    Socket.Send("<ATB>")
                    Text := Socket.Read()
                    New_File(Struct.Name)
                    New_File_Var(Struct.Name,string(Text))
                }
            }
        }
    }
}

func SyncVar(Ip string,Port string,ContName string,Files string){
    Socket := ListenSocket{}
    types:=Socket.Client("ws://"+Ip+":"+Port+"/Server_"+ContName)
    if types!=nil{
        fmt.Println("连接到堆时出现异常")
        os.Exit(0);
    }
    for{
        Socket.Send("Run")
        File:=Socket.Read()
        if(File=="<YC>"){
            break
        }
        md5:=Socket.Read()
        
        
        if Exists(File){
            data, _ := ioutil.ReadFile(File)
            FileData := string(data)
            if md5!=Md5(FileData){
                Socket.Send("No")
                Socket.Send(TypeStrings(int(GetFileModTime(File))))
                IfType := Socket.Read()
                if IfType=="<BTB>"{
                    FileText := Socket.Read()
                    New_File(File)
                    New_File_Var(File,FileText)
                }else{
                    Socket.Send(FileData)
                }
            }else{
                Socket.Send("Yes")
            }
        }else{
            Socket.Send("No")
            Socket.Send("0")
            IfType := Socket.Read()
            if IfType=="<BTB>"{
                FileText := Socket.Read()
                New_File(File)
                New_File_Var(File,FileText)
            }
        }
    }
}


func Md5(str string) string {
    h := md5.New()
    h.Write([]byte(str))
    return hex.EncodeToString(h.Sum(nil))
}


func ReadAllDir(Dir string)map[int]FileWebSocket{
    ResMap := make(map[int]FileWebSocket)
    files, _ := ioutil.ReadDir(Dir)
    for _, f := range files{
        if IsDir(Dir+f.Name()+"/"){
            TmpMap := ReadAllDir(Dir+f.Name()+"/")
            for _,Tstruct:= range TmpMap{
                ResMap[len(ResMap)]=Tstruct
            }
        }else{
            data, _ := ioutil.ReadFile(Dir+f.Name())
            FileData := string(data)
            ResMap[len(ResMap)]=FileWebSocket{
                Name : Dir+f.Name(),
                Md5  : Md5(FileData),
                File : FileData,
            }
        }
    }
    return ResMap
}



func (This *ListenSocket)Send(Text string){
    This.WebSokcet.WriteMessage(1, []byte(Text))
}
func (This *ListenSocket)Read()string{
    _, msg, err := This.WebSokcet.ReadMessage()
    if err != nil {
        return "<YC>"
    }
    return string(msg)
}

func (This *ListenSocket)Client(Host string)error{
    conn, _, err := websocket.DefaultDialer.Dial(Host, nil)
    if err != nil {
        return err
    }
    This.WebSokcet=conn
    return nil
}


func (This *ServerSocket)Send(Text string){
    This.WebSokcet.WriteMessage(This.MsgType, []byte(Text));
}


func (This *ServerSocket)Read()string{
    msgType, Res ,_ := This.WebSokcet.ReadMessage()
    This.MsgType = msgType
    return string(Res)
}







