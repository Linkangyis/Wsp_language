<p align="center"><img src="./logo.png"
         alt="Logo" width="128" height="128" style="max-width: 100%;"></p>
<h1 align="center">WSP</h1>
<p align="center">一门解释型语言(此分支为Windows插件技术预览)</p>
<p align="center">
    <a href="https://github.com/Linkangyis/Wsp_language/blob/LICENSE">
        <img src="https://img.shields.io/github/license/Ice-Hazymoon/MikuTools.svg" alt="MIT License" />
    </a>
</p>

## 已知问题
1. 如果插件调用Vm内部函数，将会出现空Vm现象，该问题由Windows特性所导致，（v5.0.0） 将来会选择虚拟vm包(利用rpc进行实现)进行修补
2. 此版本基于v4.6.2-beta.1开发 因其极其不稳定，所以不推荐进行长期使用，短期测试随便
3. 因Linux遗留问题 GC 快速变量等功能将完全不会受支持，请勿手贱开启

## 安装WSP
//此版本不支持安装 使用方法如下
```shell
go run main.go xxx.wsp
./main xxx.wsp
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
