package main

import (
	"fmt"
	"time"
)

type Error struct {
	ErrorType int
	ErrorRes  interface{}
}

var StopChan = make(chan int, 1)

func endThread() {
	panic(Error{1, "Exit"})
}

func newThread() {
	defer func() {
		if r := recover(); r != nil {
			if r.(Error).ErrorType == 1 {
				fmt.Println("信道关闭")
			}
		}
	}()
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("信道")
		if len(StopChan) > 0 {
			endThread()
		}
		/*
			select {
			default:
			case <-StopChan:
				endThread()
			}
		*/
	}
}

func main() {
	go newThread()
	time.Sleep(5 * time.Second)
	StopChan <- 20
	for {
	}
}
