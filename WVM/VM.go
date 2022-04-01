package vm

import(
  "../Token"
  "../Build"
  "../Echo"
  "../Types"
  "../Maps"
  "io/ioutil"
  "plugin"
  "fmt"
  "strings"
  "os"
)
var Vars = make(map[string]string)
type vm_func func (parameter Builds_Parameter)(string)
var vm_s =make(map[int]vm_func)
var code_ok = make(map[int]string)
var So_func_map=make(map[string]plugin.Symbol)

type Builds_Parameter struct {
    a  string
    b  string
    c  string
    opcode map[int][6]string
    lens int
    fs map[string]map[int][6]string
    ft map[string]string
}
var debugs int
func DLS_So_Start(){
    debugs = 0
    data, _ := ioutil.ReadFile(os.Getenv("WSPPATH")+"/wsp.ini")
    inis:=strings.Split(string(data),"\n" )
    for i:=0;i<=len(inis)-1;i++{
        iniss:=strings.Split(inis[i],"=" )
        if iniss[0]=="extension"{
            So_DLL_vm(iniss[1])
        }else if iniss[0]=="wsp_debug" && iniss[1]=="1"{
            debugs = 1
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
}
func print_vm(parameter Builds_Parameter)(string){
    a:=parameter.a
    opcode:=parameter.opcode
    lens:=parameter.lens
    if string(a[0])==string("$"){
        fmt.Println(Vars[types.Var_so(a)])
    }else if string(a[0])==string("\""){
        fmt.Println(types.Strings_so(a))
    }else if types.IsNum(a){
        fmt.Println(a)
    }else{
        fmt.Println("\n",echo.Arr_Echo_Opcode_View_r(50),"\n 运行错误!! \n 错误行数:",opcode[lens][5],"\n 错误内容:","print(",a,") \n 错误原因：其中括号内容存在异常，请检查\n",echo.Arr_Echo_Opcode_View_r(50),"\n")
        os.Exit(0)
    }
    return a
}
func add_vm(parameter Builds_Parameter)(string){
    a:=parameter.a
    
    str_arr := strings.Split(a, ",")
    add_num:=0
    for i:=0;i<=len(str_arr)-1;i++{
        if string(str_arr[i][0]) == string("$"){
            add_num+=types.Ints(Vars[strings.Replace(str_arr[i],"$","",-1)])
        }else{
            add_num+=types.Ints(str_arr[i])
        }
    }
    return types.Strings(add_num)
}

func vars_vm_array(parameter Builds_Parameter)(string){
    a:=parameter.a
    c:=parameter.c
    opcode:=parameter.opcode
    lens:=parameter.lens
    fs:=parameter.fs
    ft:=parameter.ft
    
    if opcode[types.Ints(c)][0]=="0"{
        Vars[a]=opcode[types.Ints(c)][1]
    }else if opcode[types.Ints(c)][0]=="302"{
        Vars[a]=Vars[opcode[types.Ints(c)][2]]
    }else if opcode[types.Ints(c)][0]=="300"{
        Vars[a]=Vars[opcode[types.Ints(c)][1]]
    }else if opcode[types.Ints(c)][0]=="304" || opcode[types.Ints(c)][0]=="301"{
        for i:=lens;i<=len(opcode)-1;i++{
            if opcode[i][0]==types.Strings(0){
                Vars[a]=opcode[i][1]
                break
            }else if opcode[i][0]==types.Strings(302){
                Vars[a]=Vars[opcode[i][2]]
                break
            }else if opcode[i][0]==types.Strings(300){
                Vars[a]=Vars[opcode[i][2]]
                break
            }else if opcode[i][0]!=types.Strings(304) && opcode[i][0]!=types.Strings(301){
                if v, ok := code_ok[i]; ok {
                    Vars[a]=v
                }else{
                    Buildse:=Builds_Parameter{
                        a      : opcode[i][2],
                        b      : opcode[i][3],
                        c      : opcode[i][4],
                        opcode : opcode,
                        lens   : lens+1,
                        fs     : fs,
                        ft     : ft,
                     }
                    Vars[a]=vm_s[types.Ints(opcode[i][0])](Buildse)
                    code_ok[i] = Vars[a]
                }
                break
            }
        }
    }else{
        if v, ok := code_ok[types.Ints(c)]; ok {
            Vars[a]=v
        }else{
            Buildse:=Builds_Parameter{
                a      : opcode[types.Ints(c)][2],
                b      : opcode[types.Ints(c)][3],
                c      : opcode[types.Ints(c)][4],
                opcode : opcode,
                lens   : lens+1,
                fs     : fs,
                ft     : ft,
            }
            Vars[a]=vm_s[types.Ints(opcode[types.Ints(c)][0])](Buildse)
            code_ok[types.Ints(c)] = Vars[a]
        }
    }
    return "True"
}

func code_null(parameter Builds_Parameter)(string){

    return "NULL"
}

func Var_so_all(var_name string)(string){
    if string(var_name[0])=="$"{
        return Vars[types.Var_so(var_name)]
    }else if types.IsNum(var_name){
        return var_name
    }else if string(var_name[0])=="\""{
        return types.Strings_so(var_name)
    }else{
        bd:=build.Wsp_Build(token.Wsp_Semantic(token.Wsp_Grammar(token.Wsp_Lexical_func(var_name))))
        Buildse:=Builds_Parameter{
            a      : bd.Codes[0][2],
            b      : bd.Codes[0][3],
            c      : bd.Codes[0][4],
            opcode : bd.Codes,
            lens   : 99999,
            fs     : bd.Funcs,
            ft     : bd.Funcs_list,
        }
        
        return vm_s[types.Ints(bd.Codes[0][0])](Buildse)
    }
}

func if_vm(parameter Builds_Parameter)(string){
    a:=parameter.a
    c:=parameter.c
    opcode:=parameter.opcode
    lens:=parameter.lens
    fs:=parameter.fs
    ft:=parameter.ft
    lone:=token.Wsp_Semantic(token.Wsp_Grammar(token.Wsp_Lexical_var(a)))
    addorabb:=lone[2][1]
    A:=lone[0][1]
    B:=lone[3][1]
    lock :=0
    for i:=lens;i<=len(opcode)-1;i++{
        if opcode[i][0]=="25"&&lock==0{
            if ifs(A,B,addorabb){
                vm_funcs_l(fs[c],fs,ft,"(NULL TO FOR)")
                break
            }
            lock=1
        }else if opcode[i][0]=="28"{
            lone=token.Wsp_Semantic(token.Wsp_Grammar(token.Wsp_Lexical_var(opcode[i][2])))
            addorabb=lone[2][1]
            A=lone[0][1]
            B=lone[3][1]
            if ifs(A,B,addorabb){
                vm_funcs_l(fs[opcode[i][4]],fs,ft,"(NULL TO FOR)")
                break
            }
        }else if opcode[i][0]=="26"{
            vm_funcs_l(fs[opcode[i][4]],fs,ft,"(NULL TO FOR)")
            break
        }else{
            break
        }
    }
    
    return "TRUE"
}
func ifs(a string ,b string,t string)(bool){
    if t=="="{
        if Var_so_all(a)==Var_so_all(b){
            return true
        }else{
            return false
        }
    }else if t==">"{
        if Var_so_all(a)>=Var_so_all(b){
            return true
        }else{
            return false
        }
    }else if t=="<"{
        if Var_so_all(a)<=Var_so_all(b){
            return true
        }else{
            return false
        }
    }else if t=="!"{
        if Var_so_all(a)!=Var_so_all(b){
            return true
        }else{
            return false
        }
    }
    return false
}

func funcs_vm_run(parameter Builds_Parameter)(string){
    a:=parameter.a
    opcode:=parameter.opcode
    lens:=parameter.lens
    fs:=parameter.fs
    ft:=parameter.ft
    
    function_name:=opcode[lens][1]
    vars_chuan := strings.Split(a, ",")
    Var_tmps:=maps.MAP_COPY_vars(Vars)
    for key,_ := range Vars {
        Vars[key]=""
        delete(Vars,key)
    }
    returns := ""
    vars_ding := strings.Split(ft[function_name],"," )
    if _, ok := fs[function_name]; ok {
        for i:=0;i<=len(vars_chuan)-1;i++{
            Vars[types.Var_so(vars_ding[i])]=vars_chuan[i]
        }
        returns = vm_funcs(fs[function_name],fs,ft,function_name)
        Vars = Var_tmps
    }else if  _, oks := So_func_map[function_name]; oks {
        So_func_map[function_name].(func(string) string)(a)
        Vars = Var_tmps
    }else{
        fmt.Println("\n",echo.Arr_Echo_Opcode_View_r(50),"\n 运行错误!! \n 错误行数:",opcode[lens][5],"\n 错误内容:",function_name+"(",a,") \n 错误原因：函数"+function_name+"不存在\n",echo.Arr_Echo_Opcode_View_r(50),"\n")
        os.Exit(0)
    }
    return returns
}
func vars_fors_vars(varse string,vs int){
    Vars[varse]=types.Strings(vs)
}
func for_vm(parameter Builds_Parameter)(string){
    a:=parameter.a
    c:=parameter.c
    fs:=parameter.fs
    ft:=parameter.ft
    lone := make(map[int]map[int][4]string)
    vars_chuan := strings.Split(a, ";")
    for i:=0;i<=len(vars_chuan)-1;i++{
        lone[i]=token.Wsp_Semantic(token.Wsp_Grammar(token.Wsp_Lexical_func(vars_chuan[i])))
    }
    lioen := types.Ints(lone[0][3][1])
    tj := lone[1][2][1]
    addfornum:=types.Ints(lone[2][0][1])
    addfornums:=lone[1][4][1]
    addorabb:=lone[3][3][1]
    bls:=lone[3][1][1]
    lock := 0
    tmeps :=""
    if _, ok := Vars[bls]; ok {
        tmeps=Vars[bls]
    }else{
        lock=1
    }
    if addorabb=="+"{
        if tj==">"{
            for i:=lioen;i>=types.Ints(addfornums);i=i+addfornum{
                vars_fors_vars(bls,i)
                vm_funcs_l(fs[c],fs,ft,"(NULL TO FOR)")
            }
        }else if tj=="<"{
            for i:=lioen;i<=types.Ints(addfornums);i=i+addfornum{
                vars_fors_vars(bls,i)
                vm_funcs_l(fs[c],fs,ft,"(NULL TO FOR)")
            }
        }else if tj=="="{
            for i:=lioen;i==types.Ints(addfornums);i=i+addfornum{
                vars_fors_vars(bls,i)
                vm_funcs_l(fs[c],fs,ft,"(NULL TO FOR)")
            }
        }else if tj=="!"{
            for i:=lioen;i!=types.Ints(addfornums);i=i+addfornum{
                vars_fors_vars(bls,i)
                vm_funcs_l(fs[c],fs,ft,"(NULL TO FOR)")
            }
        }
    }else{
        if tj==">"{
            for i:=lioen;i>=types.Ints(addfornums);i=i-addfornum{
                vars_fors_vars(bls,i)
                vm_funcs_l(fs[c],fs,ft,"(NULL TO FOR)")
            }
        }else if tj=="<"{
            for i:=lioen;i<=types.Ints(addfornums);i=i-addfornum{
                vars_fors_vars(bls,i)
                vm_funcs_l(fs[c],fs,ft,"(NULL TO FOR)")
            }
        }else if tj=="="{
            for i:=lioen;i==types.Ints(addfornums);i=i-addfornum{
                vars_fors_vars(bls,i)
                vm_funcs_l(fs[c],fs,ft,"(NULL TO FOR)")
            }
        }else if tj=="!"{
            for i:=lioen;i!=types.Ints(addfornums);i=i-addfornum{
                vars_fors_vars(bls,i)
                vm_funcs_l(fs[c],fs,ft,"(NULL TO FOR)")
            }
        }
    }
    if lock == 0{
        Vars[bls] = tmeps
    }else{
        delete(Vars,bls)
    }
    return "TRUE"
}
func Wsp_VM(Buildse build.Builds_Struct){
    vm_s[12]=print_vm
    vm_s[301] = vars_vm_array
    vm_s[0] = code_null
    vm_s[300] = code_null
    vm_s[401] = add_vm
    vm_s[304] = vars_vm_array
    vm_s[302] = code_null
    vm_s[10] = code_null
    vm_s[200] = funcs_vm_run
    vm_s[11] = for_vm
    vm_s[25] = if_vm
    vm_s[26] = code_null
    vm_s[28] = code_null
    
    
    Builds:=Buildse.Codes
    
    for i:=0;i<=len(Builds)-1;i++{
        if v, ok := code_ok[i]; ok {
            code_ok[i] = v
        }else{
            Buildse:=Builds_Parameter{
                a      : Builds[i][2],
                b      : Builds[i][3],
                c      : Builds[i][4],
                opcode : Builds,
                lens   : i,
                fs     : Buildse.Funcs,
                ft     : Buildse.Funcs_list,
            }
            code_ok[i] = vm_s[types.Ints(Builds[i][0])](Buildse)
        }
    }
    if debugs==1{
        echo.Arr_Echo_Opcode_View(Builds,Vars,"(NULL)")
    }
}
func vm_funcs(Builds map[int][6]string,fs map[string]map[int][6]string,ft map[string]string,funcname string)(string){
    vm_s[12]=print_vm
    vm_s[301] = vars_vm_array
    vm_s[0] = code_null
    vm_s[300] = code_null
    vm_s[401] = add_vm
    vm_s[304] = vars_vm_array
    vm_s[302] = code_null
    vm_s[10] = code_null
    vm_s[200] = funcs_vm_run
    vm_s[27] = code_null
    vm_s[11] = for_vm
    vm_s[26] = code_null
    
    code_ok_f:=maps.MAP_COPY_codeok(code_ok)
    for i:=0;i<=len(code_ok)-1;i++{
        delete(code_ok,i)
    }
    returns := ""
    for i:=0;i<=len(Builds)-1;i++{
        if Builds[i][0]==types.Strings(27){
            returns = Vars[Builds[i+1][1]]
            break
        }
        if v, ok := code_ok[i]; ok {
            code_ok[i] = v
        }else{
            Buildse:=Builds_Parameter{
                a      : Builds[i][2],
                b      : Builds[i][3],
                c      : Builds[i][4],
                opcode : Builds,
                lens   : i,
                fs     : fs,
                ft     : ft,
            }
            code_ok[i] = vm_s[types.Ints(Builds[i][0])](Buildse)
        }
    }
    
    code_ok = code_ok_f
    return returns
}


func vm_funcs_l(Builds map[int][6]string,fs map[string]map[int][6]string,ft map[string]string,funcname string)(string){
    vm_s[12]=print_vm
    vm_s[301] = vars_vm_array
    vm_s[0] = code_null
    vm_s[300] = code_null
    vm_s[401] = add_vm
    vm_s[304] = vars_vm_array
    vm_s[302] = code_null
    vm_s[10] = code_null
    vm_s[200] = funcs_vm_run
    vm_s[27] = code_null
    vm_s[11] = for_vm
    
    code_ok_f:=maps.MAP_COPY_codeok(code_ok)
    for i:=0;i<=len(code_ok)-1;i++{
        delete(code_ok,i)
    }
    returns := ""
    for i:=0;i<=len(Builds)-1;i++{
        if v, ok := code_ok[i]; ok {
            code_ok[i] = v
        }else{
            Buildse:=Builds_Parameter{
                a      : Builds[i][2],
                b      : Builds[i][3],
                c      : Builds[i][4],
                opcode : Builds,
                lens   : i,
                fs     : fs,
                ft     : ft,
            }
            code_ok[i] = vm_s[types.Ints(Builds[i][0])](Buildse)
        }
    }
    
    code_ok = code_ok_f
    return returns
}