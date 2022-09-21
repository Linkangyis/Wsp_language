package consts

import (
    "path/filepath"
)


type WspLoader struct{
    WspFile string
    WspDir string
    WspName string
}

var WspConst WspLoader

func (Const *WspLoader)SetWspFile(Files string){
    Dir, FileName := filepath.Split(Files)
    Const.WspFile=Files
    Const.WspDir=Dir
    Const.WspName=FileName
}