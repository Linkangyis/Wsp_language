package gc

import(
    "io/ioutil"
    "path/filepath"
    "os"
)



func PathExists(path string)(bool){
    _ , err :=os.Stat(path)
    if err == nil{//文件或者目录存在
        return true
    }
    if os.IsNotExist(err){
        return false
    }
    return false
}

//获取目录dir下的文件大小
func readDir(dirPath string) int {
    if !PathExists(dirPath){
        return 0
    }
    var dirSize int
    flist, _ := ioutil.ReadDir(dirPath)
    for _, f := range flist {
        if f.IsDir() {
            dirSize= readDir(dirPath+"/"+f.Name()) + dirSize
        } else {
            dirSize= Read(dirPath+"/"+f.Name()) + dirSize + 4096
        }
    }
    return dirSize
}

func Read(dirPath string)int{
    Ls,_:=DirSize(dirPath)
    return int(Ls)
}
func DirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}
