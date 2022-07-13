package token

import(
  "Wsp/Types"
  "Wsp/Maps"
)

func Text_ADD_strings(a int ,b int,token map[int][4]string)(string){         //在token中通过指定范围，把指定数据集中，并输出字符串 无视空行 
    text:=""
    for i:=a;i<=b;i++{
        text+=token[i][1]
    }
    return text
}
func Text_ADD_strings_KHT(a int ,b int,token map[int][4]string)(string){     //在token中通过指定范围，把指定数据集中，并输出字符串 保留空行
    text:=""
    for i:=a;i<=b;i++{
        if token[i][0] != types.Strings(5){
            text+=token[i][1]
        }else{
            text+="\n"
        }
    }
    return text
}

func Tokens_KUO_X(i int,tokens map[int][4]string)(map[int][4]string){    //获取小括号中的内容，无视空行
    z:=i
    KUO_A :=0
    KUO_B :=0
    KUO_Asu:=0
    KUO_NUM :=1
    for {
        tokenf := tokens[z]
        token_ids :=types.Ints(tokenf[0])
        z++
        if token_ids == 101 {
            KUO_A=z
            KUO_As:=z
            for {
                tokenfs := tokens[KUO_As]
                token_idss :=types.Ints(tokenfs[0])
                if types.Ints(tokens[KUO_As-1][0])==81 && types.Ints(tokens[KUO_As][0])==101{
                    KUO_As++
                    continue
                }
                if token_idss == 101 {
                    KUO_NUM++
                    KUO_Asu++
                }else if token_idss == 102{
                    KUO_NUM--
                }else if KUO_NUM==0{
                    break;
                }
                KUO_As++
            }
            break;
        }
    }
    z=i
    for {
        tokenf := tokens[z]
        token_ids :=types.Ints(tokenf[0])
        z++
        if tokens[z-1][0]==string("81") && token_ids!=101{
            z++
            continue 
        }
        if token_ids == 102 {
            if KUO_Asu==0 {
                KUO_B=z-2
                break;
            }
            KUO_Asu--
        }
    }
    if Text_ADD_strings(KUO_A,KUO_B,tokens) !=""{
        tokens[KUO_A]=[4]string{types.Strings(0),Text_ADD_strings(KUO_A,KUO_B,tokens),"T_TEXT",tokens[KUO_A][3]};
    }else{
        maps.ADD_Map(KUO_A,tokens)
    }
    KUO_A++
    for Z:=KUO_A;Z<=KUO_B;Z++{
        tokens=maps.DEL_Map(KUO_A,tokens)
    }
    return tokens
}

func Tokens_KUO_Z(i int,tokens map[int][4]string)(map[int][4]string){    //获取小括号中的内容，无视空行
    z:=i
    KUO_A :=0
    KUO_B :=0
    KUO_Asu:=0
    KUO_NUM :=1
    for {
        tokenf := tokens[z]
        token_ids :=types.Ints(tokenf[0])
        z++
        if token_ids == 121 {
            KUO_A=z
            KUO_As:=z
            for {
                tokenfs := tokens[KUO_As]
                token_idss :=types.Ints(tokenfs[0])
                if tokens[KUO_As-1][0]=="81" && types.Ints(tokens[KUO_As][0])==121{
                    KUO_As++
                    continue
                }
                if token_idss == 121 {
                    KUO_NUM++
                    KUO_Asu++
                }else if token_idss == 122{
                    KUO_NUM--
                }else if KUO_NUM==0{
                    break;
                }
                KUO_As++
            }
            break;
        }
    }
    z=i
    for {
        tokenf := tokens[z]
        token_ids :=types.Ints(tokenf[0])
        z++
        if tokens[z-1][0]==string("81")&& token_ids!=121{
            z++
            continue 
        }
        if token_ids == 122 {
            if KUO_Asu==0 {
                KUO_B=z-2
                break;
            }
            KUO_Asu--
        }
    }
    if Text_ADD_strings(KUO_A,KUO_B,tokens) !=""{
        tokens[KUO_A]=[4]string{types.Strings(0),Text_ADD_strings(KUO_A,KUO_B,tokens),"T_TEXT",tokens[KUO_A][3]};
    }else{
        maps.ADD_Map(KUO_A,tokens)
    }
    KUO_A++
    for Z:=KUO_A;Z<=KUO_B;Z++{
        tokens=maps.DEL_Map(KUO_A,tokens)
    }
    return tokens
}

