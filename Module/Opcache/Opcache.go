package op

import(
  "io/ioutil"
  "os"
  "encoding/json"
  "compress/zlib"
  "io"
  "bytes"
  "crypto/md5"
  "encoding/hex"
)

func PathExists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil {
        return true, nil
    }
    if os.IsNotExist(err) {
        return false, nil
    }
    return false, err
}


func Md5(str string) string {
    h := md5.New()
    h.Write([]byte(str))
    return hex.EncodeToString(h.Sum(nil))
}

func DoZlibCompress(src []byte) []byte {
    var in bytes.Buffer
    w := zlib.NewWriter(&in)
    w.Write(src)
    w.Close()
    return in.Bytes()
}
 
//进行zlib解压缩
func DoZlibUnCompress(compressSrc []byte) []byte {
    b := bytes.NewReader(compressSrc)
    var out bytes.Buffer
    r, _ := zlib.NewReader(b)
    io.Copy(&out, r)
    return out.Bytes()
}

func Opcaches_ADD(Buildse Opcodes,file string){
    fileName := file
    fileContent, _ := json.Marshal(Buildse)
    ioutil.WriteFile(fileName,DoZlibCompress(fileContent), 0666)
}

func Opcaches_Read(file string)(Opcodes){
    var stuRes Opcodes
    Text , _ := ioutil.ReadFile(file)
    Text = DoZlibUnCompress(Text)
    json.Unmarshal(Text, &stuRes)
    return stuRes
}