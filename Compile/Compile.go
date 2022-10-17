package compile

import(
  "Wsp/Analysis/Ast"
  "Wsp/Module/Memory"
)

func Wsp_Compile(Codes ast.Ast_Tree)Res_Struct{
    MemoryList:=center.R_Memory_FromMap()
    Code = Codes
    Res:=Res_Struct{}
    Res.Body=Wsp_Compile_l(Codes.BodyAst)
    
    Funcs:=make(map[string]map[int]map[int]Body_Struct_Run)
    for i:=0;i<=len(MemoryList)-1;i++{
        Name:=MemoryList[i]
        Funcs[Name]=Wsp_Compile_l(Codes.FuncAst.FuncList[Name])
    }
    Class := make(map[string]ClassStruct)
    ClassList:=Codes.ClassAst
    for ClassName, Value := range ClassList{
        Func_Structs:=Value.ClassFunc
        FuncClass:=make(map[string]map[int]map[int]Body_Struct_Run)
        for ClassFuncName,TValue := range Func_Structs.FuncList{
            FuncClass[ClassFuncName]=Wsp_Compile_l(TValue)
        }
        ClassFunc:=Func_Struct{FuncClass,Func_Structs.FuncVars}
        ClassBody:=Wsp_Compile_l(Value.ClassBody)
        ClassOver:=ClassStruct{ClassFunc,ClassBody}
        Class[ClassName]=ClassOver
    }
    
    Res.Func=Func_Struct{Funcs,Codes.FuncAst.FuncVars}
    Res.Class=Class
    Res.ClassLock=Codes.ClassLock
    return Res
}

func BrStick(Code ast.BodyAst_Struct)string{
    Res := ""
    List := Code.Abrk
    for i:=0;i<=len(List)-1;i++{
        Type := List[i].Type
        Text := List[i].Text
        if Type == 0{
            Res += "("+Text+")"
        }else if Type == 1{
            Res += "["+Text+"]"
        }else if Type == 3{
            Res += "->"+Text
        }
    }
    return Res
}

