package main

import ()

func main() {
	var Vars = make(map[string]interface{})
	Vars["a"] = "1"
	aaa("")

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
