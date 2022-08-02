<p align="center"><img src="./logo.png"
        alt="Logo" width="128" height="128" style="max-width: 100%;"></p>
<h1 align="center">WSP</h1>
<p align="center">一门解释型语言</p>
<p align="center">
    <a href="https://github.com/Linkangyis/Wsp_language/blob/LICENSE">
        <img src="https://img.shields.io/github/license/Ice-Hazymoon/MikuTools.svg" alt="MIT License" />
    </a>
</p>

## 安装WSP

```bash
vi /ect/profile
export WSPPATH=WSP所在目录
```
```bash
ln -s WSP所在目录/wsp /usr/bin
```
## 介绍

基于golang开发的解释型语言 使用wsp虚拟机，效率极高，当前版本 Beta3.0.0,有PHP的简单 Python的实用 Golang的效率

## 开发

```bash
wsp ./xxxx.wsp
```

## 语法
自定义函数
```php
function 变量名(参数){
    //代码块
}
```
自定义变量
```php
$xx=xx;
```
循环
```php
for(条件){
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

## 扩展开发
```golang
package main

import(
    "Wsp/WVM"  //引入核心虚拟机
    "fmt"
)

func H_Info()(map[int]string){ //定义函数列表
    info := make(map[int]string)
    info[0] = "Input"
    return info
}

func Input(a string)(string){  //定义扩展函数
    str_arr,_:=vm.Parameter_processing(a)
    var text string
	fmt.Printf(str_arr[0])
	fmt.Scanln(&text)
    return text
}
//扩展编译指令
//go build -buildmode=plugin -o scanlns.so Scanlns.go
```
## License

[MIT](https://github.com/Linkangyis/Wsp_language/blob/LICENSE)
