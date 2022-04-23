package main

import (
	"io/ioutil"
	"os"
	"plugin"
	"strings"
)

func main() {
	var Vars = make(map[string]interface{})
	Vars["a"] = "1"
	aaa("")
	DLS_So_Start()
	So_func_map["Tests"].(func(string) string)(Vars["a"].(string))

}
func aaa(a string) string {
	var Vars = make(map[string]interface{})
	Vars["a"] = a
	for z := 0; z <= 100; z = z + 1 {
		Vars["z"] = z
		for i := 0; i <= 100; i = i + 1 {
			Vars["i"] = i
		}
	}
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
