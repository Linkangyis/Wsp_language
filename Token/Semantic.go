package token

import(
  "../Echo"
  "../Types"
  "../Maps"
  "fmt"
  "os"
)
func Wsp_Semantic(Gra_Comp map[int][4]string)(map[int][4]string){
    for i:=0;i<=len(Gra_Comp)-1;i++{
        if Gra_Comp[i][0]==types.Strings(6){
            Gra_Comp = maps.DEL_Map(i,Gra_Comp)
            i--
        }
    }
    KUO_X:=0
    KUO_Z:=0
    KUO_D:=0
    for i:=0;i<=len(Gra_Comp)-1;i++{
        if Gra_Comp[i][0]==types.Strings(101) && Gra_Comp[i-1][0]!=types.Strings(81){
            KUO_X++
        }
        if Gra_Comp[i][0]==types.Strings(121) && Gra_Comp[i-1][0]!=types.Strings(81){
            KUO_Z++
        }
        if Gra_Comp[i][0]==types.Strings(111) && Gra_Comp[i-1][0]!=types.Strings(81){
            KUO_D++
        }
    }
    
    for i:=0;i<=len(Gra_Comp)-1;i++{
        if Gra_Comp[i][0]==types.Strings(102)&& Gra_Comp[i-1][0]!=types.Strings(81){
            KUO_X--
        }
        if Gra_Comp[i][0]==types.Strings(122)&& Gra_Comp[i-1][0]!=types.Strings(81){
            KUO_Z--
        }
        if Gra_Comp[i][0]==types.Strings(112)&& Gra_Comp[i-1][0]!=types.Strings(81){
            KUO_D--
        }
    }
    if (KUO_X!=0 || KUO_Z!=0 || KUO_D!=0){
        fmt.Println("在语义分析时出现错误，请检查代码是否多写漏写 ( ) [ ] { } 等定界符\n TOKENS调试界面:")
        echo.Arr_Echo(Gra_Comp)
        os.Exit(0);
    }
    return Gra_Comp
}