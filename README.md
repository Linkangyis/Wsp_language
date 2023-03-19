<p align="center"><img src="./logo.png"
         alt="Logo" width="128" height="128" style="max-width: 100%;"></p>
<h1 align="center">WSP</h1>
<p align="center">一门解释型语言(5.0版本正在开发)</p>
<p align="center">
    <a href="https://github.com/Linkangyis/Wsp_language/blob/LICENSE">
        <img src="https://img.shields.io/github/license/Ice-Hazymoon/MikuTools.svg" alt="MIT License" />
    </a>
</p>

## Wsp5.0.0进度（已经在新的分支公开开发阶段）
1. Lex       完成  （进度100%）
2. Ast       完成  （进度100%）
3. Compile   进行中（进度50%）
4. VarModule 完成  （进度100%）
5. WspVm     未开始（进度0%）
6. GcModule  未开始（进度0%）
7. 其他      未开始（进度0%）

## 安装WSP

```bash
vi /ect/profile
export WSPPATH=WSP所在目录
```
```bash
ln -s WSP所在目录/wsp /usr/bin
```
## 介绍

基于golang开发的解释型语言 使用wsp虚拟机，效率极高,有PHP的简单 Python的实用 Golang的效率

## 开发

```bash
wsp [命令行形式]
wsp ./xxxx.wsp
```

## 语法
自定义函数
```php
function 函数名(参数){
    //代码块
}
```
自定义变量
```php
$xx=xx;
```
循环
```php
xx;xx;xx形式
for(条件){
    //代码块
}
```
```php
while形式
for(条件){
    //代码块
}
```
```php
do_while形式
for{
    //代码块
}(条件)
```
```php
死循环
for{
    //代码块
}
```
判断
```php
if(条件){
    //代码块
}else if(条件){
    //代码块
}else{
    //代码块
}
```
Switch语句
```php
$a = "3";
switch($a){
    case 1:
        print(1);
    case 2:
        print(2);
    case 3:
        print(3);
    default:
        print(4);
}
```
wgo协程
```php
wgo func();//堆栈均共享(可选)
```
class类
```php
class Test{
    function PrintClass($a){
        print($a);
    }
    function _init_($c){
         $this->Var=$c;
         print($this->Var);
    }
    function TestPrint(){
        $this->PrintClass($this->Var);
    }
    $Var = 110;
}
$TestClass = new Test(001);
$TestClassB = new Test(002);
$TestClass->Var = "测试";
$TestClassB->Var = "测试2";
$TestClass->TestPrint();
$TestClassB->TestPrint();
```
Class继承
```php
class Test{
    function TestPrinta(){
        print($this->Vara);
    }
    $Var = 0;
    $Vara = 0;
}
class TestB{
    function TestPrintb(){
        print($this->Var);
    }
    $Var = 1;
}
class TestC extends Test,TestB{
    function _init_(){
        print($this->Var)
    }
    function Test(){
        print("extends");
        print($this->Var);
    }
}
$a = new TestC();
$a->Var=10086;
$a->TestPrinta();
$a->Test();
```
面向容器编程 （暂定）
```php
CurEnv->this->5555->31; //公开的容器 ID 31  //端口5555
        $TestArray[0]=10086;
        $All = 0;
CurEnv->this->0->Main;  //不公开的容器 主容器  
        $Pinter = Raflect.OpenCont(31);
        $All = Raflect.ReadCont($Pinter,"All")+1;
        $TestArray = Raflect.ReadCont($Pinter,"TestArray");
        Sys.Println($All);
        Sys.Println($TestArray[0]);
```
## 扩展开发 （未来会进行修改 [待定]）
可以查看 Ext/Test 测试插件作为实例
```golang
package main

import(
  "fmt"
)
func Func_Info()(map[int]string){  //系统核心
    info := make(map[int]string)  //函数列表
    info[0] = "Testb"
    info[1] = "Tests"
    return info
}
func Package_Info()(string){  //系统核心
    info := "Test"   //包名设置
    return info
}


func Testb(Value map[int]string)(string){    //Testb扩展函数 wsp调用 Test.Testb()
    fmt.Println("b")
    return "TRUE"
}

func Tests(Value map[int]string)(string){    //Tests扩展函数 wsp调用 Test.Tests()
    fmt.Println("s")
    return "TRUE"
}
//扩展编译指令
//go build -buildmode=plugin -o test.so Test.go
```
## License

[MIT](https://github.com/Linkangyis/Wsp_language/blob/LICENSE)
