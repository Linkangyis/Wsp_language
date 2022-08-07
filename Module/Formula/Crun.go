package crun

import (
    "bytes"
    "fmt"
    "strconv"
)

func NewStack() *Stack {
    return &Stack{nil, 0}
}

func (s *Stack) Len() int {
    return s.length
}

// 获取栈顶元素
func (s *Stack) Peek() interface{} {
    if s.length == 0 {
        return nil
    }
    return s.top.value
}

// 弹出栈顶
func (s *Stack) Pop() interface{} {
    if s.length == 0 {
        return nil
    }
    n := s.top
    s.top = n.next
    s.length--
    return n.value
}

// 压栈
func (s *Stack) Push(value interface{}) {
    n := &node{value, s.top}
    // 从头部插入
    s.top = n
    s.length++
}


func negative(Num []string)[]string{ //负数解析
    Tmp:=[]string{"1","*"}
    Tmp=append(Tmp,Num...)
    Num=Tmp
    ResMap := make(map[int]string)
    lens := 0
    for i:=2;i<=len(Num)-1;i++{
        Nums:=Num[i]
        if (Num[i-1]=="-") && isSign(Num[i-2]){
            lens--
            ResMap[lens]=Num[i-1]+Nums
        }else if (Num[i-1]=="+") && isSign(Num[i-2]){
            lens--
            ResMap[lens]=Nums
        }else{
            ResMap[lens]=Nums
        }
        lens++
    }
    Res:=[]string{}
    for i:=0;i<=len(ResMap)-1;i++{
        Res=append(Res,ResMap[i])
    }
    return Res
}

func PostfixCRun(str string)[]string{
    exp := toExp(str)
    exp = negative(exp)
    postfixExp := toPostfix(exp)
    return postfixExp
}
func RunNums(Run []string)float64{
    return calValue(Run)
}

// 将数字和操作符转为字符串数组
func toExp(str string) []string {
    s := make([]string, 0)
    var t bytes.Buffer
    n := 0 // 用于判断括号是否成对
    for _, r := range str {
        if r == ' ' {
            // 去掉空格
            continue
        }
        if isDigit(r) {
            // 是数字 就写到缓存中
            t.WriteRune(r)
        } else {
            rs := string(r)
            if !isSign(rs) {
                panic("unknown sign: " + rs)
            }
            if t.Len() > 0 {
                // 遇到符号 把缓存中的数字 输出为数
                // 例如 将缓存中的 ['1', '2', '3'] 输出为 "123"
                s = append(s, t.String())
                t.Reset()
            }
            s = append(s, rs)
            if r == '(' {
                n++
            } else if r == ')' {
                n--
            }
        }
    }
    if t.Len() > 0 {
        // 最后一个操作符后面的数字 如果最后一个操作符是 ")" 那么 t.Len() 为0
        s = append(s, t.String())
    }
    if n != 0 {
        panic("the number of '(' is not equal to the number of ')' ")
    }
    return s
}

func printExp(exp []string) {
    for _, s := range exp {
        fmt.Print(s, " ")
    }
    fmt.Println()
}

// 是否数字
func isDigit(r rune) bool {
    if r >= '0' && r <= '9' {
        return true
    }
    if string(r)=="."{
        return true
    }
    return false
}

// 是否符号
func isSign(s string) bool {
    switch s {
    case "+", "-", "*", "/", "%", "(", ")":
        return true
    default:
        return false
    }
}

// 中缀表达式转后缀表达式
func toPostfix(exp []string) []string {
    result := make([]string, 0)
    s := NewStack()
    for _, str := range exp {
        if isSign(str) {
            // 若是符号
            if str == "(" || s.Len() == 0 {
                // "(" 或者 栈为空 直接进栈
                // 括号中的计算 需要单独处理 相当于一个新的上下文
                // 如果栈为空 需要先进栈 和后续操作符比较优先级之后 才能决定计算顺序
                s.Push(str)
            } else {
                if str == ")" {
                    // 若为 ")" 依次弹出栈顶元素并输出 直到遇到 "("
                    for s.Len() > 0 {
                        if s.Peek().(string) == "(" {
                            s.Pop()
                            break
                        }
                        result = appendStr(result, s.Pop().(string))
                    }
                } else {
                    // 判断其与栈顶符号的优先级
                    // 如果栈顶是 "(" 说明是新的上下文 不能相互比较优先级
                    for s.Len() > 0 && s.Peek().(string) != "(" && signCompare(str, s.Peek().(string)) <= 0 {
                        // 当前符号的优先级 不大于栈顶元素 弹出栈顶元素并输出
                        // 优先级高的操作 需要先计算
                        // 优先级相同 因为栈中的操作是先放进去的 也需要先计算
                        result = appendStr(result, s.Pop().(string))
                    }
                    // 当前符号入栈
                    s.Push(str)
                }
            }
        } else {
            // 若是数字就输出
            result = appendStr(result, str)
        }
    }
    for s.Len() > 0 {
        result = appendStr(result, s.Pop().(string))
    }
    return result
}

func appendStr(slice []string, str string) []string {
    if str == "(" || str == ")" {
        // 后缀表达式 不包含括号
        return slice
    }
    return append(slice, str)
}

// 比较符号优先级
func signCompare(a, b string) int {
    return getSignValue(a) - getSignValue(b)
}

// 优先级越高 值越大
func getSignValue(a string) int {
    switch a {
    case "(", ")":
        return 2
    case "*", "/","%":
        return 1
    default:
        return 0
    }
}

// 通过后缀表达式 计算值
func calValue(exp []string) float64 {
    s := NewStack()
    for _, str := range exp {
        if isSign(str) {
            // 如果是符号 弹出栈顶的两个元素 进行计算
            // 因为栈结构先进后出 所以先弹出b
            b := getInt(s)
            a := getInt(s)
            var n float64
            switch str {
            case "+":
                n = a + b
            case "-":
                n = a - b
            case "*":
                n = a * b
            case "/":
                n = a / b
            case "%":
                n = float64(int64(a) % int64(b))
            }
            // 计算结果压栈
            s.Push(n)
        } else {
            // 数字直接压栈 也可以在这里做类型转换 使栈中的值均为int
            s.Push(str)
        }
    }
    // 栈顶元素 为最终结果
    return getInt(s)
}

// 弹出栈顶元素 并转为int
func getInt(s *Stack) float64 {
    v := s.Pop()
    switch v.(type) {
    case float64: // push进去的计算结果为int
        return v.(float64)
    case string: // exp中的数据为string
        if i, err := strconv.ParseFloat(v.(string),64); err != nil {
            panic(err)
        } else {
            return float64(i)
        }
    }
    panic(fmt.Sprintf("unknown value type: %T", v))
}