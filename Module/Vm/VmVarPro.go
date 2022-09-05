package vm

import(
  "os"
  "io/ioutil"
  "io"
  "path"
  "regexp"
  "fmt"
  "Wsp/Compile"
)
var FILE string = "./.<Var_Temps>/"
var AllOverPaths string = "./.<Var_Temps>/"
var paths string = "./.<Var_Temps>/Main/"
var TmpPaths string = "./.<Var_Temps>/Main/"
var FuncName string = "Main"
var TmpFuncName  = "Main"

var Pointer =make(map[string]string)
var Pointere =make(map[string]string)
var Pointerf =make(map[string]string)
var Pointeref =make(map[string]string)

var PointerLen int
var retuenlon int=0
var retuenpath string
var tmpcs string
var delmap = make(map[int]string)

func RootCd(File string){
    AllOverPaths+=File+"/"
}

func DelDirs(path string){
    delmap[len(delmap)]=path
}
func DelDirl(){
    tmpcs=""
    for i:=0;i<=len(delmap)-1;i++{
        Del_Dir(delmap[i])
        delete(delmap,i)
    }
    
}
func SetPaths(path string){
    paths = path
}
func ReadPaths()string{
    return paths
}
func CopyArray(start string,stop string){
    start = FuncName+start
    stop = FuncName+stop
    Var_Pointer(So_Array_Io(stop)[0])
    start = So_Array_Stick(start)
    stop = So_Array_Stick(stop)
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
            if err = Copy_File(srcfp, dstfp); err != nil {
                fmt.Println(err)
            }
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
func Exists(path string) bool{
    _, err := os.Stat(path)    //os.Stat获取文件信息  
    if err != nil {  
        if os.IsExist(err){  
            return true  
        }  
        return false  
    }  
    return true  
}  
func IsDir(path string) bool{  
    s, err := os.Stat(path)  
    if err != nil {  
        return false  
    }  
    return s.IsDir()  
}  

func IsFile(path string) bool{  
    return !IsDir(path)  
}

func Read_Array(file string)string{
    tmp:=file
    Altp:=So_Array_Io(file)
    VarName:=FuncName+Altp[0]
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
    file = So_Array_Stick(VarName)
    
    if len(tmp)>2{
        if tmp[0:2]=="0x"{
            Init:=So_Array_Io(tmp)
            Init[0]=Pointeref[Init[0]]
            file = ""
            for i:=0;i<=len(Init)-1;i++{
                file+="/"+Init[i]
            }
            file = file[1:]
        }
    }
    if IsDir(file){
        if Locks{
            return VarName+Lisr
        }
        return Pointer[VarName]+Lisr
    }else if IsFile(file){
        return Read_File(file)
    }
    
    return "NULL"
}

func Read_Os_File(files error){
    file:=string(files.Error())
    r, _ := regexp.Compile("r(.*):")
    file = r.FindString(file)[1:]
    file = file[1:len(file)-1]
    Del_File(file)
    Del_Dir(file)
}
func Del_Dir(file string){
    os.RemoveAll(file)
}
func Del_File(file string){
    file=file[0:len(file)-1]
    os.Remove(file)
}
func New_File(file string){
   err := os.MkdirAll(file, 0666)
   if err != nil {
      Read_Os_File(err)
      New_File(file)
   }
}

func New_File_Var(file string,text string)string{
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

func Read_File(filepath string) string {
   fi, _ := os.Open(filepath)
   fd, _ := ioutil.ReadAll(fi)
   fi.Close()
   return string(fd)
}

func Del_Array(ar string){
    ar = So_Array_Stick(ar)
    Del_File(ar)
    Del_Dir(ar)
}

func So_Array_Io(Arrs string)(map[int]string){
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

func So_Array_Stick(Arrs string)string{
    Maps:=So_Array_Io(Arrs)
    file := paths
    for i:=0;i<=len(Maps)-2;i++{
        file+=Maps[i]+"/"
    }
    file = file+Maps[len(Maps)-1]
    return file
}

func Var_Pointer(VarName string)string{
    VarNameFile := paths+VarName
    if _,ok:=Pointer[VarName];!ok{
        Pointerf[VarName]=VarNameFile
        Pointer[VarName]="0x"+TypeStrings(PointerLen)
        Pointeref["0x"+TypeStrings(PointerLen)]=VarNameFile
        Pointere["0x"+TypeStrings(PointerLen)]=VarName
        PointerLen++
    }
    return Pointer[VarName]
}

func AddArray(Arrs string,Var string)string{
    Arrs = FuncName+Arrs
    if len(Var)>2{
        if Var[0:2]=="0x"{
            Init:=So_Array_Io(Var)
            Init[0]=Pointeref[Init[0]]
            Var = ""
            for i:=0;i<=len(Init)-1;i++{
                Var+="/"+Init[i]
            }
            Var = Var[1:]
            Var_Pointer(Arrs)
            CopyVmArray(Var,So_Array_Stick(Arrs))
            return ""
        }
    }
    Maps:=So_Array_Io(Arrs)
    file := paths
    Var_Pointer(Maps[0])
    for i:=0;i<=len(Maps)-2;i++{
        file+=Maps[i]+"/"
    }
    New_File(file)
    New_File_Var(file+Maps[len(Maps)-1],Var)
    return ""
}

func SetFunc(cdFile string){
    TmpPaths = paths
    TmpFuncName = FuncName
    paths = AllOverPaths+cdFile+"/"
    FuncName= cdFile
}

func VarNameGenerate(Code compile.Body_Struct_Run)string{
    List := Code.Abrk
    Res := Code.Text
    for i:=0;i<=len(List)-1;i++{
        if List[i].Type==1{
            Res+="["+VarSoAll(List[i].Text)+"]"
        }else{
            break
        }
    }
    return Res
}

func VmEnd(){
    Del_Dir("./.<Var_Temps>")
}

