# Wsp_language
【简介】基于golang开发的解释型语言 使用wsp虚拟机，效率极高，当前版本 BETA 4<br>
【优点】高效率 低内存 高性能 百次循环4~5ms<br>
【支持系统】目前只支持Linux Uinx<br>
【使用方法】<br>
&nbsp;&nbsp;&nbsp;以centos为例：<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;【第一步】：<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;在 /ect/profile文件加入以下环境<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;【export WSPPATH=WSP所在目录】<br><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;【第二步】：<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;在/usr/bin 目录新建一个软链接文件，指向 【WSP编译后的位置/wsp】<br><br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;在命令行执行<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; wsp 文件位置 <br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;即可查看执行结果<br>