func Wsp_Compile_l(TCode map[int]ast.BodyAst_Struct)map[int]map[int]Body_Struct_Run{
    Res:=make(map[int]map[int]Body_Struct_Run)
    Len_Line := 0
    Res[Len_Line]=make(map[int]Body_Struct_Run)
    for i:=0;i<=len(TCode)-1;i++{
        //fmt.Println(TCode[i].Text,TCode[i].Type,i)
        switch TCode[i].Type{
            case 50:
                if TCode[i+1].Type==95{
                    tmp:=0
                    for z:=i;z<=len(TCode)-1;z++{
                        if TCode[z].Type!=50&&TCode[z].Type!=95{
                            break
                        }else if TCode[z].Type==50&&TCode[z+1].Type==95{
                            Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                                Type : 301,
                                Abrk : TCode[z].Abrk,
                                Name : "SET_VAR",
                                Text : TCode[z].Text,
                                Movs : "<NIL>",
                                Line : TCode[z].Line,
                            }
                        }else{
                            break
                        }
                        tmp++
                    }
                    //fmt.Println(i)
                    i+=tmp-1
                }else if TCode[i].Abrk[0].Type!=2 && (TCode[i+1].Type!=90 && TCode[i+2].Type!=90) && (TCode[i+1].Type!=91 && TCode[i+2].Type!=91) && (TCode[i+1].Type!=92 && TCode[i+1].Type!=93 && TCode[i+1].Type!=94){
                    Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                        Type : 302,
                        Abrk : TCode[i].Abrk,
                        Name : "SO_VAR",
                        Text : TCode[i].Text,
                        Movs : "<NIL>",
                        Line : TCode[i].Line,
                    }
                }else if TCode[i+1].Type==90 && TCode[i+2].Type==90{
                    Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                        Type : 303,
                        Abrk : TCode[i].Abrk,
                        Name : "ADD_VAR",
                        Text : TCode[i].Text,
                        Movs : "<NIL>",
                        Line : TCode[i].Line,
                    }
                }else if TCode[i+1].Type==91 && TCode[i+2].Type==91{
                    Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                        Type : 303,
                        Abrk : TCode[i].Abrk,
                        Name : "ABB_VAR",
                        Text : TCode[i].Text,
                        Movs : "<NIL>",
                        Line : TCode[i].Line,
                    }
                }else if TCode[i].Abrk[0].Type!=2{
                    Mov := ""
                    if TCode[i-1].Type==90||TCode[i-1].Type==91{
                        Mov+=TCode[i-1].Text
                    }
                    tmp:=0
                    for z:=i;z<=len(TCode)-1;z++{
                        if TCode[z].Type==80 || TCode[z].Type==96{
                            break
                        }
                        Text := ""
                        if TCode[z].Type==50{
                            Text = "$"+TCode[z].Text
                        }else{
                            Text = TCode[z].Text
                        }
                        Mov += Text+BrStick(TCode[z])
                        tmp++
                    }
                    Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                        Type : 304,
                        Name : "CRUN_VAR",
                        Text : Mov,
                        Line : TCode[i].Line,
                    }
                    i+=tmp-1
                }
                
            case 7:
                Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                    Type : 200,
                    Abrk : TCode[i].Abrk,
                    Name : "PRINT",
                    Text : TCode[i].Sbrk[0],
                    Movs : "<NIL>",
                    Line : TCode[i].Line,
                }
            case 80:
                Len_Line++
                Res[Len_Line]=make(map[int]Body_Struct_Run)
            case 2:
                Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                    Type : 201,
                    Abrk : TCode[i].Abrk,
                    Name : "FOR",
                    Text : TCode[i].Xbrk[0],
                    Movs : "<NIL>",
                    Line : TCode[i].Line,
                }
                Len_Line++
                Res[Len_Line]=make(map[int]Body_Struct_Run)
            case 3:
                for z:=i;z<=len(TCode)-1;z++{
                    if TCode[z].Type==3 || TCode[z].Type==4 || TCode[z].Type==5{
                        Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                            Type : 199+TCode[z].Type,
                            Abrk : TCode[z].Abrk,
                            Name : TCode[z].Name,
                            Text : TCode[z].Xbrk[0],
                            Movs : "<NIL>",
                            Line : TCode[z].Line,
                        }
                    }else{
                        break
                    }
                }
                Len_Line++
                Res[Len_Line]=make(map[int]Body_Struct_Run)
            case 0:
                if len(TCode[i].Sbrk)>0 && TCode[i+1].Type!=90&&TCode[i+1].Type!=91&&TCode[i+1].Type!=92&&TCode[i+1].Type!=93&&TCode[i+1].Type!=94{
                    Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                        Type : 205,
                        Abrk : TCode[i].Abrk,
                        Name : TCode[i].Text,
                        Text : TCode[i].Sbrk[0],
                        Movs : "<NIL>",
                        Line : TCode[i].Line,
                    }
                }else if TCode[i+1].Type==90||TCode[i+1].Type==91||TCode[i+1].Type==92||TCode[i+1].Type==93||TCode[i+1].Type==94 || (TCode[i+1].Type==0 && (TCode[i+2].Type==90||TCode[i+2].Type==91||TCode[i+2].Type==92||TCode[i+2].Type==93||TCode[i+2].Type==94)){
                    Mov := ""
                    if TCode[i-1].Type==90||TCode[i-1].Type==91{
                        Mov+=TCode[i-1].Text
                    }
                    tmp:=0
                    for z:=i;z<=len(TCode)-1;z++{
                        if TCode[z].Type==80 || TCode[z].Type==96{
                            break
                        }
                        Text := ""
                        if TCode[z].Type==50{
                            Text = "$"+TCode[z].Text
                        }else{
                            Text = TCode[z].Text
                        }
                        Mov += Text+BrStick(TCode[z])
                        tmp++
                    }
                    Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                        Type : 304,
                        Name : "CRUN_VAR",
                        Text : Mov,
                        Line : TCode[i].Line,
                    }
                    i+=tmp-1
                }else if TCode[i].Text!=""{
                    _,ok:=Code.FuncAst.FuncVars[TCode[i].Text]
                    if string(TCode[i].Text[0])=="\""||ast.IsNum(TCode[i].Text)||ok{
                        Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                            Type : 0,
                            Name : TCode[i].Name,
                            Text : TCode[i].Text,
                            Movs : "<NIL>",
                            Line : TCode[i].Line,
                        }
                    }else if _,ok:=Code.FuncAst.FuncList[TCode[i].Text];ok{
                        Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                            Type : 206,
                            Name : TCode[i].Name,
                            Text : TCode[i].Text,
                            Movs : "<NIL>",
                            Line : TCode[i].Line,
                        }
                    }else{
                        Errors("函数 ["+TCode[i].Text+"] 不存在",TCode[i].Line)
                    }
                }
            case 11:
                Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                    Type : 207,
                    Abrk : TCode[i].Abrk,
                    Name : "RETURN",
                    Text : "",
                    Movs : "<NIL>",
                    Line : TCode[i].Line,
                }
            case 8:
                Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                    Type : 208,
                    Abrk : TCode[i].Abrk,
                    Name : "WGO",
                    Text : "",
                    Movs : "<NIL>",
                    Line : TCode[i].Line,
                }
            case 9:
                Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                    Type : 209,
                    Abrk : TCode[i].Abrk,
                    Name : "ADD",
                    Text : TCode[i].Sbrk[0],
                    Movs : "<NIL>",
                    Line : TCode[i].Line,
                }
            case 6:
                Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                    Type : 210,
                    Name : "BREAK",
                    Text : "",
                    Movs : "<NIL>",
                    Line : TCode[i].Line,
                }
            case 10:
                Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                    Type : 211,
                    Name : "CONTINUE",
                    Text : "",
                    Movs : "<NIL>",
                    Line : TCode[i].Line,
                }
            case 100:
                Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                    Type : 212,
                    Abrk : TCode[i].Abrk,
                    Name : "START",
                    Text : TCode[i].Xbrk[0],
                    Movs : "<NIL>",
                    Line : TCode[i].Line,
                }
                Len_Line++
                Res[Len_Line]=make(map[int]Body_Struct_Run)
            case 12:
                Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                    Type : 213,
                    Abrk : TCode[i].Abrk,
                    Name : TCode[i].Sbrk[0],
                    Text : TCode[i].Xbrk[0],
                    Movs : "<NIL>",
                    Line : TCode[i].Line,
                }
                Len_Line++
                Res[Len_Line]=make(map[int]Body_Struct_Run)
            case 13:
                Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                    Type : 214,
                    Abrk : TCode[i].Abrk,
                    Name : "SWITCH_CASE",
                    Text : TCode[i].Xbrk[0],
                    Movs : "",
                    Line : TCode[i].Line,
                }
            case 14:
                Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                    Type : 215,
                    Abrk : TCode[i].Abrk,
                    Name : "SWITCH_DEFAULT",
                    Text : TCode[i].Xbrk[0],
                    Movs : "",
                    Line : TCode[i].Line,
                }
            case 95:
                if TCode[i+1].Type==90||TCode[i+1].Type==91||TCode[i+1].Type==92||TCode[i+1].Type==93||TCode[i+1].Type==94{
                    Mov := ""
                    tmp:=0
                    for z:=i;z<=len(TCode)-1;z++{
                        if TCode[z].Type==80 || TCode[z].Type==96{
                            break
                        }
                        Text := ""
                        if TCode[z].Type==50{
                            Text = "$"+TCode[z].Text
                        }else{
                            Text = TCode[z].Text
                        }
                        Mov += Text+BrStick(TCode[z])
                        tmp++
                    }
                    Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                        Type : 304,
                        Name : "CRUN_VAR",
                        Text : Mov[1:],
                        Line : TCode[i].Line,
                    }
                    i+=tmp-1
                }
            case 15:
                if TCode[i+1].Type==90||TCode[i+1].Type==91||TCode[i+1].Type==92||TCode[i+1].Type==93||TCode[i+1].Type==94{
                    Mov := ""
                    if TCode[i-1].Type==90||TCode[i-1].Type==91{
                        Mov+=TCode[i-1].Text
                    }
                    tmp:=0
                    for z:=i;z<=len(TCode)-1;z++{
                        if TCode[z].Type==80 || TCode[z].Type==96{
                            break
                        }
                        Text := ""
                        if TCode[z].Type==50{
                            Text = "$"+TCode[z].Text
                        }else{
                            Text = TCode[z].Text
                        }
                        Mov += Text+BrStick(TCode[z])
                        tmp++
                    }
                    Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                        Type : 304,
                        Name : "CRUN_VAR",
                        Text : Mov[1:],
                        Line : TCode[i].Line,
                    }
                    i+=tmp-1
                }else{
                    Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                        Type : 216,
                        Abrk : TCode[i].Abrk,
                        Name : "EVAL",
                        Text : TCode[i].Sbrk[0],
                        Movs : "<NIL>",
                        Line : TCode[i].Line,
                    }
                }
            case 16:
                Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                    Type : 217,
                    Abrk : TCode[i].Abrk,
                    Name : "STICK",
                    Text : TCode[i].Sbrk[0],
                    Movs : "<NIL>",
                    Line : TCode[i].Line,
                }
            case 18:
                Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                    Type : 218,
                    Abrk : TCode[i].Abrk,
                    Name : TCode[i].Text,
                    Text : TCode[i].Sbrk[0],
                    Movs : "<NIL>",
                    Line : TCode[i].Line,
                }
            case 20:
                if TCode[i+1].Type==90||TCode[i+1].Type==91||TCode[i+1].Type==92||TCode[i+1].Type==93||TCode[i+1].Type==94{
                    Mov := ""
                    if TCode[i-1].Type==90||TCode[i-1].Type==91{
                        Mov+=TCode[i-1].Text
                    }
                    tmp:=0
                    for z:=i;z<=len(TCode)-1;z++{
                        if TCode[z].Type==80 || TCode[z].Type==96{
                            break
                        }
                        Text := ""
                        if TCode[z].Type==50{
                            Text = "$"+TCode[z].Text
                        }else{
                            Text = TCode[z].Text
                        }
                        Mov += Text+BrStick(TCode[z])
                        tmp++
                    }
                    Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                        Type : 304,
                        Name : "CRUN_VAR",
                        Text : Mov,
                        Line : TCode[i].Line,
                    }
                    i+=tmp-1
                }else{
                    Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                        Type : 219,
                        Abrk : TCode[i].Abrk,
                        Name : "LEN",
                        Text : TCode[i].Sbrk[0],
                        Movs : "<NIL>",
                        Line : TCode[i].Line,
                    }
                }
            case 51:
                Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                    Type : 220,
                    Abrk : TCode[i].Abrk,
                    Name : "FUNCNEW",
                    Text : TCode[i].Xbrk[0],
                    Movs : "<NIL>",
                    Line : TCode[i].Line,
                }
                Len_Line++
                Res[Len_Line]=make(map[int]Body_Struct_Run)
            case 52:
                Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                    Type : 221,
                    Abrk : TCode[i].Abrk,
                    Name : TCode[i].Text,
                    Text : TCode[i].Xbrk[0],
                    Movs : "<NIL>",
                    Line : TCode[i].Line,
                }
                Len_Line++
                Res[Len_Line]=make(map[int]Body_Struct_Run)
            case 21:
                Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                    Type : 222,
                    Abrk : TCode[i].Abrk,
                    Name : TCode[i].Text,
                    Text : TCode[i].Name,
                    Movs : "<NIL>",
                    Line : TCode[i].Line,
                }
                Len_Line++
                Res[Len_Line]=make(map[int]Body_Struct_Run)
            case 17:
                Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                    Type : 223,
                    Abrk : make(map[int]ast.Brks),
                    Name : "CLASSLOCK",
                    Text : TCode[i].Text,
                    Movs : "<NIL>",
                    Line : TCode[i].Line,
                }
                Len_Line++
                Res[Len_Line]=make(map[int]Body_Struct_Run)
            case 19:
                Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                    Type : 223,
                    Abrk : make(map[int]ast.Brks),
                    Name : "CLASSLOCK",
                    Text : TCode[i].Text,
                    Movs : "<NIL>",
                    Line : TCode[i].Line,
                }
                Len_Line++
                Res[Len_Line]=make(map[int]Body_Struct_Run)
        }
    }
    Check(Res)
    return Res
}