package jit

import(
  "unsafe"
  "fmt"
  "syscall"
)

func Asm(Code []uint16){
    executablePrintFunc, err := syscall.Mmap(
        -1,
        0,
        len(Code)*2,
        syscall.PROT_READ|syscall.PROT_WRITE|syscall.PROT_EXEC,
        syscall.MAP_PRIVATE|syscall.MAP_ANONYMOUS)
    if err != nil {
        fmt.Printf("内存映射错误 JITError: %v", err)
    }
    j := 0
    for i := range Code {
        executablePrintFunc[j] = byte(Code[i] >> 8)
        executablePrintFunc[j+1] = byte(Code[i])
        j = j + 2
    }
    type Run func()
    unsafePrintFunc := (uintptr)(unsafe.Pointer(&executablePrintFunc))
    Runs := *(*Run)(unsafe.Pointer(&unsafePrintFunc))
    Runs()
}

//暂时不会做