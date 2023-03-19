package ast

type AstStruct struct{
    Type int
    UName string
    Name string
    Line int
    ValueList map[int]ValueStruct
}

type ValueStruct struct{
    Type int
    Value string
}