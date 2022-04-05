package token

import(
  "strconv"
  "strings"
  "io/ioutil"
)

func token_str_ifs(str string)(string){      //输出函数替换切片  可以自行定义 T_KH只是演示  代表空行
    var maps = map[string]string{}
    maps["\n"] = "";
    //maps["\\"] = "";
    
    if v, ok := maps[str]; ok {      //判断是否有对应值
        return v
    }else{
        return str
    }
}
func token_str_if(str string)(string){      //输出函数替换切片  可以自行定义 T_KH只是演示  代表空行
    var maps = map[string]string{}
    maps["\n"] = "";
    
    if v, ok := maps[str]; ok {      //判断是否有对应值
        return v
    }else{
        return str
    }
}

func token(str string,line int)([4]string){//判断token主题内容是否为字符串，如为字符串则返回T_TEXT token 否则匹配config.go设置的token进行判断
    var tokens = map[string]int{}
    tokens = token_map()
    
    var tokens_name = map[string]string{}
    tokens_name = token_text_map()
    
    if v, ok := tokens[str]; ok {         //判断是否存在于config.go中
        return [4]string{strconv.Itoa(v),token_str_if(str),tokens_name[str],strconv.Itoa(line)}    //存在  获取类型
    }else{
        return [4]string{strconv.Itoa(0),str,"T_TEXT",strconv.Itoa(line)}    //不存在 类型直接定义为sting
    }
}

var Files string

func Wsp_File_P()(string){
    return Files
}
func Wsp_Lexical(file string)(map[int][4]string){
                                                    /*读取文件 并放在data变量  ---start*/
    data, _ := ioutil.ReadFile(file)
    code :=string(data)+"\n"+"\n\n\n"
    lenf:=strings.Count(code,"function")
    for i:=0;i<=lenf-1;i++{
        code+="\n"
    }
    code=Code_Notes(code)
    Files = file
                                                    /*读取文件 并放在data变量   --- end*/
                                                    /*设置全局循环数值         ---start*/
    num := 0;           //该opcode指令集位置
    num_TEXTS :=0       //需要删除掉的字符串数组数量
    lock_TEXT := 0      //字符串自增锁
    line :=1            //行数叠加
                                                    /*设置全局循环数值           ---end*/
                                                    /*token分析开始            ---start*/
    var return_map = map[int][4]string{}    //定义token_map  类型为map
    code_wsp := strings.Split(code,"")      //将字符串转化为数组
    code_len := len(code_wsp)-1             //统计数组总量并减一
    
    for i:=0;i<=code_len;i++{               //开启循环
        tk := token(string(code_wsp[i][len(code_wsp[i])-1]),line);   //获取opcode主题最后一个字符串的token类型  只作为缓存作用
        if tk[0] == strconv.Itoa(0){     //判断类型
            if i!=len(code_wsp)-1 {      //判断是否不为最后一个字符
                tmp := code_wsp[i+1]     //把字符串数组中的下一个数组保存至tmp变量
                code_wsp[i+1]=code_wsp[i]+code_wsp[i+1]    //将两个变量拼接
                num_TEXTS++    //需要删除的字符串数组数量加一
                if token(string(code_wsp[i+1][len(code_wsp[i+1])-1]),line)[0] != strconv.Itoa(0){    //判断下一个字符是否不为字符串类型
                    code_wsp[i+1] = tmp
                    for z:=i-num_TEXTS+1;z<=i-1;z++{
                        code_wsp[z]=""          //根据num_TEXTS破坏字符串
                    }
                    num_TEXTS = 0               //需要破坏的数组归零
                    lock_TEXT = 1               //本次字符串循环结束，激活文本锁
                }
            }
        }else{
            code_wsp[i]=string(code_wsp[i][len(code_wsp[i])-1])   //清空上次opcode留存的字符串，只保留本次opcode关键字
        }
        if lock_TEXT==1 {     //判断文本锁状态
            return_map[num]=token(code_wsp[i],line)       //写入opcode
            num++                                         //opcode位置加一
            lock_TEXT = 0                                 //关闭文本锁
        }else if token(string(code_wsp[i][len(code_wsp[i])-1]),line)[0] != strconv.Itoa(0){    //否则判断最后一个字符的类型是否不为字符串
            return_map[num]=token(code_wsp[i],line)       //写入opcode
            num++                                         //opcode位置加一
        }
        if code_wsp[i]=="\n"{                             //统计行数
            line++
        }
    }
    return return_map
                                                    /*token分析结束              ---end*/
}

