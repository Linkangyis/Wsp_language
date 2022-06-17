# Wsp_language
【简介】基于golang开发的解释型语言 使用wsp虚拟机，效率极高，当前版本 V1.0.0<br>
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


[introduction] the interpretive language developed based on golang uses the WSP virtual machine, which is highly efficient. The current version of V1.0.0<br>
[advantages] high efficiency, low memory, high performance, 100 cycles, 4~5ms<br>
[support system] currently only supports Linux UINX<br>
[usage]<br>
&nbsp;&nbsp;&nbsp;Take CentOS as an example:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[step 1]:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Add the following environment to the /ect/profile file<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[export WSPPATH=WSP directory]<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[step 2]:<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Create a new soft link file in /usr/bin directory and point to [WSP compiled location /wsp]<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Execute on the command line<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;WSP file location<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;You can view the execution results<br>
