package initCode

func Complex(Code string)CodeStruct{
    newRes := newInitCode(Code)
    newRes.LexRun()
    newRes.AstRun()
    return newRes
}