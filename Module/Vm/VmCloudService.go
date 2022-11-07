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

func SververSocketClient(Port string,Files string){
    var upgrader = websocket.Upgrader{
        ReadBufferSize:   1024,
        WriteBufferSize:  1024,
        HandshakeTimeout: 5 * 1000,
    }
    http.HandleFunc("/Auth", func(w http.ResponseWriter, r *http.Request) {
        conn, _ := upgrader.Upgrade(w, r, nil) 
        for{
            Maps := ReadAllDir(Files)
            for _,Struct := range Maps {
                msgType, _, _ := conn.ReadMessage()
                
                conn.WriteMessage(msgType, []byte(Struct.Name));
                conn.WriteMessage(msgType, []byte(Struct.Md5));
                
                msgType, Type,_ := conn.ReadMessage()
                if string(Type)=="No"{
                    msgType, Times ,_ := conn.ReadMessage()
                    if GetFileModTime(Struct.Name)>int64(TypeInts(string(Times))){
                        conn.WriteMessage(msgType, []byte("<BTB>"));
                        conn.WriteMessage(msgType, []byte(Struct.File));
                    }else{
                        conn.WriteMessage(msgType, []byte("<ATB>"));
                        _, Text ,_ := conn.ReadMessage()
                        New_File(Struct.Name)
                        New_File_Var(Struct.Name,string(Text))
                    }
                    
                }
            }
        }
    })
    http.ListenAndServe(":"+Port, nil)
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

func SyncVar(Ip string,Port string){
    Socket := ListenSocket{}
    types:=Socket.Client("ws://"+Ip+":"+Port+"/Auth")
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













