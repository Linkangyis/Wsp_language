package memory

import (
	"Wsp/Module/Library"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

type MallocSTRING string

func init() {
	MemoryLoad = make([]sync.Map, 0)
	MemoryLoad = append(MemoryLoad, sync.Map{})
}

func thisTime() int64 {
	return time.Now().UnixNano()
}

func (this MallocSTRING) Open() *MemoryStruct {
	return OpenPointer(this)
}

func lenSyncMap(maps *sync.Map) int {
	length := 0
	maps.Range(func(k, v interface{}) bool {
		length++
		return true
	})
	return length
}

func Malloc() MallocSTRING {
	Lock.Lock()
	defer Lock.Unlock()
	CurrentLen := len(MemoryLoad) - 1

	CurrentLenString := lib.TypeStrings(CurrentLen)
	Map := MemoryLoad[CurrentLen]
	Pointer := CurrentLenString + "x" + Hex(lenSyncMap(&Map))
	MemoryLoad[CurrentLen].Store(lenSyncMap(&Map), &MemoryStruct{NewTime: thisTime()})
	if lenSyncMap(&Map)-1 >= 0xFF {
		MemoryLoad = append(MemoryLoad, sync.Map{})
	}

	OpenPointer(MallocSTRING(Pointer)).NewTime = thisTime()
	return MallocSTRING(Pointer)
}

func OpenPointer(Pointer MallocSTRING) *MemoryStruct {
	PageString := ""
	i := 0
	for {
		if Pointer[i] == 'x' {
			break
		}
		PageString += string(Pointer[i])
		i++
	}
	i = 0
	lock := false
	MapPageString := ""
	for {
		if i >= len(Pointer) {
			break
		}
		if lock {
			MapPageString += string(Pointer[i])
		}
		if Pointer[i] == 'x' {
			lock = true
		}
		i++
	}
	Page := lib.TypeInts(PageString)
	MapPage := Hex2Dec(MapPageString)
	value, ok := MemoryLoad[Page].Load(MapPage)
	if !ok {
		panic("内存溢出")
	}
	return value.(*MemoryStruct)
}

func Hex2Dec(val string) int {
	n, err := strconv.ParseUint(val, 16, 32)
	if err != nil {
		fmt.Println(err)
	}
	return int(n)
}

func Hex(ten int) string {
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
	for i := len(hex) - 1; i >= 0; i-- {
		if hex[i] >= 10 {
			hexStr = append(hexStr, fmt.Sprintf("%c", 'A'+hex[i]-10))
		} else {
			hexStr = append(hexStr, fmt.Sprintf("%d", hex[i]))
		}
	}
	HexRes := strings.Join(hexStr, "")
	if len(HexRes) == 1 {
		return "0" + HexRes
	}
	return HexRes
}

/*
	func SetValue[T any](Value T,this *MemoryStruct){
	    var TmpInterFace interface{}
	    TmpInterFace = Value
	    this.Value = TmpInterFace
	    this.ReadTime = thisTime()
	    this.SizeByte = int64(unsafe.Sizeof(Value))
	}
*/
func (this *MemoryStruct) SetValue(Value interface{}) {
	this.Value = Value
	this.SetTime = thisTime()
}
func (this *MemoryStruct) Read() *interface{} {
	this.ReadTime = thisTime()
	if this.Value == "<FREE>" {
		panic("This Memory in free")
	}
	return &this.Value
}
func (this *MemoryStruct) Free() {
	this.Value = "<FREE>"
	this = nil
}

func FreeAll() {
	MemoryLoad = make([]sync.Map, 0)
	MemoryLoad = append(MemoryLoad, sync.Map{})
}
