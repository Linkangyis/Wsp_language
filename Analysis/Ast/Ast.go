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
    return Ast_Tree{FuncAst,BodyAst}
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
func Wsp_Ast_One(lex map[int]lex.Lex_Struct)map[int]BodyAst_Struct{
    Res:=make(map[int]BodyAst_Struct)
    for i:=0;i<=len(lex)-1;i++{
        Name := lex[i].Name
        Type := lex[i].Type
        Line := lex[i].Line
        if lex[i].Type==50||lex[i].Type==1{
            i++
        }
        Text := lex[i].Text
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
                }
                if lex[i+1].Type!=71&&lex[i+1].Type!=73&&lex[i+1].Type!=75{
                    break
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
        if Res[i].Type==1{
            funcMap[Res[i].Text]=Wsp_Ast_One(Complex(Res[i].Xbrk[0]))
            funcVarMap[Res[i].Text]=Func_Stc(Res[i].Sbrk[0])
            center.A_Memory_FromMap(Res[i].Text)
            Res[i].Xbrk[0]="True"
        }else if len(Res[i].Xbrk)!=0{
            for z:=0;z<=len(Res[i].Xbrk)-1;z++{
                Mer:=center.New_Memory()
                funcMap[Mer]=Wsp_Ast_One(Complex(Res[i].Xbrk[z]))
                Res[i].Xbrk[z]=Mer
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