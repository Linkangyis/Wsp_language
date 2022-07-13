package main

import (
    "strconv"
    "Wsp/WVM"
)

type StackNode struct {
    Data interface{}
    next *StackNode
}

type LinkStack struct {
    top *StackNode
    Count int
}

func (this *LinkStack) Init() {
    this.top = nil
    this.Count = 0
}

func (this *LinkStack) Push(data interface{}) {
    var node *StackNode = new(StackNode)
    node.Data = data
    node.next = this.top
    this.top = node
    this.Count++
}

func (this *LinkStack) Pop() interface{} {
    if this.top == nil {
        return nil
    }
    returnData := this.top.Data
    this.top = this.top.next
    this.Count--
    return returnData
}

//Look up the top element in the stack, but not pop.
func (this *LinkStack) LookTop() interface{} {
    if this.top == nil {
        return nil
    }
    return this.top.Data
}

var str string

func H_Info()(map[int]string){
    info := make(map[int]string)
    info[0] = "Crun"
    return info
}

func Crun(a string)(string){
    str_arr,_:=vm.Parameter_processing(a)
    add_num:=str_arr[0]
    s2 := strconv.FormatFloat(Count(add_num),'f',-1,64)//float64
    return s2
}

func Count(data string) float64 {
    var arr []string = generateRPN(data)
    return calculateRPN(arr)
}

func calculateRPN(datas []string) float64 {
    var stack LinkStack
    stack.Init()
    for i := 0; i < len(datas); i++ {
        if isNumberString(datas[i]) {
            if f, err := strconv.ParseFloat(datas[i], 64); err != nil {
                panic("operatin process go wrong.")
            } else {
                stack.Push(f)
            }
        } else {
            p1 := stack.Pop().(float64)
            p2 := stack.Pop().(float64)
            p3 := normalCalculate(p2, p1, datas[i])
            stack.Push(p3)
        }
    }
    res := stack.Pop().(float64)
    return res
}

func normalCalculate(a,b float64, operation string) float64 {
    switch operation{
    case "*":
        return a * b
    case "-":
        return a - b
    case "+":
        return a + b
    case "/":
        return a / b
    default:
        panic("invalid operator")
    }
}

func generateRPN(exp string) []string {

    var stack LinkStack
    stack.Init()

    var spiltedStr []string = convertToStrings(exp)
    var datas []string

    for i := 0; i < len(spiltedStr); i++ { // 遍历每一个字符
        tmp := spiltedStr[i] //当前字符

        if !isNumberString(tmp) { //是否是数字
            // 四种情况入栈
            // 1 左括号直接入栈
            // 2 栈内为空直接入栈
            // 3 栈顶为左括号，直接入栈
            // 4 当前元素不为右括号时，在比较栈顶元素与当前元素，如果当前元素大，直接入栈。
            if tmp == "(" ||
                stack.LookTop() == nil || stack.LookTop().(string) == "(" ||
                ( compareOperator(tmp, stack.LookTop().(string)) == 1 && tmp != ")" ) {
                stack.Push(tmp)
            } else { // ) priority
                if tmp == ")" { //当前元素为右括号时，提取操作符，直到碰见左括号
                    for {
                        if pop := stack.Pop().(string); pop == "(" {
                            break
                        } else {
                            datas = append(datas, pop)
                        }
                    }
                } else { //当前元素为操作符时，不断地与栈顶元素比较直到遇到比自己小的（或者栈空了），然后入栈。
                    for {
                        pop := stack.LookTop()
                        if pop != nil && compareOperator(tmp, pop.(string)) != 1 {
                            datas = append(datas, stack.Pop().(string))
                        } else {
                            stack.Push(tmp)
                            break
                        }
                    }
                }
            }

        } else {
            datas = append(datas, tmp)
        }
    }

    //将栈内剩余的操作符全部弹出。
    for {
        if pop := stack.Pop(); pop != nil {
            datas = append(datas, pop.(string))
        } else {
            break
        }
    }
    return datas
}

// if return 1, o1 > o2.
// if return 0, o1 = 02
// if return -1, o1 < o2
func compareOperator(o1, o2 string) int {
    // + - * /
    var o1Priority int
    if o1 == "+" || o1 == "-" {
        o1Priority = 1
    } else {
        o1Priority = 2
    }
    var o2Priority int
    if o2 == "+" || o2 == "-" {
        o2Priority = 1
    } else {
        o2Priority = 2
    }
    if o1Priority > o2Priority {
        return 1
    } else if o1Priority == o2Priority {
        return 0
    } else {
        return -1
    }
}

func isNumberString(o1 string) bool {
    if o1 == "+" || o1 == "-" || o1 == "*" || o1 == "/" || o1 == "(" || o1 == ")" {
        return false
    } else {
        return true
    }
}

func convertToStrings(s string) []string {
    var strs []string
    bys := []byte(s)
    var tmp string
    for i := 0; i < len(bys); i++ {
        if !isNumber(bys[i]) {
            if tmp != "" {
                strs = append(strs, tmp)
                tmp = ""
            }
            strs = append(strs, string(bys[i]))
        } else {
            tmp = tmp+string(bys[i])
        }
    }
    strs = append(strs, tmp)
    return strs
}

func isNumber(o1 byte) bool {
    if o1 == '+' || o1 == '-' || o1 == '*' || o1 == '/' || o1 == '(' || o1 == ')' {
        return false
    } else {
        return true
    }
}
//go build -buildmode=plugin -o crun.so Crun.go