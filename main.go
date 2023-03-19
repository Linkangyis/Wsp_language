package main

import(
    "Wsp/Module/InitCode"
    "Wsp/Module/Memory"
    "Wsp/Analysis/Ast"
    "Wsp/Compile"
    "fmt"
)

func main(){
    Code := `
    /*
        这些是注释
        这些是注释
        这些是注释
        这些是注释
        2023/2/19
    */
    Sys.Println(1);
    /*123123*/Sys.Println(1);
    //12123
    
    1++1;
    $i++;
    function($a){
        Sys.Println(123);
    }
    Sys.Println(1)+1;
    (1+1+2+3,Test(1+1));
    
    (function Test(  $a,  $b){
        Sys.Println(1);
    },1);
    
    $a = "123123";
    $a = '123123';
    $a = 123123;
    $a = 123.123;
    function Test(  $a,  $b){
        Sys.Println(1)
    }
    $a=-1+1--1;
    $a++;
    $a--;
    $a[0]++;
    $a[0]--;
    $a[1]--;
    ($a++)+1;
    ($a[0])--2;
    $a(123)->124+Test("123")*3*1>=1&&1>=0;
    `
    
    id := memory.Malloc()
    id.Open().SetValue(initCode.Complex(Code).Ast)
    fmt.Println("-----------------------------------------------")
    for i:=0;i<=len((*id.Open().Read()).(map[int]ast.AstStruct));i++{
        fmt.Println(i,(*id.Open().Read()).(map[int]ast.AstStruct)[i])
    }
    fmt.Println("-----------------------------------------------")
    Opcode:=compile.Compile((*id.Open().Read()).(map[int]ast.AstStruct))
    fmt.Println("-----------------------------------------------")
    for i:=0;i<=len(Opcode.Opcode)-1;i++{
        fmt.Println(i,"OPCODE",Opcode.Opcode[i])
    }
    fmt.Println("-----------------------------------------------")
    fmt.Println(Opcode.FuncList)
    fmt.Println("-----------------------------------------------")
    memory.FreeAll()
}
