package vm

import(
    "strings"
    "os"
    "io/ioutil"
    "io"
    "path"
    "regexp"
    "fmt"
    "Wsp/Compile"
    "Wsp/Module/RamDisk"
    "time"
)

func PathFileStick(file string,str string)string{
    Ls:=strings.Split(file,"/")
    var Res string
    lock:=false
    if file[0]=='.'{
        lock=true
    }
    for i:=0;i<=len(Ls)-1;i++{
        if Ls[i]==""{
            Res += "/"
        }else if Ls[i]=="ForMain"&&i==2{
            Res += "/For"+str
        }else if len(Ls[i])>3&&(Ls[i][0:3]=="For"&&i==2){
            Res += "/For"+str
        }else{
            Res += "/"+Ls[i]
        }
    }
    if lock{
        return Res[1:]
    }
    return Res+"/"
}

func InitVar(Id string,ifs int,Father FileValue)FileValue{
    if ifs==0{
        return FileValue{
            FILE : "./.<Var_Temps>/For"+Id+"/",
            AllOverPaths  : "./.<Var_Temps>/For"+Id+"/",
            paths  : "./.<Var_Temps>/For"+Id+"/Main/",
            TmpPaths  : "./.<Var_Temps>/For"+Id+"/Main/",
            FuncName  : "Main",
            TmpFuncName  : "Main",
            AllCodeStop : false,
            ResLock : false,
            Govm  : true,
            Func : &FuncResTmp{},
        }
    }else if ifs==1{
        return FileValue{
            FILE : "./.<Var_Temps>/For"+Id+"/",
            AllOverPaths  : PathFileStick(Mains.AllOverPaths,Id),
            paths  : PathFileStick(Mains.paths,Id),
            TmpPaths  : PathFileStick(Mains.TmpPaths,Id),
            FuncName  : Mains.FuncName,
            TmpFuncName  : Mains.TmpFuncName,
            AllCodeStop : false,
            ResLock : false,
            Govm  : true,
            Func : &FuncResTmp{},
        }
    }else if ifs==4{
        return FileValue{
            FILE : "./.<Var_Temps>/For"+Id+"/",
            AllOverPaths  : "./.<Var_Temps>/For"+Id+"/",
            paths  : "./.<Var_Temps>/For"+Id+"/"+Id+"-Main/",
            TmpPaths  : "./.<Var_Temps>/For"+Id+"/"+Id+"-Main/",
            FuncName  : Id+"-Main",
            TmpFuncName  : Id+"-Main",
            AllCodeStop : false,
            ResLock : false,
            Govm  : true,
            Func : &FuncResTmp{},
        }
    }else if ifs==5{
        return FileValue{
            FILE : "./.<Var_Temps>/For"+Id+"/",
            AllOverPaths  : PathFileStick(Father.AllOverPaths,Id),
            paths  : PathFileStick(Father.paths,Id),
            TmpPaths  : PathFileStick(Father.TmpPaths,Id),
            FuncName  : Father.FuncName,
            TmpFuncName  : Father.TmpFuncName,
            AllCodeStop : false,
            ResLock : false,
            Govm  : true,
            Func : &FuncResTmp{},
        }
    }
    return FileValue{
        FILE : "./.<Var_Temps>/For"+Id+"/",
        AllOverPaths  : "./.<Var_Temps>/For"+Id+"/",
        paths  : "./.<Var_Temps>/For"+Id+"/Main/",
        TmpPaths  : "./.<Var_Temps>/For"+Id+"/Main/",
        FuncName  : "Main",
        TmpFuncName  : "Main",
        AllCodeStop : false,
        ResLock : false,
        Govm  : true,
        Func : &FuncResTmp{},
    }
}

var Pointer =make(map[string]string)
var Pointere =make(map[string]string)
var Pointerf =make(map[string]string)
var Pointeref =make(map[string]string)

var PointerLen int
var retuenlon int=0
var retuenpath string
var tmpcs string
var delmap = make(map[int]string)

func (ls *FileValue)RootCd(File string){
    ls.AllOverPaths+=File+"/"
}
func (ls *FileValue)SetPaths(path string){
    ls.paths = path
}
func (ls *FileValue)ReadPaths()string{
    return ls.paths
}
func (ls *FileValue)SetWgoId(Name string){
    ls.WgoIdName=Name;
}
func (St *FileValue)SetFunc(cdFile string){
    St.TmpPaths = St.paths
    St.TmpFuncName = St.FuncName
    St.paths = St.AllOverPaths+cdFile+"/"
    St.FuncName= cdFile
}

func CopyVmArray(start string,stop string){
    if Exists(start){
        if IsDir(start){
            Del_File(stop)
            Del_Dir(stop)
            Copy_Dir(start,stop)
        }
        if IsFile(start){
            Del_File(stop)
            Del_Dir(stop)
            Copy_File(start,stop)
        }
    }
}


