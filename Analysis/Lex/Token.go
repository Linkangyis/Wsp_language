package lex

import(
    "regexp"
)

type Lex_Struct struct{
    Type int 
    Text string
    Name string
    Line int
}
var Lines int=1

func Line_Echo()int{
    return Lines
}
func Line_Set(a int){
    Lines=a
}

func CodeNotes(text string)(string){
    match:=regexp.MustCompile(`//(.*)\n`).ReplaceAllString(text,"\n")
    match = regexp.MustCompile(`#(.*)\n`).ReplaceAllString(match,"\n")
    return match
}

func Wsp_Lexical(Code string)map[int]Lex_Struct{
    Code = CodeNotes(Code)
    Res:=make(map[int]Lex_Struct)
    var String_region string
    var ALLLock int = 0
    var SBrLock int = 0
    var MBrLock int = 0
    var XBrLock int = 0
    for i,Cpoints:=range Code{
        Cpoint := string(Cpoints)
        String_region+=Cpoint
        if Cpoint=="\n"{
            Lines++
        }
        if Cpoint=="\""{
            if ALLLock==0{
                ALLLock=1
            }else{
                ALLLock=0
                Res[len(Res)]=Lex_Struct{0,String_region,"String",Line_Echo()}
                String_region=""
            }
        }
        if Cpoint=="("&&XBrLock==0&&MBrLock==0&&ALLLock==0{
            if SBrLock<1{
                Type:=Token_Contrast_Map_Type(String_region)
                Name:=Token_Contrast_Map_Name(String_region)
                Res[len(Res)]=Lex_Struct{Type,String_region,Name,Line_Echo()}
                String_region=""
            }
            SBrLock++
        }
        if Cpoint==")"&&XBrLock==0&&MBrLock==0&&ALLLock==0{
            SBrLock--
            if SBrLock<1{
                String_region=String_region[0:len(String_region)-1]
                Res[len(Res)]=Lex_Struct{0,String_region,"String",Line_Echo()}
                String_region=")"
            }
            
        }
        
        
        if Cpoint=="["&&XBrLock==0&&SBrLock==0&&ALLLock==0{
            if MBrLock<1{
                Type:=Token_Contrast_Map_Type(String_region)
                Name:=Token_Contrast_Map_Name(String_region)
                Res[len(Res)]=Lex_Struct{Type,String_region,Name,Line_Echo()}
                String_region=""
            }
            MBrLock++
        }
        if Cpoint=="]"&&XBrLock==0&&SBrLock==0&&ALLLock==0{
            MBrLock--
            if MBrLock<1{
                String_region=String_region[0:len(String_region)-1]
                Res[len(Res)]=Lex_Struct{0,String_region,"String",Line_Echo()}
                String_region="]"
            }
            
        }
        
        
        if Cpoint=="{"&&SBrLock==0&&MBrLock==0&&ALLLock==0{
            if XBrLock<1{
                Type:=Token_Contrast_Map_Type(String_region)
                Name:=Token_Contrast_Map_Name(String_region)
                Res[len(Res)]=Lex_Struct{Type,String_region,Name,Line_Echo()}
                String_region=""
            }
            XBrLock++
        }
        if Cpoint=="}"&&MBrLock==0&&SBrLock==0&&ALLLock==0{
            XBrLock--
            if XBrLock<1{
                String_region=String_region[0:len(String_region)-1]
                Res[len(Res)]=Lex_Struct{0,String_region,"String",Line_Echo()}
                String_region="}"
            }
            
        }
        
        
        if Token_Contrast_Map_Name(String_region)!="NULL"&&SBrLock==0&&XBrLock==0&&MBrLock==0&&ALLLock==0{
            Type:=Token_Contrast_Map_Type(String_region)
            Name:=Token_Contrast_Map_Name(String_region)
            String_R:=Token_Replace_String(String_region)
            if Name!="SPACE"{
                Res[len(Res)]=Lex_Struct{Type,String_R,Name,Line_Echo()}
            }
            String_region=""
        }else if SBrLock==0&&XBrLock==0&&MBrLock==0&&ALLLock==0{
            if string(Code[i+1])=="(" || string(Code[i+1])==" " || string(Code[i+1])=="[" || string(Code[i+1])=="{"||string(Code[i+1])=="+" || string(Code[i+1])=="-" || string(Code[i+1])=="*" || string(Code[i+1])=="/" || string(Code[i+1])=="%" || string(Code[i+1])=="\n" || string(Code[i+1])=="=" || string(Code[i+1])=="," || string(Code[i+1])==";" || string(Code[i+1])=="<" || string(Code[i+1])==">" || string(Code[i+1])=="!" || string(Code[i+1])==":"  {
                Res[len(Res)]=Lex_Struct{0,String_region,"String",Line_Echo()}
                String_region=""
            }
        }
    }
    return Res
}
