package initCode

import(
    "Wsp/Analysis/Lexical"
    "Wsp/Analysis/Ast"
)

type CodeStruct struct{
    Code string
    TokenList map[int]*lex.LexicalStruct
    Ast map[int]ast.AstStruct
}

func newInitCode(Code string)CodeStruct{
    return CodeStruct{Code:Code}
}

func (this *CodeStruct)LexRun(){
    this.TokenList = lex.LexRun(this.Code,1)
}

func (this *CodeStruct)AstRun(){
    this.Ast = ast.AstCode(this.TokenList)
}