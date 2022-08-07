package crun

type node struct {
    value interface{}
    next  *node
}

// 栈的链式结构实现
type Stack struct {
    top    *node
    length int
}
