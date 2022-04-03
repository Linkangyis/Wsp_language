package main

import(
    "time"
    "strconv"
)

func H_Info()(map[int]string){
    info := make(map[int]string)
    info[0] = "Time"
    return info
}

func nulls(a string)(string){
    return a
}
func Time(a string)(string){
    nulls(a)
    return strconv.FormatInt(time.Now().UnixNano(),10)
}



//go build -buildmode=plugin -o time.so Time.go