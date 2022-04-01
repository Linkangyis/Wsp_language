package token

func token_map()(map[string]int){     //token对应的token序号
    var maps = map[string]int{}
    maps["("] = 101;
    maps[")"] = 102;
    maps["{"] = 111;
    maps["}"] = 112;
    maps["["] = 121;
    maps["]"] = 122;
    maps["\n"] = 5;
    maps[" "] = 6;
    maps["="] = 7;
    maps[">"] = 8;
    maps["<"] = 9;
    maps["+"] = 500;
    maps["function"] = 10;
    maps["for"] = 11;
    maps["print"] = 12;
    maps["class"] = 20;
    maps["public"] = 21;
    maps["new"] = 21;
    maps["-"] = 22;
    maps["eval"] = 23;
    maps["null"] = 24;
    maps["if"] = 25;
    maps["else"] = 26;
    maps["return"] = 27;
    maps["\\"] = 81;
    maps["/"] = 82;
    maps["$"] = 300;
    maps["array"] = 400;
    maps["ADD"] = 401;
    return maps;
}
func token_text_map()(map[string]string){    //token对应的token值
    var maps = map[string]string{}
    maps["("] = "D_X_KUO";
    maps[")"] = "D_X_KUO";
    maps["{"] = "D_H_KUO";
    maps["}"] = "D_H_KUO";
    maps["["] = "D_Z_KUO";
    maps["]"] = "D_Z_KUO";
    maps["\n"] = "T_KH";
    maps[" "] = "T_K";
    maps["="] = "Y_DE";
    maps[">"] = "Y_DA - CLASS_START_B";
    maps["<"] = "Y_XA";
    maps["+"] = "ADD";
    maps["function"] = "H_FUNCTION";
    maps["for"] = "H_FOR";
    maps["print"] = "H_PRINT";
    maps["class"] = "CLASS_SET";
    maps["public"] = "CLASS_PUBLIC";
    maps["new"] = "CLASS_NEW";
    maps["-"] = "CLASS_START_A";
    maps["eval"] = "H_EVAL";
    maps["null"] = "V_NULL";
    maps["if"] = "H_IF";
    maps["else"] = "H_IF_ELSE";
    maps["return"] = "H_RETURN";
    maps["\\"] = "T_TS_SWS";
    maps["/"] = "T_ZS";
    maps["$"] = "V_VAR";
    maps["array"] = "ARRAY";
    maps["ADD"] = "ADD";
    return maps;
}