func token_funcs(str string,line int)([4]string){//判断token主题内容是否为字符串，如为字符串则返回T_TEXT token 否则匹配config.go设置的token进行判断
    var tokens = map[string]int{}
    tokens = token_map()
    
    var tokens_name = map[string]string{}
    tokens_name = token_text_map()
    
    if v, ok := tokens[str]; ok {         //判断是否存在于config.go中
        return [4]string{strconv.Itoa(v),token_str_ifs(str),tokens_name[str],strconv.Itoa(line)}    //存在  获取类型
    }else{
        return [4]string{strconv.Itoa(0),str,"T_TEXT",strconv.Itoa(line)}    //不存在 类型直接定义为sting
    }
}

func Wsp_Lexical_func(code string)(map[int][4]string){
                                                    /*读取文件 并放在data变量  ---start*/
    code =string(code)+"\n"+"\n"
                                                    /*读取文件 并放在data变量   --- end*/
                                                    /*设置全局循环数值         ---start*/
    num := 0;           //该opcode指令集位置
    num_TEXTS :=0       //需要删除掉的字符串数组数量
    lock_TEXT := 0      //字符串自增锁
    line :=1            //行数叠加
                                                    /*设置全局循环数值           ---end*/
                                                    /*token分析开始            ---start*/
    var return_map = map[int][4]string{}    //定义token_map  类型为map
    code_wsp := strings.Split(code,"")      //将字符串转化为数组
    code_len := len(code_wsp)-1             //统计数组总量并减一
    
    for i:=0;i<=code_len;i++{               //开启循环
        tk := token_funcs(string(code_wsp[i][len(code_wsp[i])-1]),line);   //获取opcode主题最后一个字符串的token类型  只作为缓存作用
        if tk[0] == strconv.Itoa(0){     //判断类型
            if i!=len(code_wsp)-1 {      //判断是否不为最后一个字符
                tmp := code_wsp[i+1]     //把字符串数组中的下一个数组保存至tmp变量
                code_wsp[i+1]=code_wsp[i]+code_wsp[i+1]    //将两个变量拼接
                num_TEXTS++    //需要删除的字符串数组数量加一
                if token_funcs(string(code_wsp[i+1][len(code_wsp[i+1])-1]),line)[0] != strconv.Itoa(0){    //判断下一个字符是否不为字符串类型
                    code_wsp[i+1] = tmp
                    for z:=i-num_TEXTS+1;z<=i-1;z++{
                        code_wsp[z]=""          //根据num_TEXTS破坏字符串
                    }
                    num_TEXTS = 0               //需要破坏的数组归零
                    lock_TEXT = 1               //本次字符串循环结束，激活文本锁
                }
            }
        }else{
            code_wsp[i]=string(code_wsp[i][len(code_wsp[i])-1])   //清空上次opcode留存的字符串，只保留本次opcode关键字
        }
        if lock_TEXT==1 {     //判断文本锁状态
            return_map[num]=token_funcs(code_wsp[i],line)       //写入opcode
            num++                                         //opcode位置加一
            lock_TEXT = 0                                 //关闭文本锁
        }else if token_funcs(string(code_wsp[i][len(code_wsp[i])-1]),line)[0] != strconv.Itoa(0){    //否则判断最后一个字符的类型是否不为字符串
            return_map[num]=token_funcs(code_wsp[i],line)       //写入opcode
            num++                                         //opcode位置加一
        }
        if code_wsp[i]=="\n"{                             //统计行数
            line++
        }
    }
    return return_map
                                                    /*token分析结束              ---end*/
}

