package build

import(
    "fmt"
    "strings"
    "../Token"
    "../Types"
)
var imports = make(map[int]string)
var num int = 0
var ALLS_ABD Builds_Struct
var Body_func string
var funcre int = 0
var Bodys string
var Body_funcs string

func import_ADD(a string)(string){
    lock :=1
    for i:=0;i<=len(imports)-1;i++{
        if imports[i]==a{
            lock=0
        }
    }
    if lock==1{
        imports[num]=a
        num++
    }
    return a
}
func var_so(a string)string{
    if string(a[0])=="$"{
        return string("Vars[\""+a[1:]+"\"]")
    }else if string(a[0])=="\""{
        return a
    }else if types.IsNum(a){
        return "\""+a+"\""
    }else{
        Lex:=token.Wsp_Lexical_func(a)
        Sem:=token.Wsp_Semantic(Lex)
        Gra:=token.Wsp_Grammar(Sem)
        Buildse:=Wsp_Build(Gra)
        return strings.Replace(Wsp_GTo_Go(Buildse.Codes,ALLS_ABD.Funcs), "\n", "", -1)
    }
    return a
}
func var_soa(a string)string{
    if string(a[0])=="$"{
        return string("Vars[\""+a[1:]+"\"].(string)")
    }else if string(a[0])=="\""{
        return "`"+a+"`"
    }else if types.IsNum(a){
        return "\""+a+"\""
    }else{
        Lex:=token.Wsp_Lexical_func(a)
        Sem:=token.Wsp_Semantic(Lex)
        Gra:=token.Wsp_Grammar(Sem)
        Buildse:=Wsp_Build(Gra)
        return strings.Replace(Wsp_GTo_Go(Buildse.Codes,ALLS_ABD.Funcs), "\n", "", -1)
    }
    return a
}
func var_sos(a string)string{
    if string(a[0])=="$"{
        return string(a[1:])
    }
    return a
}
var sosdl int =0
func Wsp_GTo_Go(opcode map[int][6]string,Funsx map[string]map[int][6]string)(string){
    Body:=""
    for i:=0;i<=len(opcode)-1;i++{
        if opcode[i][0]=="12"{
            Body+="\nfmt.Println("+var_so(opcode[i][2])+")"
            Body+="\n"
            import_ADD("fmt")
        }else if opcode[i][0]=="301"{
            Body+="Vars[\""+opcode[i][1]+"\"]="
        }else if opcode[i][0]=="300"{
            Body+="Vars[\""+opcode[i][1]+"\"].(string)\n"
        }else if opcode[i][0]=="0"{
            Body+="\""+opcode[i][2]+"\""
            Body+="\n"
        }else if opcode[i][0]=="200"{
            if _,ok:=Funsx[opcode[i][1]];ok{
                str := var_soa(opcode[i][2])
                Body+=opcode[i][1]+"("+str+")"
                Body+="\n"
            }else{
                import_ADD("os")
                import_ADD("plugin")
                import_ADD("io/ioutil")
                import_ADD("strings")
                if sosdl!=3{
                    sosdl=1
                }
                str := var_soa(opcode[i][2])
                Body+="So_func_map[\""+opcode[i][1]+"\"].(func(string) string)("+str+")\n"
            }
        }else if opcode[i][0]=="11"{
            a := opcode[i][2]
            lone := make(map[int]map[int][4]string)
            vars_chuan := strings.Split(a, ";")
            for i:=0;i<=len(vars_chuan)-1;i++{
                lone[i]=token.Wsp_Semantic(token.Wsp_Grammar(token.Wsp_Lexical_func(vars_chuan[i])))
            }
            lioen := lone[0][3][1]
            tj := lone[1][2][1]
            addfornum:=lone[2][0][1]
            addfornums:=lone[1][4][1]
            addorabb:=lone[3][3][1]
            bls:=lone[3][1][1]
            Body+="for "+bls+":="+lioen+";"+bls+tj+"="+addfornums+";"+bls+"="+bls+addorabb+""+addfornum+"{\nVars[\""+bls+"\"]="+bls+"\n"
            Body+=Wsp_GTo_Go(Funsx[opcode[i][4]],Funsx)
            Body+="}\n"
        }else if opcode[i][0]=="10"{
            Vars_List:=""
            Vars_Lists:=""
            str_arr := strings.Split(opcode[i][2], ",")
            for z:=0;z<=len(str_arr)-1;z++{
                Vars_List+=var_so(strings.Trim(str_arr[z]," "))+"="+var_sos(strings.Trim(str_arr[z]," "))+"\n"
            }
            for z:=0;z<=len(str_arr)-2;z++{
                Vars_Lists+=var_sos(strings.Trim(str_arr[z]," "))+" string ,"
            }
            Vars_Lists+=var_sos(strings.Trim(str_arr[len(str_arr)-1]," "))+" string"
            Body_func+="func "+opcode[i][1]+"("+Vars_Lists+")(string){\nvar Vars = make(map[string]interface{})\n"
            Body_func+= Vars_List
            Body_func+=Wsp_GTo_Go(Funsx[opcode[i][1]],Funsx)
            if funcre==0{
                Body_func+= "return \"\"\n"
            }else{
                funcre=0
            }
            Body_func+="}\n"
        }else if opcode[i][0]=="27"{
            Body+=opcode[i][1]
            funcre = 1
        }else if opcode[i][0]=="25"{
            lone:=token.Wsp_Semantic(token.Wsp_Grammar(token.Wsp_Lexical_var(opcode[i][2])))
            addorabb:=lone[2][1]
            A:=lone[0][1]
            B:=lone[3][1]
            Body+="if "+var_soa(A)+addorabb+"="+var_soa(B)+" {\n"
            Body+=Wsp_GTo_Go(Funsx[opcode[i][4]],Funsx)
            Body+="}"
        }else if opcode[i][0]=="28"{
            Body+="else if "+opcode[i][2]+" {\n"
            Body+=Wsp_GTo_Go(Funsx[opcode[i][4]],Funsx)
            Body+="}"
        }else if opcode[i][0]=="26"{
            Body+="else{\n"
            Body+=Wsp_GTo_Go(Funsx[opcode[i][4]],Funsx)
            Body+="}"
        }
        
    }
    if sosdl==1{
                        Bodys+="DLS_So_Start()\n"
                        Body_funcs=`
                        var So_func_map=make(map[string]plugin.Symbol)
                        func DLS_So_Start(){
                            data, _ := ioutil.ReadFile(os.Getenv("WSPPATH")+"/wsp.ini")
                            inis:=strings.Split(string(data),"\n" )
                            for i:=0;i<=len(inis)-1;i++{
                                iniss:=strings.Split(inis[i],"=" )
                                if iniss[0]=="extension"{
                                    So_DLL_vm(iniss[1])
                                }
                            }
                        }
                        func So_DLL_vm(file string)(map[string]plugin.Symbol){
                            p, _ := plugin.Open(file)
                            add, _ := p.Lookup("H_Info")
                            funcmaps:=add.(func() map[int]string)()
                            
                            for i:=0;i<=len(funcmaps)-1;i++{
                                add, _ = p.Lookup(funcmaps[i])
                                So_func_map[funcmaps[i]]=add
                            }
                            return So_func_map
                        }`
        sosdl=3
    }
    return Body
}
func GO_BUILD(s Builds_Struct)(string){
    ALLS_ABD = s
    opcode:=s.Codes
    Funsx := s.Funcs
    Head :="package main \n\n"
    
    BodyT:=Wsp_GTo_Go(opcode,Funsx)
    Body :="func main(){\n"
    Body += Bodys
    Body += "var Vars = make(map[string]interface{})\n"
    Body += BodyT
    Body+="\n}\n"
    
    Import :="import("
    for i:=0;i<=len(imports)-1;i++{
        Import+="\n"+"    \""+imports[i]+"\""
    }
    Import+="\n)\n\n"
    
    fmt.Println("建议使用gofmt -w ./xxx.go 格式化代码\n注意！！！ 如果你引入了动态库 兼容性将会非常差！")
    return Head+Import+Body+Body_func+Body_funcs
}