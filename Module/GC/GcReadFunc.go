package gc

import(
    "io/ioutil"
)

//获取目录dir下的文件大小
func readDir(dirPath string) int {
    var dirSize int
    flist, _ := ioutil.ReadDir(dirPath)
    for _, f := range flist {
        if f.IsDir() {
            dirSize= readDir(dirPath+"/"+f.Name()) + dirSize
        } else {
            dirSize= int(f.Size()) + dirSize
        }
    }
    return dirSize
}