func token_var(str string,line int)([4]string){//判断token主题内容是否为字符串，如为字符串则返回T_TEXT token 否则匹配config.go设置的token进行判断
    var tokens = map[string]int{}
    tokens["="] = 7
    tokens["\n"] = 7
    
    var tokens_name = map[string]string{}
    tokens_name = token_text_map()
    
    if v, ok := tokens[str]; ok {         //判断是否存在于config.go中
        return [4]string{strconv.Itoa(v),token_str_ifs(str),tokens_name[str],strconv.Itoa(line)}    //存在  获取类型
    }else{
        return [4]string{strconv.Itoa(0),str,"T_TEXT",strconv.Itoa(line)}    //不存在 类型直接定义为sting
    }
}

func Wsp_Lexical_var(code string)(map[int][4]string){
                                                    /*读取文件 并放在data变量  ---start*/
    code =string(code)+"\n"
                                                    /*读取文件 并放在data变量   --- end*/
                                                    /*设置全局循环数值         ---start*/
    num := 0;           //该opcode指令集位置
    num_TEXTS :=0       //需要删除掉的字符串数组数量
    lock_TEXT := 0      //字符串自增锁
    line :=1            //行数叠加
                                                    /*设置全局循环数值           ---end*/
                                                    /*token分析开始            ---start*/
    var return_map = map[int][4]string{}    //定义token_map  类型为map
    code_wsp := strings.Split(code,"")      //将字符串转化为数组
    code_len := len(code_wsp)-1             //统计数组总量并减一
    
    for i:=0;i<=code_len;i++{               //开启循环
        tk := token_var(string(code_wsp[i][len(code_wsp[i])-1]),line);   //获取opcode主题最后一个字符串的token类型  只作为缓存作用
        if tk[0] == strconv.Itoa(0){     //判断类型
            if i!=len(code_wsp)-1 {      //判断是否不为最后一个字符
                tmp := code_wsp[i+1]     //把字符串数组中的下一个数组保存至tmp变量
                code_wsp[i+1]=code_wsp[i]+code_wsp[i+1]    //将两个变量拼接
                num_TEXTS++    //需要删除的字符串数组数量加一
                if token_var(string(code_wsp[i+1][len(code_wsp[i+1])-1]),line)[0] != strconv.Itoa(0){    //判断下一个字符是否不为字符串类型
                    code_wsp[i+1] = tmp
                    for z:=i-num_TEXTS+1;z<=i-1;z++{
                        code_wsp[z]=""          //根据num_TEXTS破坏字符串
                    }
                    num_TEXTS = 0               //需要破坏的数组归零
                    lock_TEXT = 1               //本次字符串循环结束，激活文本锁
                }
            }
        }else{
            code_wsp[i]=string(code_wsp[i][len(code_wsp[i])-1])   //清空上次opcode留存的字符串，只保留本次opcode关键字
        }
        if lock_TEXT==1 {     //判断文本锁状态
            return_map[num]=token_var(code_wsp[i],line)       //写入opcode
            num++                                         //opcode位置加一
            lock_TEXT = 0                                 //关闭文本锁
        }else if token_var(string(code_wsp[i][len(code_wsp[i])-1]),line)[0] != strconv.Itoa(0){    //否则判断最后一个字符的类型是否不为字符串
            return_map[num]=token_var(code_wsp[i],line)       //写入opcode
            num++                                         //opcode位置加一
        }
        if code_wsp[i]=="\n"{                             //统计行数
            line++
        }
    }
    return return_map
                                                    /*token分析结束              ---end*/
}