package ast

import(
  "Wsp/Analysis/Lex"
  "Wsp/Module/Memory"
  "strconv"
  "strings"
)

func IsNum(s string) bool {
    _, err := strconv.ParseFloat(s, 64)
    return err == nil
}
func Wsp_Ast(Code map[int]lex.Lex_Struct)Ast_Tree{
    BodyAst:=Wsp_Ast_One(Code)
    FuncAst:=FuncAst_Struct{funcMap,funcVarMap}
    return Ast_Tree{FuncAst,BodyAst,classMap,classlocks}
}
func Stick_Brk(BodyAst_Struct BodyAst_Struct)string{
    Res:=BodyAst_Struct.Text
    for i:=0;i<=len(BodyAst_Struct.Abrk)-1;i++{
        switch BodyAst_Struct.Abrk[i].Type{
            case 0:
                Res+="("+BodyAst_Struct.Abrk[i].Text+")"
            case 1:
                Res+="["+BodyAst_Struct.Abrk[i].Text+"]"
            case 2:
                Res+="{"+BodyAst_Struct.Abrk[i].Text+"}"
        }
    }
    return Res
}
func So_Run(BodyAst BodyAst_Struct)string{
    if BodyAst.Type!=0{
        return BodyAst.Sbrk[0]
    }
    return BodyAst.Text
}
func Complex(code string)map[int]lex.Lex_Struct{
    return lex.Wsp_Lexical(code+"  ")
}
func TypeTrims(a string)string{
    return strings.Trim(a," ")
}
func Func_Stc(code string)map[int]string{
    codelex:=strings.Split(code, ",")
    res:=make(map[int]string)
    if len(codelex[0])==0{
        return res
    }
    for i:=0;i<=len(codelex)-1;i++{
        res[len(res)]=TypeTrims(codelex[i][1:])
    }
    return res
}
var funcMap = make(map[string]map[int]BodyAst_Struct)
var funcVarMap = make(map[string]map[int]string)

var classfuncMap = make(map[string]map[int]BodyAst_Struct)
var classfuncVarMap = make(map[string]map[int]string)
var classMap = make(map[string]ClassAstStruct)