func Tokens_KUO_D(i int,tokens map[int][4]string)(map[int][4]string){    //获取大括号中的内容，保留空行
    z:=i
    KUO_A :=0
    KUO_B :=0
    KUO_Asu:=0
    KUO_NUM :=1
    for {
        tokenf := tokens[z]
        token_ids :=types.Ints(tokenf[0])
        z++
        if token_ids == 111 {
            KUO_A=z
            KUO_As:=z
            for {
                tokenfs := tokens[KUO_As]
                token_idss :=types.Ints(tokenfs[0])
                if tokens[KUO_As-1][0]=="81" && types.Ints(tokens[KUO_As][0])==111{
                    KUO_As++
                    continue
                }
                if token_idss == 111 {
                    KUO_NUM++
                    KUO_Asu++
                }else if token_idss == 112{
                    KUO_NUM--
                }else if KUO_NUM==0{
                    break;
                }
                KUO_As++
            }
            break;
        }
    }
    z=i
    for {
        tokenf := tokens[z]
        token_ids :=types.Ints(tokenf[0])
        z++
        if tokens[z-1][0]==string("81")&& token_ids!=111{
            z++
            continue 
        }
        if token_ids == 112 {
            if KUO_Asu==0 {
                KUO_B=z-2
                break;
            }
            KUO_Asu--
        }
    }
    if Text_ADD_strings(KUO_A,KUO_B,tokens) !=""{
        tokens[KUO_A]=[4]string{types.Strings(0),Text_ADD_strings_KHT(KUO_A,KUO_B,tokens),"T_TEXT",tokens[KUO_A][3]};
    }else{
        maps.ADD_Map(KUO_A,tokens)
    }
    KUO_A++
    for Z:=KUO_A;Z<=KUO_B;Z++{
        tokens=maps.DEL_Map(KUO_A,tokens)
    }
    return tokens
}
func Wsp_Grammar(tokens map[int][4]string)(map[int][4]string){
    for i:=0;i<=len(tokens)-1;i++{
        token := tokens[i]
        token_id :=types.Ints(token[0])
        if token_id==12{            //PRINT输出语句
            tokens=Tokens_KUO_X(i,tokens)
        }else if token_id==0 && types.Ints(tokens[i+1][0])==101 && types.Ints(tokens[i-1][0])!=10{   //用户自定义函数
            tokens[i]=[4]string{types.Strings(200),tokens[i][1],"H_V_"+tokens[i][1],tokens[i][3]};
            i--
        }else if token_id==10{       //用户定义自定义函数
            tokens=Tokens_KUO_X(i,tokens)
            tokens=Tokens_KUO_D(i,tokens)
        }else if token_id==25 {      //if语句
            tokens=Tokens_KUO_X(i,tokens)
            tokens=Tokens_KUO_D(i,tokens)
        }else if token_id==26 {      //if语句else分支
            IS_TRUE_As:=0
            for z:=i;z<=len(tokens)-1;z++{
                if tokens[z][0]==types.Strings(25){
                    IS_TRUE_As=1
                }else if tokens[z][0]==types.Strings(111){
                    break
                }
            }
            if IS_TRUE_As == 1{
                tokens=Tokens_KUO_X(i,tokens)
            }
            tokens=Tokens_KUO_D(i,tokens)
            i=i+2
            
        }else if token_id==11{         //for循环
            tokens=Tokens_KUO_X(i,tokens)
            tokens=Tokens_KUO_D(i,tokens)
        }else if token_id==121{
            tokens=Tokens_KUO_Z(i-1,tokens)
        }else if token_id==400{
            tokens=Tokens_KUO_X(i,tokens)
        }else if token_id==200{
            tokens=Tokens_KUO_X(i,tokens)
        }else if token_id==401{
            tokens=Tokens_KUO_X(i,tokens)
        }else if token_id==402{
            tokens=Tokens_KUO_X(i,tokens)
        }else{
            continue
        }
    }
    return tokens
}