package compile

import(
    "os"
    "fmt"
    "Wsp/Module/Const"
)

func Error(err string,line int,errs string){
    fmt.Println("语法错误:\n"+"在第",line,"行时"+err+errs,"\n目录：",consts.WspConst.WspFile)
    os.Exit(0)
}

func Errors(err string,line int){
    fmt.Println("语法错误:\n  "+err,"\n  在第",line,"行","\n  目录：",consts.WspConst.WspFile)
    os.Exit(0)
}


func Check(Opcode map[int]map[int]Body_Struct_Run){
    for i:=0;i<=len(Opcode)-1;i++{
        for z:=0;z<=len(Opcode[i])-1;z++{
            if Opcode[i][z].Type==205&&z!=len(Opcode[i])-1{
                if Opcode[i][z+1].Type!=212 && Opcode[i][z+1].Type!=90&&(Opcode[i][z+1].Type!=91||(Opcode[i][z+1].Type==91&&Opcode[i][z+2].Type==98))&& Opcode[i][z+1].Type!=92&& Opcode[i][z+1].Type!=93&& Opcode[i][z+1].Type!=94{
                    Error("结尾缺少一个",Opcode[i][z].Line,";")
                }
            }
            if Opcode[i][z].Type==202 || Opcode[i][z].Type==203 || Opcode[i][z].Type==204{
                if len(Opcode[i][z].Abrk)<2 && Opcode[i][z].Type!=203{
                    Error("缺少需要条件",Opcode[i][z].Line,"() or {}")
                }else if len(Opcode[i][z].Abrk)<1 && Opcode[i][z].Type==203{
                    Error("缺少需要条件",Opcode[i][z].Line,"{}")
                }
                
                if Opcode[i][z].Abrk[0].Text=="" && Opcode[i][z].Abrk[0].Type==1{
                    Error("()判断条件不能为",Opcode[i][z].Line,"空")
                }
            }
        }
    }
}