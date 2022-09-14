package center

import(
  "fmt"
  "strings"
)

func New_Memory()string{
    T:=Num_Hex(Memory_Id)
    Memory_Id++
    LenTmp:=len(T)
    for i:=0;i<=7-LenTmp;i++{
        T="0"+T
    }
    Memory_From_Map[len(Memory_From_Map)]="0x"+T
    return "0x"+T
}

func New_MemoryFunc()string{
    T:=Num_Hex(Memory_Id)
    Memory_Id++
    LenTmp:=len(T)
    for i:=0;i<=7-LenTmp;i++{
        T="0"+T
    }
    Memory_From_Map[len(Memory_From_Map)]="0F"+T
    return "0F"+T
}

func Num_Hex(ten int) string {
    m := 0
    hex := make([]int, 0)
    for {
        m = ten % 16
        ten = ten / 16
        if ten == 0 {
            hex = append(hex, m)
            break
        }
        hex = append(hex, m)
    }
    hexStr := []string{}
    for i:=len(hex)-1;i>=0;i--{
        if hex[i] >= 10 {
            hexStr = append(hexStr, fmt.Sprintf("%c", 'A'+hex[i]-10))
        } else {
            hexStr = append(hexStr, fmt.Sprintf("%d", hex[i]))
        }
    }
    return strings.Join(hexStr, "")
}
