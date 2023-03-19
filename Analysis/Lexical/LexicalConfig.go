package lex

func LexInit(){
    TokenInit()
    SeparateInit()
}
func TokenInit(){
    //2系列 关键词
    TokenConfigMap["function"] = TokenConfig{
        Name       : "FUNCTION",
        Type       : 200,
    }
    
    TokenConfigMap["class"] = TokenConfig{
        Name       : "CLASS",
        Type       : 201,
    }
    
    TokenConfigMap["return"] = TokenConfig{
        Name       : "RETURN",
        Type       : 202,
    }
    
    TokenConfigMap["for"] = TokenConfig{
        Name       : "FOR",
        Type       : 203,
    }
    
    //3系列 系统级函数
    TokenConfigMap["print"] = TokenConfig{
        Name       : "PRINT",
        Type       : 300,
    }
    
    //4系列 符号
    TokenConfigMap[" "] = TokenConfig{
        Name       : "SPACE",
        Type       : 400,
        Hide       : true,
    }
    
    TokenConfigMap["\n"] = TokenConfig{
        Name       : "WARP",
        Type       : 401,
        Replace    : "_WARP_",
        Hide       : true,
    }
    
    TokenConfigMap["("] = TokenConfig{
        Name       : "PARENTHESES_0",
        Type       : 402,
    }
    
    TokenConfigMap[")"] = TokenConfig{
        Name       : "PARENTHESES_1",
        Type       : 403,
    }
    
    TokenConfigMap["["] = TokenConfig{
        Name       : "BRACKETS_0",
        Type       : 404,
    }
    
    TokenConfigMap["]"] = TokenConfig{
        Name       : "BRACKETS_1",
        Type       : 405,
    }
    
    TokenConfigMap["{"] = TokenConfig{
        Name       : "BRACES_0",
        Type       : 406,
    }
    
    TokenConfigMap["}"] = TokenConfig{
        Name       : "BRACES_1",
        Type       : 407,
    }
    
    TokenConfigMap[";"] = TokenConfig{
        Name       : "END",
        Type       : 408,
    }
    
    TokenConfigMap["'"] = TokenConfig{
        Name       : "QUOTES_A",
        Type       : 409,
    }
    
    TokenConfigMap["\""] = TokenConfig{
        Name       : "QUOTES_B",
        Type       : 410,
    }
    
    TokenConfigMap["`"] = TokenConfig{
        Name       : "QUOTES_C",
        Type       : 411,
    }
    
    TokenConfigMap["$"] = TokenConfig{
        Name       : "VAR",
        Type       : 412,
    }
    
    TokenConfigMap[","] = TokenConfig{
        Name       : "FENGE",
        Type       : 413,
    }
    
    //5系列 运算符
    TokenConfigMap["+"] = TokenConfig{
        Name       : "ADD",
        Type       : 500,
    }
    
    TokenConfigMap["-"] = TokenConfig{
        Name       : "ABB",
        Type       : 501,
    }
    
    TokenConfigMap["*"] = TokenConfig{
        Name       : "CHENG",
        Type       : 502,
    }
    
    TokenConfigMap["/"] = TokenConfig{
        Name       : "CHU",
        Type       : 503,
    }
    
    TokenConfigMap[">"] = TokenConfig{
        Name       : "DAYU",
        Type       : 504,
    }
    
    TokenConfigMap["<"] = TokenConfig{
        Name       : "XIAOYU",
        Type       : 505,
    }
    
    TokenConfigMap["!"] = TokenConfig{
        Name       : "NO",
        Type       : 506,
    }
    
    TokenConfigMap["="] = TokenConfig{
        Name       : "EQUAL",
        Type       : 507,
    }
    
    TokenConfigMap["%"] = TokenConfig{
        Name       : "QUYU",
        Type       : 508,
    }
    
    TokenConfigMap["|"] = TokenConfig{
        Name       : "HUO",
        Type       : 509,
    }
    
    TokenConfigMap["&"] = TokenConfig{
        Name       : "YU",
        Type       : 510,
    }
}
func JudgeSeparate(Text string)bool{
    if Text=="<NIL>"{
        return false
    }else if _ , Type := SeparateConfigMap[Text];Type{
        return true
    }
    return false
}
func JudgeToken(Text string,Type bool)bool{
    if _ , Type := TokenConfigMap[Text];Type{
        return true
    }
    return false
}
func SeparateInit(){
    //基础
    SeparateConfigMap[" "] = true
    SeparateConfigMap["\n"] = true
    SeparateConfigMap["$"] = true
    
    //运算
    SeparateConfigMap["+"] = true
    SeparateConfigMap["-"] = true
    SeparateConfigMap["*"] = true
    SeparateConfigMap["/"] = true
    SeparateConfigMap["%"] = true
    SeparateConfigMap[">"] = true
    SeparateConfigMap["<"] = true
    SeparateConfigMap["="] = true
    SeparateConfigMap["!"] = true
    SeparateConfigMap["&"] = true
    SeparateConfigMap["|"] = true
    
    //字符串
    SeparateConfigMap["\""] = true
    SeparateConfigMap["'"] = true
    SeparateConfigMap["`"] = true
    
    //定界
    SeparateConfigMap["{"] = true
    SeparateConfigMap["}"] = true
    SeparateConfigMap["["] = true
    SeparateConfigMap["]"] = true
    SeparateConfigMap["("] = true
    SeparateConfigMap[")"] = true
    
    //END
    SeparateConfigMap[";"] = true
    
    //分割
    SeparateConfigMap[","] = true
}