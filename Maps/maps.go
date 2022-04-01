package maps

import(
  "../Types"
)

func DEL_Map(a int ,token map[int][4]string)(map[int][4]string){  //删除map中的一个指定元素
    for i:=a;i<=len(token)-1;i++{
        token[i]=token[i+1]
    }
    delete(token,len(token)-1)
    return token
}
func MAP_COPY(token map[int][4]string)(map[int][4]string){         //copy一个map
    tok := make(map[int][4]string)
    for i:=0;i<=len(token)-1;i++{
        tok[i]=token[i]
    }
    return tok
}
func MAP_COPY_vars(token map[string]string)(map[string]string){         //copy一个map
    tok := make(map[string]string)
    for key,vu := range token {
        tok[key]=vu
    }
    return tok
}
func MAP_COPY_codeok(token map[int]string)(map[int]string){         //copy一个map
    tok := make(map[int]string)
    for key,vu := range token {
        tok[key]=vu
    }
    return tok
}
func ADD_Map(a int ,token map[int][4]string)(map[int][4]string){    //在map指定位置ADD一个T_TEXT元素
    token_tmp:=MAP_COPY(token)
    tmps:=token[a]
    token[a] = [4]string{types.Strings(0),"","T_TEXT",tmps[3]};
    a++
    for i:=a;i<=len(token_tmp)-1;i++{
        t_a:=token_tmp[i-1]
        token[i]=t_a
    }
    
    return token
}