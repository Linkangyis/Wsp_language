package ramdisk

import (
    "os"
    "Wsp/Module/RamDisk/ramdisk"
    "Wsp/Module/RamDisk/datasize"
)
func New(Size uint64,File string){
    os.Mkdir(File, os.ModePerm)
    var opts ramdisk.Options
    opts.Size = Size * datasize.MB
    opts.MountPath = File

    _, err := ramdisk.Create(opts)
    if err != nil {
        os.Exit(1)
    }
}

func Del(File string){
    device := File
    err := ramdisk.Destroy(device)
    if err != nil {
        os.RemoveAll(File)
        os.Exit(1)
    }
    os.RemoveAll(File)
}