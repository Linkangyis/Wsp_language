package array

import(
  "Wsp/Types"
  "os"
  "io/ioutil"
  "io"
  "path"
  "regexp"
  "fmt"
)
var paths string = "./Var_Temps/"
var Pointer =make(map[string]string)
var Pointere =make(map[string]string)
var PointerLen int
var retuenlon int=0
var retuenpath string
var tmpcs string
var delmap = make(map[int]string)

func Del_Dirs(path string){
    delmap[len(delmap)]=path
}
func Del_Dirl(){
    tmpcs=""
    for i:=0;i<=len(delmap)-1;i++{
        Del_Dir(delmap[i])
        delete(delmap,i)
    }
    
}
func Set_Paths(path string){
    paths = path
}
func Read_Paths()string{
    return paths
}
func Copy_Array(start string,stop string){
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
func Copy_ArrayVm(start string,stop string){
    start = So_Array_Stick_c(start)
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
func Set_Res(a int){
    retuenlon=a
}
func Set_Ress(a string){
    retuenpath=a
}
func Copy_File(src, dst string) error {
    var err error
    var srcfd *os.File
    var dstfd *os.File
    var srcinfo os.FileInfo
 
    if srcfd, err = os.Open(src); err != nil {
        return err
    }
    defer srcfd.Close()
 
    if dstfd, err = os.Create(dst); err != nil {
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
    Head_VarName:=So_Array_Io(file)[0]
    for i:=1;i<=len(So_Array_Io(file))-1;i++{
        tmpcs+=So_Array_Io(file)[i]
    }
    if file==""{
        return "NULL"
    }
    filetemp:=So_Array_Stick_b(file)
    filetemplen:=So_Array_Stick_b_len(file)
    file = So_Array_Stick(file)
    if IsDir(file){
        //return "array("+Get_All_Array(file)+")"
        return Var_Pointer(Head_VarName)
    }
    if !Exists(file){
        if Exists(filetemp){
            if IsFile(filetemp){
                return string(Read_File(filetemp)[filetemplen])
            }
        }
        return "NULL"
    }
    if IsFile(file){
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
func So_Array_Stick_b(Arrs string)string{
    Maps:=So_Array_Io(Arrs)
    file := paths
    for i:=0;i<=len(Maps)-2;i++{
        file+=Maps[i]+"/"
    }
    file = file[0:len(file)-1]
    return file
}
func So_Array_Stick_c(Arrs string)string{
    Maps:=So_Array_Io(Arrs)
    file:=""
    if retuenlon==0{
        file = "./Var_Temps/"
    }else{
        file = retuenpath
    }
    for i:=0;i<=len(Maps)-1;i++{
        file+=Maps[i]+"/"
    }
    if tmpcs!=""{
        file+=tmpcs+"/"
        tmpcs=""
    }
    file = file[0:len(file)-1]
    return file
}
func So_Array_Stick_b_len(Arrs string)int{
    Maps:=So_Array_Io(Arrs)
    if len(Arrs)==1{
        return 0
    }
    reslen := types.Ints(Maps[len(Maps)-1][1:len(Maps[len(Maps)-1])-1])
    return reslen
}
func Var_Pointer(VarName string)string{
    if _,ok:=Pointer[VarName];!ok{
        Pointer[VarName]="0x"+types.Strings(PointerLen)
        Pointere["0x"+types.Strings(PointerLen)]=VarName
        PointerLen++
    }
    return Pointer[VarName]
}
func Add_Array(Arrs string,Var string)string{
    if len(Var)>2{
        if Var[0:2]=="0x"{
            Copy_ArrayVm(Pointere[Var],Arrs)
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