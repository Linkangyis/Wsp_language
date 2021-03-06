package vm

import(
  "Wsp/Token"
  "Wsp/Build"
  "Wsp/Echo"
  "Wsp/Types"
  "Wsp/Maps"
  "Wsp/WVM/Array"
  "io/ioutil"
  "plugin"
  "fmt"
  "strings"
  "os"
)
//var Vars = make(map[string]string)
type vm_func func (parameter Builds_Parameter)(string)
var vm_s =make(map[int]vm_func)
var code_ok = make(map[int]string)
var So_func_map=make(map[string]plugin.Symbol)
var wsp_func_del []string

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
func CodesOk(s map[int]string){
    code_ok=s
}
func CodesOkre()(map[int]string){
    return code_ok
}
func E_So_func_map()(map[string]plugin.Symbol){
    return So_func_map;
}
func DLS_So_Start()(map[string]string){
    re := make(map[string]string)
    debugs = 0
    data, _ := ioutil.ReadFile(os.Getenv("WSPPATH")+"/wsp.ini")
    inis:=strings.Split(string(data),"\n" )
    for i:=0;i<=len(inis)-1;i++{
        iniss:=strings.Split(inis[i],"=" )
        if iniss[0]=="extension"{
            So_DLL_vm(iniss[1])
        }else if iniss[0]=="wsp_debug" && iniss[1]=="1"{
            debugs = 1
        }else if iniss[0]=="wsp_func_del"{
            wsp_func_del = strings.Split(iniss[1], ",")
        }else if iniss[0] != ""{
            re[iniss[0]]=iniss[1]
        }
    }
    return re
}
func Wsp_File_E()(string){
    return token.Wsp_File_P();
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

func Parameter_processing(a string)(map[int]string,map[int]string){
    map_snum:=0
    returns :=make(map[int]string)
    tokenser:=token.Wsp_Grammar(token.Wsp_Semantic(token.Wsp_Lexical_func(a)))
    for i:=0;i<=len(tokenser)-1;i++{
        if tokenser[i][1]==","{
            map_snum++
        }else{
            returns[map_snum]+=tokenser[i][1]
        }
    }
    returnss :=make(map[int]string)
    for i:=0;i<=len(returns)-1;i++{
        returnss[i]=Var_so_all(returns[i])
    }
    return returnss,returns
}

func print_vm(parameter Builds_Parameter)(string){
    a:=parameter.a
    ms,_:=Parameter_processing(a)
    rs:=ms[0]
    fmt.Println(rs)
    return rs
}
func add_vm(parameter Builds_Parameter)(string){
    a:=parameter.a
    
    str_arr := strings.Split(a, ",")
    add_num:=0
    for i:=0;i<=len(str_arr)-1;i++{
        if string(str_arr[i][0]) == string("$"){
            add_num+=types.Ints(array.Read_Array(strings.Replace(str_arr[i],"$","",-1)))
        }else{
            add_num+=types.Ints(str_arr[i])
        }
    }
    return types.Strings(add_num)
}

func arrays(Arrs string)(string){
    lock := 0
    avrs := make(map[int]string)
    avrs_l := 0
    for i:=0;i<=len(Arrs)-1;i++{
        if string(Arrs[i])=="]"{
            lock--
        }else if string(Arrs[i])=="["{
            lock++
        }
        avrs[avrs_l]+=string(Arrs[i])
        if lock==0{
            avrs_l++
        }
    }
    res :=""
    for i:=0;i<=len(avrs)-1;i++{
        res+="["+Var_so_all(types.Strings_so(avrs[i]))+"]"
    }
    return res
}
func vars_vm_array(parameter Builds_Parameter)(string){
    a:=parameter.a
    c:=parameter.c
    opcode:=parameter.opcode
    lens:=parameter.lens
    fs:=parameter.fs
    ft:=parameter.ft
    
    if opcode[lens][0]=="302" || opcode[lens][0]=="304" {
        a=opcode[lens][1]+arrays(parameter.b)
    }
    if opcode[types.Ints(c)][0]=="0"{
        array.Add_Array(a,Var_so_all(opcode[types.Ints(c)][1]))
    }else if opcode[types.Ints(c)][0]=="302"{
        as:=opcode[types.Ints(c)][1]+arrays(opcode[types.Ints(c)][3])
        array.Copy_Array(as,a)
        //Vars[a]=Vars[as]
    }else if opcode[types.Ints(c)][0]=="300"{
        array.Add_Array(a,array.Read_Array(opcode[types.Ints(c)][1]))
        //Vars[a]=Vars[opcode[types.Ints(c)][1]]
    }else if opcode[types.Ints(c)][0]=="304" || opcode[types.Ints(c)][0]=="301"{
        for i:=lens;i<=len(opcode)-1;i++{
            if opcode[i][0]==types.Strings(0){
                array.Add_Array(a,Var_so_all(opcode[i][1]))
                break
            }else if opcode[i][0]==types.Strings(302){
                as:=opcode[i][1]+arrays(opcode[i][3])
                //Vars[a]=Vars[as]
                array.Copy_Array(as,a)
                break
            }else if opcode[i][0]==types.Strings(300){
                array.Add_Array(a,array.Read_Array(opcode[i][2]))
                //Vars[a]=Vars[opcode[i][2]]
                break
            }else if opcode[i][0]!=types.Strings(304) && opcode[i][0]!=types.Strings(301){
                if v, ok := code_ok[i]; ok {
                    array.Add_Array(a,v)
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
                    array.Add_Array(a,vm_s[types.Ints(opcode[i][0])](Buildse))
                    code_ok[i] = array.Read_Array(a)
                }
                break
            }
        }
    }else{
        if v, ok := code_ok[types.Ints(c)]; ok {
            array.Add_Array(a,v)
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
            array.Add_Array(a,vm_s[types.Ints(opcode[types.Ints(c)][0])](Buildse))
            code_ok[types.Ints(c)] = array.Read_Array(a)
        }
    }
    return "True"
}

func code_null(parameter Builds_Parameter)(string){
    return "NULLS"
}

func Var_so_all(var_name string)(string){
    var_name = types.Trims(var_name)
    if var_name=="TRUE"{
        return "1"
    }else if var_name=="FALSE"{
        return "0"
    }
    if string(var_name[0])=="$"{
        bsd:=build.Wsp_Build(token.Wsp_Semantic(token.Wsp_Grammar(token.Wsp_Lexical_func(var_name))))
        if bsd.Codes[0][0]=="300"{
            return array.Read_Array(types.Var_so(var_name))
        }else{
            as:=bsd.Codes[0][1]+arrays(bsd.Codes[0][3])
            tres := array.Get_All_Array(array.So_Array_Stick(as))
            var returns string
            if tres ==""{
                returns = array.Read_Array(as)
            }else{
                returns = array.Read_Array(as)
                //returns = "array("+array.Get_All_Array(array.So_Array_Stick(as))+")"
            }
            return returns
        }
    }else if types.IsNum(var_name){
        return var_name
    }else if string(var_name[0])=="\""{
        return types.Strings_so(types.Trims(var_name))
    }else{
        
        bd:=build.Wsp_Build(token.Wsp_Semantic(token.Wsp_Grammar(token.Wsp_Lexical_func(var_name))))
        Buildse:=Builds_Parameter{
            a      : bd.Codes[0][2],
            b      : bd.Codes[0][3],
            c      : bd.Codes[0][4],
            opcode : bd.Codes,
            lens   : 0,
            fs     : bd.Funcs,
            ft     : bd.Funcs_list,
        }
        res := vm_s[types.Ints(bd.Codes[0][0])](Buildse)
        //fmt.Println(2)
        return res
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
    addorabb:=lone[1][1]
    if addorabb==""{
        if Var_so_all(a)=="1"{
            vm_funcs_l(fs[c],fs,ft,"(NULL TO IF)")
            return "TRUE"
        }
    }
    A:=lone[0][1]
    B:=lone[3][1]
    lock :=0
    for i:=lens;i<=len(opcode)-1;i++{
        if opcode[i][0]=="25"&&lock==0{
            if ifs(A,B,addorabb){
                vm_funcs_l(fs[c],fs,ft,"(NULL TO IF)")
                break
            }
            lock=1
        }else if opcode[i][0]=="28"{
            lone=token.Wsp_Semantic(token.Wsp_Grammar(token.Wsp_Lexical_var(opcode[i][2])))
            addorabb=lone[1][1]
            A=lone[0][1]
            B=lone[3][1]
            if addorabb==""{
                if Var_so_all(opcode[i][2])=="1"{
                    vm_funcs_l(fs[opcode[i][4]],fs,ft,"(NULL TO IF)")
                    return "TRUE"
                }
            }
            if ifs(A,B,addorabb){
                vm_funcs_l(fs[opcode[i][4]],fs,ft,"(NULL TO IF)")
                break
            }
        }else if opcode[i][0]=="26"{
            vm_funcs_l(fs[opcode[i][4]],fs,ft,"(NULL TO IF)")
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
    ft:=Ec_Ft()
    
    function_name:=strings.Trim(opcode[lens][1]," ")
    
    for i:=0;i<=len(wsp_func_del)-1;i++{
        if wsp_func_del[i]==function_name{
            array.Del_Dir("./Var_Temps")
            fmt.Println("\n",echo.Arr_Echo_Opcode_View_r(50),"\n ????????????!! \n ????????????:",opcode[lens][5],"\n ????????????:",function_name+"(",a,") \n ?????????????????????"+function_name+"?????????\n",echo.Arr_Echo_Opcode_View_r(50),"\n")
            os.Exit(0)
        }
    }
    
    temps:=array.Read_Paths()
    
    returns := ""
    if _, ok := fs[function_name]; ok {
        array.Set_Res(0)
        vars_chuan,_ := Parameter_processing(a)
        array.Set_Paths(temps+function_name+"/")
        vars_ding := strings.Split(ft[function_name],"," )
        for i:=0;i<=len(vars_chuan)-1;i++{
            array.Add_Array(types.Var_so(vars_ding[i]),vars_chuan[i])
        }
        array.Set_Paths(temps+function_name+"/")
        returns = vm_funcs(fs[function_name],fs,ft,function_name)
        array.Del_Dirs(array.Read_Paths())
        array.Set_Ress(array.Read_Paths())
        array.Set_Paths(temps)
    }else if  _, oks := So_func_map[function_name]; oks {
        array.Set_Paths(temps)
        returns=So_func_map[function_name].(func(string) string)(a)
    }else{
        array.Del_Dir("./Var_Temps")
        fmt.Println("\n",echo.Arr_Echo_Opcode_View_r(50),"\n ????????????!! \n ????????????:",opcode[lens][5],"\n ????????????:",function_name+"(",a,") \n ?????????????????????"+function_name+"?????????\n",echo.Arr_Echo_Opcode_View_r(50),"\n")
        os.Exit(0)
    }
    
    array.Set_Res(1)
    return returns
}

func vars_fors_vars(varse string,vs int){
    array.Add_Array(varse,types.Strings(vs))
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
    if array.Read_Array(bls)!="NULL"{
        tmeps=array.Read_Array(bls)
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
        array.Add_Array(bls,tmeps)
    }else{
        array.Del_Array(bls)
    }
    return "TRUE"
}

func vars_csc(parameter Builds_Parameter)(string){
    a:=parameter.a
    as,_:=Parameter_processing(a)
    return as[0]
}
var ALLS_OPCODE=make(map[int][6]string)
func Ec_Op()(map[int][6]string){
    return ALLS_OPCODE
}
var ALLS_FS=make(map[string]map[int][6]string)
func Ec_Fs()(map[string]map[int][6]string){
    return ALLS_FS
}
var ALLS_FT=make(map[string]string)
func Ec_Ft()(map[string]string){
    return ALLS_FT
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
    vm_s[402] = vars_csc
    
    Builds:=Buildse.Codes
    ALLS_OPCODE = Buildse.Codes
    ALLS_FS = Buildse.Funcs
    ALLS_FT = Buildse.Funcs_list
    for i:=0;i<=len(Builds)-1;i++{
        array.Del_Dirl()
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
        echo.Arr_Echo_Opcode_View(Builds,array.Get_All_Array(array.Read_Paths()),"(NULL)")
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
    vm_s[402] = vars_csc
    
    code_ok_f:=maps.MAP_COPY_codeok(code_ok)
    code_ok = make(map[int]string)
    returns := ""
    for i:=0;i<=len(Builds)-1;i++{
        array.Del_Dirl()
        if Builds[i][0]==types.Strings(27){
            returns = array.Read_Array(Builds[i+1][1])
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
    vm_s[402] = vars_csc
    
    code_ok_f:=maps.MAP_COPY_codeok(code_ok)
    code_ok = make(map[int]string)
    returns := ""
    for i:=0;i<=len(Builds)-1;i++{
        array.Del_Dirl()
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

func Vm_Code_Run(Builds map[int][6]string)(string){
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
    vm_s[402] = vars_csc
    
    code_ok_f:=maps.MAP_COPY_codeok(code_ok)

    code_ok = make(map[int]string)
    returns := ""
    for i:=0;i<=len(Builds)-1;i++{
        array.Del_Dirl()
        if v, ok := code_ok[i]; ok {
            code_ok[i] = v
        }else{
            Buildse:=Builds_Parameter{
                a      : Builds[i][2],
                b      : Builds[i][3],
                c      : Builds[i][4],
                opcode : Builds,
                lens   : i,
                fs     : ALLS_FS,
                ft     : ALLS_FT,
            }
            code_ok[i] = vm_s[types.Ints(Builds[i][0])](Buildse)
        }
    }
    
    code_ok = code_ok_f
    return returns
}