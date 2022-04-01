package build

import(
  "../Types"
  "../Echo"
  "../Token"
)
type Builds_Struct struct {
    Codes  map[int][6]string
    Xovar  map[string]string     
    Funcs  map[string]map[int][6]string
    Funcs_list  map[string]string 
}
func Wsp_Build(Sem_Comp map[int][4]string)(Builds_Struct){
    Code_Build_NUM:=0
    Build := make(map[int][6]string)
    Xovars := make(map[string]string)
    flist := make(map[string]string)
    funxs :=make(map[string]map[int][6]string)
    Xovars_lens := 0
    for i:=0;i<=len(Sem_Comp)-1;i++{
        if Sem_Comp[i][0]==types.Strings(12){
            Build[Code_Build_NUM]=[6]string{Sem_Comp[i][0],Sem_Comp[i][1],Sem_Comp[i+2][1],"","",Sem_Comp[i][3]}
            Code_Build_NUM++
            i=i+2
        }else if Sem_Comp[i][0]==types.Strings(10){
            Build[Code_Build_NUM]=[6]string{Sem_Comp[i][0],Sem_Comp[i+1][1],Sem_Comp[i+3][1],"","0x"+echo.Zero_int(6-len(types.Strings(Xovars_lens)))+types.Strings(Xovars_lens),Sem_Comp[i][3]}
            Xovars["0x"+echo.Zero_int(6-len(types.Strings(Xovars_lens)))+types.Strings(Xovars_lens)]=Sem_Comp[i+6][1]
            funxs[Sem_Comp[i+1][1]]=Wsp_Build(token.Wsp_Semantic(token.Wsp_Grammar(token.Wsp_Lexical_func(Sem_Comp[i+6][1])))).Codes
            flist[Sem_Comp[i+1][1]]=Sem_Comp[i+3][1]
            Code_Build_NUM++
            Xovars_lens++
            i=i+3
        }else if Sem_Comp[i][0]==types.Strings(11){
            Build[Code_Build_NUM]=[6]string{Sem_Comp[i][0],Sem_Comp[i][1],Sem_Comp[i+2][1],"","0x"+echo.Zero_int(6-len(types.Strings(Xovars_lens)))+types.Strings(Xovars_lens),Sem_Comp[i][3]}
            funxs["0x"+echo.Zero_int(6-len(types.Strings(Xovars_lens)))+types.Strings(Xovars_lens)]=Wsp_Build(token.Wsp_Semantic(token.Wsp_Grammar(token.Wsp_Lexical_func(Sem_Comp[i+5][1])))).Codes
            Code_Build_NUM++
            Xovars_lens++
            i=i+3
        }else if Sem_Comp[i][0]==types.Strings(25){
            if Sem_Comp[i-1][0]==types.Strings(26){
                Build[Code_Build_NUM]=[6]string{types.Strings(28),Sem_Comp[i][1],Sem_Comp[i+2][1],"","0x"+echo.Zero_int(6-len(types.Strings(Xovars_lens)))+types.Strings(Xovars_lens),Sem_Comp[i][3]}
                funxs["0x"+echo.Zero_int(6-len(types.Strings(Xovars_lens)))+types.Strings(Xovars_lens)]=Wsp_Build(token.Wsp_Semantic(token.Wsp_Grammar(token.Wsp_Lexical_func(Sem_Comp[i+5][1])))).Codes
                Code_Build_NUM++
                Xovars_lens++
            }else{
                Build[Code_Build_NUM]=[6]string{Sem_Comp[i][0],Sem_Comp[i][1],Sem_Comp[i+2][1],"","0x"+echo.Zero_int(6-len(types.Strings(Xovars_lens)))+types.Strings(Xovars_lens),Sem_Comp[i][3]}
                funxs["0x"+echo.Zero_int(6-len(types.Strings(Xovars_lens)))+types.Strings(Xovars_lens)]=Wsp_Build(token.Wsp_Semantic(token.Wsp_Grammar(token.Wsp_Lexical_func(Sem_Comp[i+5][1])))).Codes
                Code_Build_NUM++
                Xovars_lens++
            }
        }else if Sem_Comp[i][0]==types.Strings(26){
            if Sem_Comp[i+1][0]!=types.Strings(25){
                Build[Code_Build_NUM]=[6]string{Sem_Comp[i][0],Sem_Comp[i][1],"","","0x"+echo.Zero_int(6-len(types.Strings(Xovars_lens)))+types.Strings(Xovars_lens),Sem_Comp[i][3]}
                funxs["0x"+echo.Zero_int(6-len(types.Strings(Xovars_lens)))+types.Strings(Xovars_lens)]=Wsp_Build(token.Wsp_Semantic(token.Wsp_Grammar(token.Wsp_Lexical_func(Sem_Comp[i+2][1])))).Codes
                Code_Build_NUM++
                Xovars_lens++
            }
        }else if Sem_Comp[i][0]==types.Strings(300){
            if Sem_Comp[i+2][0]==types.Strings(121) && Sem_Comp[i+5][0]!=types.Strings(7){
                Build[Code_Build_NUM]=[6]string{types.Strings(302),Sem_Comp[i+1][1],Sem_Comp[i+1][1]+"["+Sem_Comp[i+3][1]+"]","["+Sem_Comp[i+3][1]+"]","",Sem_Comp[i][3]}
                Code_Build_NUM++
            }else if Sem_Comp[i+2][0]==types.Strings(121) && Sem_Comp[i+5][0]==types.Strings(7){
                Build[Code_Build_NUM]=[6]string{types.Strings(304),Sem_Comp[i+1][1],Sem_Comp[i+1][1]+"["+Sem_Comp[i+3][1]+"]","["+Sem_Comp[i+3][1]+"]",types.Strings(Code_Build_NUM+1),Sem_Comp[i][3]}
                Code_Build_NUM++
            }else if Sem_Comp[i+2][0]==types.Strings(7){
                Build[Code_Build_NUM]=[6]string{types.Strings(301),Sem_Comp[i+1][1],Sem_Comp[i+1][1],"",types.Strings(Code_Build_NUM+1),Sem_Comp[i][3]}
                Code_Build_NUM++
            }else if Sem_Comp[i+2][0]==types.Strings(500){
                Build[Code_Build_NUM]=[6]string{types.Strings(303),Sem_Comp[i+1][1],"","",types.Strings(Code_Build_NUM+1),Sem_Comp[i][3]}
                Code_Build_NUM++
            }else{
                Build[Code_Build_NUM]=[6]string{Sem_Comp[i][0],Sem_Comp[i+1][1],"","","",Sem_Comp[i][3]}
                Code_Build_NUM++
            }
        }else if Sem_Comp[i][0]==types.Strings(200){
            Build[Code_Build_NUM]=[6]string{Sem_Comp[i][0],Sem_Comp[i][1],Sem_Comp[i+2][1],"","",Sem_Comp[i][3]}
            Code_Build_NUM++
        }else if Sem_Comp[i][0]==types.Strings(400){
            Build[Code_Build_NUM]=[6]string{Sem_Comp[i][0],Sem_Comp[i][1],Sem_Comp[i+2][1],"","",Sem_Comp[i][3]}
            Code_Build_NUM++
        }else if Sem_Comp[i][0]==types.Strings(0) && Sem_Comp[i-1][0]==types.Strings(7){
            Build[Code_Build_NUM]=[6]string{Sem_Comp[i][0],Sem_Comp[i][1],Sem_Comp[i][1],Sem_Comp[i][1],Sem_Comp[i][1],Sem_Comp[i][3]}
            Code_Build_NUM++
        }else if Sem_Comp[i][0]==types.Strings(401){
            Build[Code_Build_NUM]=[6]string{Sem_Comp[i][0],Sem_Comp[i][1],Sem_Comp[i+2][1],"","",Sem_Comp[i][3]}
            Code_Build_NUM++
        }else if Sem_Comp[i][0]==types.Strings(27){
            Build[Code_Build_NUM]=[6]string{Sem_Comp[i][0],Sem_Comp[i][1],"","",types.Strings(Code_Build_NUM+1),Sem_Comp[i][3]}
            Code_Build_NUM++
        }
    }
    Builds:=Builds_Struct{
            Codes      : Build,
            Xovar      : Xovars,
            Funcs      : funxs,
            Funcs_list : flist,
    }
    return Builds
}