package array

import(
  "os"
  "io/ioutil"
  "io"
  "path"
  "regexp"
  "fmt"
)
var paths string = "./Var_Temps/"
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
    file = So_Array_Stick(file)
    if !Exists(file){
        return "NULL"
    }
    if IsDir(file){
        //return "array("+Get_All_Array(file)+")"
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

func New_File_Var(file string,text string){
    filename := file
    f, _ := os.Create(filename)
    defer f.Close()
    content := text
    f.Write([]byte(content))
    if IsDir(file){
        os.RemoveAll(file)
        New_File_Var(file,text)
    }
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
func Add_Array(Arrs string,Var string){
    Maps:=So_Array_Io(Arrs)
    file := paths
    for i:=0;i<=len(Maps)-2;i++{
        file+=Maps[i]+"/"
    }
    New_File(file)
    New_File_Var(file+Maps[len(Maps)-1],Var)
}