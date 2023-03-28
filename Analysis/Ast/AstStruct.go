package ast

type AstStruct struct{
    Type int
    UName string
    Name string
    Line int
    ValueList map[int]ValueStruct
    ClassExtends map[int]string
}

type ValueStruct struct{
    Type int
    Value string
}