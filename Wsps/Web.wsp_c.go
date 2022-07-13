package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"plugin"
	"strings"
)

func main() {
	DLS_So_Start()
	var Vars = make(map[string]interface{})
	Vars["a"] = "114514"
	Vars["a"] = "10086"
	Run(`""`)
	TS3("")
//	Starts(`""`)
	Exit(`""`)

}
func Exit(a string) string {
	var Vars = make(map[string]interface{})
	Vars["a"] = So_func_map["Input"].(func(string) string)(`"输入Exit关闭服务:"`)
	if Vars["a"].(string) != `"Exit"` {
		Exit(`""`)
	}
	return "1"
}
func A(a string) string {
	var Vars = make(map[string]interface{})

	fmt.Println("测试服务A")

	fmt.Println(So_func_map["Time"].(func(string) string)(`""`))
	Vars["a"] = So_func_map["Web_Print"].(func(string) string)(Vars["a"].(string))
	for i := 0; i <= 100; i = i + 1 {
		Vars["i"] = i
		So_func_map["Web_Print"].(func(string) string)(Vars["i"].(string))
	}
	So_func_map["Web_Print"].(func(string) string)(`"测试服务A结束"`)
	return a
}
func B(a string) string {
	var Vars = make(map[string]interface{})
	So_func_map["Web_Header_Set"].(func(string) string)(`"Content-Type","text/html; charset=utf-8"`)

	fmt.Println("测试服务B")

	fmt.Println(So_func_map["Time"].(func(string) string)(`""`))
	Vars["a"] = So_func_map["Web_Print"].(func(string) string)(Vars["a"].(string))
	return a
}
func C(a string) string {
	var Vars = make(map[string]interface{})

	So_func_map["Web_Header_Set"].(func(string) string)(`"Content-Type","text/html; charset=utf-8"`)

	fmt.Println("测试服务C")

	fmt.Println(So_func_map["Time"].(func(string) string)(`""`))
	Vars["a"] = So_func_map["Web_Print"].(func(string) string)(Vars["a"].(string))
	return a
}
func TS2(a string) string {

	fmt.Println("服务A已开启")
	So_func_map["Web_Start"].(func(string) string)(`"9968"`)
	return a
}
func Run(a string) string {

	So_func_map["New_WebFiles"].(func(string) string)(`"Public","./"`)
	So_func_map["New_Web"].(func(string) string)(`"AS","A"`)
	So_func_map["New_Web"].(func(string) string)(`"","A"`)
	So_func_map["New_Web"].(func(string) string)(`"B","B"`)
	So_func_map["New_Web"].(func(string) string)(`"B/C","C"`)
	So_func_map["New_WebFiles"].(func(string) string)(`"P","../Ext"`)
	So_func_map["New_WebFiles"].(func(string) string)(`"Publics","../"`)
	So_func_map["New_Web"].(func(string) string)(`"B/Cs","B"`)
	return a
}
func TS3(a string) string {

	fmt.Println("服务B已开启")
	So_func_map["Web_Start"].(func(string) string)(`"9958"`)
	return a
}
func TS4(a string) string {

	fmt.Println("服务C已开启")
	So_func_map["Web_Start"].(func(string) string)(`"9948"`)
	return ""
}
func res(a string) string {
	var Vars = make(map[string]interface{})
	Vars["a"] = a
	So_func_map["Multithreading"].(func(string) string)(Vars["a"].(string))
	return ""
}
func Starts(a string) string {


	res(`"TS2"`)
	So_func_map["Sleep"].(func(string) string)("1")
	res(`"TS3"`)
	res(`"TS4"`)
	return ""
}

var So_func_map = make(map[string]plugin.Symbol)

func DLS_So_Start() {
	data, _ := ioutil.ReadFile(os.Getenv("WSPPATH") + "/wsp.ini")
	inis := strings.Split(string(data), "\n")
	for i := 0; i <= len(inis)-1; i++ {
		iniss := strings.Split(inis[i], "=")
		if iniss[0] == "extension" {
			So_DLL_vm(iniss[1])
		}
	}
}
func So_DLL_vm(file string) map[string]plugin.Symbol {
	p, _ := plugin.Open(file)
	add, _ := p.Lookup("H_Info")
	funcmaps := add.(func() map[int]string)()

	for i := 0; i <= len(funcmaps)-1; i++ {
		add, _ = p.Lookup(funcmaps[i])
		So_func_map[funcmaps[i]] = add
	}
	return So_func_map
}
