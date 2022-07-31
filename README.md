<p align="center"><img src="./static/icon.png"
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

基于golang开发的解释型语言 使用wsp虚拟机，效率极高，当前版本 V2.0.0,有PHP的简单 Python的实用 Golang的效率

## 开发

```bash
wsp ./xxxx.wsp
```

## 语法
自定义函数
```bash
function(参数){
    //代码块
}
```
自定义变量
```bash
$xx=xx;
```
循环
```bash
for(条件){
    //代码块
}
```

判断
```bash
if(条件){
    //代码块
}else if(条件){
    //代码块
}else{
    //代码块
}
```


## License

[MIT](https://github.com/Ice-Hazymoon/MikuTools/blob/master/LICENSE)
