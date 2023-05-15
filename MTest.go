package main

import (
	memory "Wsp/Module/Memory"
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i <= 512; i++ {
		wg.Add(1)
		go func() {
			memory.Malloc()
			wg.Done()
		}()
	}
	wg.Wait()
	a := memory.Malloc()
	a.Open().SetValue("你好")
	fmt.Println((*a.Open().Read()).(string))
	a.Open().SetValue("不好")
	fmt.Println((*a.Open().Read()).(string))
	memory.FreeAll()
}
