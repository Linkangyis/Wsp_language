package lex

func Token_Contrast_Map_Type(Text string)int{
    Maps:=make(map[string]int)
    Maps["function"]=1
    Maps["for"]=2
    Maps["if"]=3
    Maps["else"]=4
    Maps["return"]=11
    Maps["break"]=6
    Maps["continue"]=10
    Maps["print"]=7
    Maps["wgo"]=8
    Maps["add"]=9
    Maps["switch"]=12
    Maps["case"]=13
    Maps["default"]=14
    Maps["stick"]=16
    Maps["class"]=17
    Maps["new"]=18
    Maps["extends"]=19
    Maps["len"]=20
    Maps["global"]=21
    Maps["CurEnv"]=22
    Maps["free"]=23
    Maps["$"]=50
    Maps["("]=71
    Maps[")"]=72
    Maps["["]=73
    Maps["]"]=74
    Maps["{"]=75
    Maps["}"]=76
    Maps["\""]=77
    Maps["\n"]=78
    Maps[" "]=79
    Maps["\t"]=79
    Maps[";"]=80
    Maps["+"]=90
    Maps["-"]=91
    Maps["*"]=92
    Maps["/"]=93
    Maps["%"]=94
    Maps["="]=95
    Maps[","]=96
    Maps["<"]=97
    Maps[">"]=98
    Maps["!"]=99
    Maps[":"]=100
    if _,ok:=Maps[Text];ok{
        return Maps[Text]
    }
    //Error("未定义关键词 ["+Text+"] 强制退出")
    return -1
}
func Token_Contrast_Map_Name(Text string)string{
    Maps:=make(map[string]string)
    Maps["print"]="PRINT"
    Maps["wgo"]="WGO"
    Maps["function"]="FUNCTION"
    Maps["for"]="FOR"
    Maps["if"]="IF"
    Maps["else"]="ELSE"
    Maps["return"]="RETURN"
    Maps["break"]="BREAK"
    Maps["continue"]="CONTINUE"
    Maps["add"]="ADD"
    Maps["switch"]="SWITCH"
    Maps["case"]="SWITCH_CASE"
    Maps["default"]="SWITCH_DEFAULT"
    Maps["stick"]="TEXTSTICK"
    Maps["class"]="CLASS"
    Maps["new"]="NEWCLASS"
    Maps["extends"]="CLASSEXTENDS"
    Maps["len"]="ALLLEN"
    Maps["global"]="ALLGLOBAL"
    Maps["CurEnv"]="SETENV"
    Maps["free"]="FREE"
    Maps["$"]="VAR"
    Maps["("]="S_BRACKETS_A"
    Maps[")"]="S_BRACKETS_B"
    Maps["["]="M_BRACKETS_A"
    Maps["]"]="M_BRACKETS_B"
    Maps["{"]="X_BRACKETS_A"
    Maps["}"]="X_BRACKETS_B"
    Maps["\""]="STRING_QUOTE"
    Maps["\n"]="LINE_ADD"
    Maps[" "]="SPACE"
    Maps["\t"]="SPACETAB"
    Maps[";"]="END"
    Maps["+"]="CRUN_ADD"
    Maps["-"]="CRUN_SUB"
    Maps["*"]="CRUN_MUL"
    Maps["/"]="CRUN_DIV"
    Maps["%"]="CRUN_RES"
    Maps["="]="EQUAL"
    Maps[","]="STC"
    Maps["<"]="GT"
    Maps[">"]="LT"
    Maps["!"]="NT"
    Maps[":"]="START"
    if _,ok:=Maps[Text];ok{
        return Maps[Text]
    }
    //Error("未定义关键词 ["+Text+"] 强制退出")
    return "NULL"
}

func TabReturn()map[string]int{
    Maps:=make(map[string]int)
    Maps["function"]=1
    Maps["for"]=2
    Maps["if"]=3
    Maps["else"]=4
    Maps["return"]=11
    Maps["break"]=6
    Maps["continue"]=10
    Maps["print"]=7
    Maps["wgo"]=8
    Maps["add"]=9
    Maps["switch"]=12
    Maps["case"]=13
    Maps["default"]=14
    Maps["stick"]=16
    Maps["class"]=17
    Maps["new"]=18
    Maps["extends"]=19
    Maps["len"]=20
    Maps["global"]=21
    Maps["CurEnv"]=22
    Maps["free"]=23
    return Maps
}

func Token_Replace_String(Text string)string{
    Maps:=make(map[string]string)
    Maps["\n"]=""
    Maps["\t"]=""
    if _,ok:=Maps[Text];ok{
        return Maps[Text]
    }
    return Text
}

func LegitimateText(OneB string)bool{
    Maps := make(map[string]bool)
    Maps["_"] = true
    Maps["."] = true
    if StrAllLetter(OneB){
        return true
    }else if IsNum(OneB){
        return true
    }else if _,ok:=Maps[OneB];ok{
        return true
    }
    return false
}