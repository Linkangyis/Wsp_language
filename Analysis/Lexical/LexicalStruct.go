package lex

var TokenConfigMap = make(map[string]TokenConfig)
var SeparateConfigMap = make(map[string]bool)
var LockCompete = TokenLexLock{}

type TokenLexLock struct{
    Num int
    AddType string
    DoneType string
}

type TmpLexical struct{
    Text string
    Type bool
    Line int
}

type LexicalStruct struct{
    Type int
    Text string
    Name string
    Line int
}


type TokenConfig struct{
    Name string
    Type int
    ConfigType bool
    Replace string
    Hide bool
}


func newLexical()*LexicalStruct{
    return &LexicalStruct{}
}

func newLexMap()map[int]*LexicalStruct {
    return make(map[int]*LexicalStruct)
}


func (this *TokenLexLock)SetType(AddType string,DoneType string){
    this.AddType = AddType
    this.DoneType = DoneType
}

func (this *TokenLexLock)Add(Type string)bool{
    if Type==this.AddType{
        this.Num++
        return true
    }
    return false
}

func (this *TokenLexLock)Done(Type string)bool{
    if Type==this.DoneType{
        this.Num--
        return true
    }
    return false
}

func (this *TokenLexLock)Wait()bool{
    if this.Num==0{
        return true
    }
    return false
}

func (this *LexicalStruct)SetType(Type int){
    this.Type = Type
}

func (this *LexicalStruct)SetText(String string){
    this.Text = String
}

func (this *LexicalStruct)SetName(String string){
    this.Name = String
}