func Copy_File(src, dst string) error {
    var err error
    var srcfd *os.File
    var dstfd *os.File
    var srcinfo os.FileInfo
    if srcfd, err = os.Open(src); err != nil {
        srcfd.Close()
        return err
    }
    defer srcfd.Close()
 
    if dstfd, err = os.Create(dst); err != nil {
        dstfd.Close()
        return err
    }
    defer dstfd.Close()
 
    if _, err = io.Copy(dstfd, srcfd); err != nil {
        return err
    }
    if srcinfo, err = os.Stat(src); err != nil {
        return err
    }
    return os.Chmod(dst, srcinfo.Mode())
}
 

func Copy_Dir(src string, dst string) error {
    var err error
    var fds []os.FileInfo
    var srcinfo os.FileInfo
 
    if srcinfo, err = os.Stat(src); err != nil {
        return err
    }
 
    if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
        New_File(dst)
    }
 
    if fds, err = ioutil.ReadDir(src); err != nil {
        return err
    }
    for _, fd := range fds {
        srcfp := path.Join(src, fd.Name())
        dstfp := path.Join(dst, fd.Name())
 
        if fd.IsDir() {
            if err = Copy_Dir(srcfp, dstfp); err != nil {
                fmt.Println(err)
            }
        } else {
            Copy_File(srcfp, dstfp)
        }
    }
    return nil
}
func Get_All_Array(pathname string) string {
    rd, _ := ioutil.ReadDir(pathname)
    res := ""
    for _, fi := range rd {
        if fi.IsDir() {
            if IsDir(pathname +"/" + fi.Name()){
                res+=fi.Name()+"=>array("
            }
            res+=Get_All_Array(pathname +"/" + fi.Name())
            if IsDir(pathname +"/" + fi.Name()){
                res+=")"
            }
        } else {
            res+=fi.Name()+"=>"+Read_File(pathname +"/" + fi.Name())+" "
        }
    }
    return res
}
func Exists(path string) bool{  //获取文件是否存在
    _, err := os.Stat(path)
    if err != nil {  
        if os.IsExist(err){  
            return true  
        }  
        return false  
    } 
    return true  
}  
func IsDir(path string) bool{    //目录
    s, err := os.Stat(path)  
    if err != nil {  
        return false  
    }  
    return s.IsDir()  
}  

func IsFile(path string) bool{    //文件
    return !IsDir(path)  
}

func Read_Array(file string,Vales *FileValue)string{   //vm读取变量
    tmp:=file
    Altp:=So_Array_Io(file)
    VarName:=Vales.FuncName+Altp[0]
    var Locks bool = false
    if len(Altp[0])>2{
        if Altp[0][0:2]=="0x"{
            VarName=Altp[0]
            Locks = true
        }
    }
    Lisr := ""
    for i:=1;i<=len(Altp)-1;i++{
        Lisr +=Altp[i]
    }
    file = So_Array_Stick(VarName,Vales)
    file_SHANGJI:=So_Array_Stick_SHANGJI(VarName,Vales)
    porinterTextNum:=""
    if len(tmp)>2{
        if tmp[0:2]=="0x"{
            Init:=So_Array_Io(tmp)
            Init[0]=Pointeref[Init[0]]
            file = ""
            for i:=0;i<=len(Init)-1;i++{
                file+="/"+Init[i]
            }
            file = file[1:]
            
            Inits:=So_Array_Io(tmp)
            Inits[0]=Pointeref[Inits[0]]
            file_SHANGJI = ""
            for i:=0;i<=len(Inits)-2;i++{
                file_SHANGJI+="/"+Inits[i]
            }
            porinterTextNum = Inits[len(Inits)-1]
            file_SHANGJI = file_SHANGJI[1:]
        }
    }
    if Exists(file){
        if IsDir(file){
            if Locks{
                return VarName+Lisr
            }
            return Pointer[VarName]+Lisr
        }else if IsFile(file){
            return Read_File(file)
        }
    }
    if Exists(file_SHANGJI){
        return string(Read_File(file_SHANGJI)[TypeInts(string(porinterTextNum[1:len(porinterTextNum)-1]))])
    }
    return "NULL"
}

func Read_Os_File(files error){  //错误处理
    file:=string(files.Error())
    r, _ := regexp.Compile("r(.*):")
    file = r.FindString(file)[1:]
    file = file[1:len(file)-1]
    Del_File(file)
    Del_Dir(file)
}
func Del_Dir(file string){  //删除目录
    os.RemoveAll(file)
}
func Del_File(file string){   //删除文件
    file=file[0:len(file)-1]
    os.Remove(file)
}
func Del_Files(file string){   //删除文件
    os.Remove(file)
}
func New_File(file string){   //无视报错，新建目录
    err := os.MkdirAll(file, 0666)
    if err != nil {
       Read_Os_File(err)
       New_File(file)
    }
}

func New_File_Var(file string,text string)string{    //新建变量
    filename := file
    f, _ := os.Create(filename)
    defer f.Close()
    content := text
    
    f.Write([]byte(content))
    if IsDir(file){
        os.RemoveAll(file)
        New_File_Var(file,text)
    }
    return ""
}