func Wsp_Ast_One(lex map[int]lex.Lex_Struct)map[int]BodyAst_Struct{
    Res:=make(map[int]BodyAst_Struct)
    for i:=0;i<=len(lex)-1;i++{
        Name := lex[i].Name
        Type := lex[i].Type
        Line := lex[i].Line
        if lex[i].Type==17&&lex[i+2].Type==19{
            Type = 19
            var EndLien int
            var Addid int
            for z:=i;z<=len(lex)-1;z++{
                if lex[z].Type==75{
                    EndLien = z
                    break;
                }
                Addid++
            }
            var List string
            for z:=i+3;z<=EndLien-1;z++{
                if lex[z].Type==0{
                    List+=lex[z].Text+","
                }
            }
            lex[EndLien-1] = lex[i+1]
            List = List[0:len(List)-1]
            Name = List
            i+=Addid-1
        }
        if lex[i].Type==1{
            if funcnewlock>0 && classlock==0{
                Type = 52
            }
            if lex[i+1].Type==0{
                i++
            }else if classlock==0{
                Type = 51
            }
        }
        if lex[i].Type==50||lex[i].Type==17||lex[i].Type==18{
            i++
        }
        if Type==18{
            if lex[i].Text=="$"{
                Error("语法错误，NEW类不允许为变量")
            }
        }
        
        Text := lex[i].Text
        
        
        if lex[i].Type==21{
            Type = 21
            var EndLien int
            var Addid int
            for z:=i;z<=len(lex)-1;z++{
                if lex[z].Type==80{
                    EndLien = z
                    break;
                }
                Addid++
            }
            var List string
            for z:=i;z<=EndLien-1;z++{
                if lex[z].Type==0{
                    List+=lex[z].Text+","
                }
            }
            lex[EndLien-1] = lex[i+1]
            List = List[0:len(List)-1]
            Name = List
            i+=Addid-1
        }
        
        if lex[i].Type==4&&lex[i+1].Type==3{
            Type=5
            Name="ELIF"
            Text="elif"
            i++
        }
        if lex[i].Type==78{
            continue
        }
        StLock := 0
        if lex[i].Type==71{
            Name="EVAL"
            Type=15
            i--
            StLock=1
        }
        var slock_string = make(map[int]string)
        var mlock_string = make(map[int]string)
        var xlock_string = make(map[int]string)
        var alock_string = make(map[int]Brks)
        if lex[i].Type!=95 || StLock==1{
            for {
                switch lex[i+1].Type {
                    case 71:
                        if lex[i+2].Type==72{
                            slock_string[len(slock_string)]=""
                            alock_string[len(alock_string)]=Brks{}
                            i=i+2
                        }else{
                            slock_string[len(slock_string)]=lex[i+2].Text
                            alock_string[len(alock_string)]=Brks{0,lex[i+2].Text}
                            i=i+3
                        }
                    case 73:
                        if lex[i+2].Type==74{
                            alock_string[len(alock_string)]=Brks{}
                            mlock_string[len(mlock_string)]=""
                            i=i+2
                        }else{
                            mlock_string[len(mlock_string)]=lex[i+2].Text
                            alock_string[len(alock_string)]=Brks{1,lex[i+2].Text}
                            i=i+3
                        }
                    case 75:
                        if lex[i+2].Type==76{
                            alock_string[len(alock_string)]=Brks{}
                            xlock_string[len(xlock_string)]=""
                            i=i+2
                        }else{
                            xlock_string[len(xlock_string)]=lex[i+2].Text
                            alock_string[len(alock_string)]=Brks{2,"<SPACE>"}
                            i=i+3
                        }
                    case 91:
                        if lex[i+2].Type==98{
                            alock_string[len(alock_string)]=Brks{3,lex[i+3].Text}
                            i+=3
                        }
                }
                if lex[i+1].Type!=71&&lex[i+1].Type!=73&&lex[i+1].Type!=75{
                    break
                }
                if lex[i+1].Type!=71&&lex[i+1].Type!=73&&lex[i+1].Type!=75{
                    if lex[i+1].Type!=91 && lex[i+2].Type!=98{
                        break
                    }
                }
            }
        }
        Res[len(Res)]=BodyAst_Struct{Type,Name,Text,slock_string,mlock_string,xlock_string,alock_string,Line}
        if StLock==1{
            StLock=0
        }
    }
    for i:=0;i<=len(Res)-1;i++{
        Tmplen:=Line_Echo()
        Line_Set(Res[i].Line)
        if Res[i].Type==1 && classlock==0 && funcnewlock == 0{          //花括号内容编译 正常函数 域外
            funcnewlock++
            _,ok1:=funcMap[Res[i].Text];
            _,ok2:=funcMap["9C"+Res[i].Text]
            if ok1 || ok2{
                Error("Error: 函数 "+Res[i].Text+" 不允许重载")
            }
            funcMap[Res[i].Text]=Wsp_Ast_One(Complex(Res[i].Xbrk[0]))
            funcVarMap[Res[i].Text]=Func_Stc(Res[i].Sbrk[0])
            center.A_Memory_FromMap(Res[i].Text)
            Res[i].Xbrk[0]="True"
            funcnewlock--
        }else if Res[i].Type==52 && classlock==0{          //花括号内容编译 正常函数 域内
            _,ok1:=funcMap[Res[i].Text];
            _,ok2:=funcMap["9C"+Res[i].Text]
            if ok1 || ok2{
                Error("Error: 函数 "+Res[i].Text+" 不允许重载")
            }
            funcMap["9C"+Res[i].Text]=Wsp_Ast_One(Complex(Res[i].Xbrk[0]))
            funcVarMap["9C"+Res[i].Text]=Func_Stc(Res[i].Sbrk[0])
            center.A_Memory_FromMap("9C"+Res[i].Text)
            Res[i].Xbrk[0]="True"
        }else if Res[i].Type==1 && classlock==1{          //花括号内容编译 Class函数
            classfuncMap[Res[i].Text]=Wsp_Ast_One(Complex(Res[i].Xbrk[0]))
            classfuncVarMap[Res[i].Text]=Func_Stc(Res[i].Sbrk[0])
            Res[i].Xbrk[0]="True"
        }else if Res[i].Type==17{
            classlock=1
            classfuncMap = make(map[string]map[int]BodyAst_Struct)
            classfuncVarMap = make(map[string]map[int]string)
            
            Body := Wsp_Ast_One(Complex(Res[i].Xbrk[0]))
            FuncLists:=FuncAst_Struct{classfuncMap,classfuncVarMap}
            
            if _,ok:=classMap[Res[i].Text];ok{
                Error("Error: Class "+Res[i].Text+" 不允许重载")
            }
            classlocks[Res[i].Text]=false
            if funcnewlock==0{
                classlocks[Res[i].Text]=true
            }
            classMap[Res[i].Text] = ClassAstStruct{FuncLists,Body}
            classlock=0
        }else if Res[i].Type==19{
            classlock=1
            classfuncMap = make(map[string]map[int]BodyAst_Struct)
            classfuncVarMap = make(map[string]map[int]string)
            FatherList := strings.Split(Res[i].Name,",")
            
            for z:=0;z<=len(FatherList)-1;z++{
                Values:=classMap[FatherList[z]]
                classfuncMap = ExtendMapStick(classfuncMap,Values.ClassFunc.FuncList)
                classfuncVarMap = ExtendMapStickVar(classfuncVarMap,Values.ClassFunc.FuncVars)
            }
            
            Body := make(map[int]BodyAst_Struct)
            for z:=0;z<=len(FatherList)-1;z++{
                Values:=classMap[FatherList[z]].ClassBody
                Body = ExtendMapStickBody(Body,Values)
            }
            
            Body = ExtendMapStickBody(Body,Wsp_Ast_One(Complex(Res[i].Xbrk[0])))
            FuncLists:=FuncAst_Struct{classfuncMap,classfuncVarMap}
            if _,ok:=classMap[Res[i].Text];ok{
                Error("Error: Class "+Res[i].Text+" 不允许重载")
            }
            classlocks[Res[i].Text]=false
            if funcnewlock==0{
                classlocks[Res[i].Text]=true
            }
            classMap[Res[i].Text] = ClassAstStruct{FuncLists,Body}
            classlock=0
        }else if Res[i].Type==51{
            for z:=0;z<=len(Res[i].Xbrk)-1;z++{
                funcnewlock++
                Mer:=center.New_MemoryFunc()
                funcMap[Mer]=Wsp_Ast_One(Complex(Res[i].Xbrk[z]))
                Res[i].Xbrk[z]=Mer
                funcVarMap[Mer]=Func_Stc(Res[i].Sbrk[0])
                funcnewlock--
            }
        }else if len(Res[i].Xbrk)!=0{
            for z:=0;z<=len(Res[i].Xbrk)-1;z++{
                funcnewlock++
                Mer:=center.New_Memory()
                funcMap[Mer]=Wsp_Ast_One(Complex(Res[i].Xbrk[z]))
                Res[i].Xbrk[z]=Mer
                funcnewlock--
            }
        }
        Line_Set(Tmplen)
    }
    return Res
}
func Line_Echo()int{
    return lex.Line_Echo()
}
func Line_Set(a int){
    lex.Line_Set(a)
}
func ExtendMapStick(x map[string]map[int]BodyAst_Struct,y map[string]map[int]BodyAst_Struct)map[string]map[int]BodyAst_Struct{
    n:=make(map[string]map[int]BodyAst_Struct)
    for i,v := range x {
        n[i]=v
    }
    for i,v := range y {
        n[i]=v
    }
    return n
}
func ExtendMapStickVar(x map[string]map[int]string,y map[string]map[int]string)map[string]map[int]string{
    n:=make(map[string]map[int]string)
    for i,v := range x {
        n[i]=v
    }
    for i,v := range y {
        n[i]=v
    }
    return n
}
func ExtendMapStickBody(x map[int]BodyAst_Struct,y map[int]BodyAst_Struct)map[int]BodyAst_Struct{
    n:=make(map[int]BodyAst_Struct)
    for i:=0;i<=len(x)-1;i++{
        n[len(n)]=x[i]
    }
    for i:=0;i<=len(y)-1;i++{
        n[len(n)]=y[i]
    }
    return n
}