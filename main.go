package main

import (
	"Wsp/Analysis/Ast"
	"Wsp/Compile"
	"Wsp/Module/InitCode"
	"Wsp/Module/Memory"
	vm "Wsp/Module/Vm"
	"fmt"
)

func main() {
	Code := `

	Println("TestPrint")
	Println("TestPrint")
	Println("TestPrint")
	Println("TestPrint")
	Println("TestPrint")
    /*
        这些是注释
        这些是注释
        这些是注释
        这些是注释
        2023/2/19
    */


    /*
    function TestFunc($a){
        Sys.Println($a);
    }
    
    
    function TestFunc($a,$b){
        Sys.Println($a,$b);
    }
    
    class TestClassA{
        function Aclass(){}
        Sys.Println($a);
    }
    class TestClassB{
        function Bclass(){}
        Sys.Println($a);
    }
    class Test extends TestClassA,TestClassB{
        function Cclass(){}
        Sys.Println($a);
    }
    1+1*2;
    $a(123)->124+Test("123")*3*1>=1&&1>=0;
    1==3;

    for($i=0;$i<=10;$i++){
        Sys.Println(1)
    }


    if(1==1){
        Sys.Println(1);
    }else if(1==2){
        Sys.Println(2);
    }else{
        Sys.Println(3);
    }

    if(1==1){
        Sys.Println(1);
    }else if(1==2){
        Sys.Println(2);
    }

    if(1==1){
        Sys.Println(1);
    }else{
        Sys.Println(3);
    }
    */

/*
    1+a(123);
*/


    `
	id := memory.Malloc()
	id.Open().SetValue(initCode.Complex(Code).Ast)

	fmt.Println("-----------------------------------------------")
	fmt.Println("Debug Log")
	fmt.Println("-----------------------------------------------")
	for i := 0; i <= len((*id.Open().Read()).(map[int]ast.AstStruct)); i++ {
		fmt.Println(i, (*id.Open().Read()).(map[int]ast.AstStruct)[i])
	}
	fmt.Println("-----------------------------------------------")
	Opcode := compile.Compile((*id.Open().Read()).(map[int]ast.AstStruct))
	fmt.Println("-----------------------------------------------")
	for i := 0; i <= len(Opcode.Opcode)-1; i++ {
		fmt.Println(i, "OPCODE", Opcode.Opcode[i])
	}
	fmt.Println("-----------------------------------------------")
	fmt.Println(Opcode.FuncList)
	fmt.Println("-----------------------------------------------")
	fmt.Println(Opcode.ClassList)
	fmt.Println("-----------------------------------------------")
	vm.WspVm(Opcode)
	fmt.Println("-----------------------------------------------")

	memory.FreeAll()
}