func Read_File(filepath string) string {    //读取文件
    fi, _ := os.Open(filepath)
    fd, _ := ioutil.ReadAll(fi)
    fi.Close()
    return string(fd)
}

func Del_Array(ar string,Vales *FileValue){    //删除变量
    ar = So_Array_Stick(Vales.FuncName+ar,Vales)
    Del_File(ar)
    Del_Dir(ar)
}

func So_Array_Io(Arrs string)(map[int]string){    //解析数组
    lock := 0
    start:=0
    avrs := make(map[int]string)
    avrs_l := 0
    for i:=0;i<=len(Arrs)-1;i++{
        if string(Arrs[i])=="]"{
            lock--
        }else if string(Arrs[i])=="["{
            if start==0{
                start = 1
                avrs_l++
            }
            lock++
        }
        avrs[avrs_l]+=string(Arrs[i])
        if lock==0 && start==1{
            avrs_l++
        }
    }
    return avrs
}

func So_Array_Stick(Arrs string,Vales *FileValue)string{    //数组转路径
    Vales.paths = Vales.paths
    Maps:=So_Array_Io(Arrs)
    file := Vales.paths
    for i:=0;i<=len(Maps)-2;i++{
        file+=Maps[i]+"/"
    }
    file = file+Maps[len(Maps)-1]
    return file
}

func So_Array_Stick_SHANGJI(Arrs string,Vales *FileValue)string{    //数组转路径上级
    Maps:=So_Array_Io(Arrs)
    file := Vales.paths
    for i:=0;i<=len(Maps)-3;i++{
        file+=Maps[i]+"/"
    }
    file = file+Maps[len(Maps)-1]
    return file
}

func Var_Pointer(VarName string,Vales *FileValue)string{   //声明指针
    VarNameFile := Vales.paths+VarName
    if _,ok:=Pointer[VarName];!ok{
        Pointerf[VarName]=VarNameFile
        Pointer[VarName]="0x"+TypeStrings(PointerLen)
        Pointeref["0x"+TypeStrings(PointerLen)]=VarNameFile
        Pointere["0x"+TypeStrings(PointerLen)]=VarName
        PointerLen++
    }
    return Pointer[VarName]
}

func AddArray(Arrs string,Var string,Vales *FileValue)string{    //添加数组
    Arrs = Vales.FuncName+Arrs
    
    if len(Var)>2{
        if Var[0:2]=="0x"{
            Init:=So_Array_Io(Var)
            Init[0]=Pointeref[Init[0]]
            Var = ""
            for i:=0;i<=len(Init)-1;i++{
                Var+="/"+Init[i]
            }
            Var = Var[1:]
            Var_Pointer(So_Array_Io(Arrs)[0],Vales)
            CopyVmArray(Var,So_Array_Stick(Arrs,Vales))
            return ""
        }
    }
    Maps:=So_Array_Io(Arrs)
    file := Vales.paths
    Var_Pointer(Maps[0],Vales)
    for i:=0;i<=len(Maps)-2;i++{
        file+=Maps[i]+"/"
    }
    New_File(file)
    New_File_Var(file+Maps[len(Maps)-1],Var)
    return ""
}

func VarNameGenerate(Code compile.Body_Struct_Run,Vales *FileValue)string{   //解析数组
    List := Code.Abrk
    Res := Code.Text
    for i:=0;i<=len(List)-1;i++{
        if List[i].Type==1{
            Res+="["+VarSoAll(List[i].Text,Vales)+"]"
        }else{
            break
        }
    }
    return Res
}

func VarNameGenerateClass(Code compile.Body_Struct_Run,Vales *FileValue)string{   //解析数组
    List := Code.Abrk
    Res := List[0].Text
    for i:=1;i<=len(List)-1;i++{
        if List[i].Type==1{
            Res+="["+VarSoAll(List[i].Text,Vales)+"]"
        }else{
            break
        }
    }
    return Res
}

/*MAP TO ARRAY*/
func CopyArrayStudio(Values map[string]interface{},Path string){
    for index, data := range Values{
        index = "["+index+"]"
        switch data.(type) {
            case map[string]interface{}:
                CopyArrayStudio(data.(map[string]interface{}),Path+index+"/")
            case string:
                New_File(Path+index)
                New_File_Var(Path+index,data.(string))
        }
    }
}


func VmEnd(){
    if VarRam{
        ramdisk.Del("./.<Var_Temps>")
    }else{
        Del_Dir("./.<Var_Temps>")
    }
    VmEndApis = true
    WebSocketWg.Wait()
}
func VmStart(){
    if VarRam{
        ramdisk.New(1024,"./.<Var_Temps>")
    }else{
        Del_Dir("./.<Var_Temps>")
    }
}

func GetFileModTime(path string) int64 {
    f, err := os.Open(path)
    if err != nil {
        return time.Now().Unix()
    }
    defer f.Close()

    fi, err := f.Stat()
    if err != nil {
        return time.Now().Unix()
    }

    return fi.ModTime().Unix()
}