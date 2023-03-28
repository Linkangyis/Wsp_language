package compile

import (
	"Wsp/Analysis/Ast"
	"Wsp/Analysis/Lexical"
	"fmt"
	"strings"
)

var TmpRunCode map[int]int
var SeparateConfigMap map[string]bool

func init() {
	TmpRunCode = map[int]int{
		500: 500,
		501: 501,
		502: 502,
		503: 503,
		508: 508,
		600: 600,
		601: 601,
		602: 602,
		603: 603,
		604: 604,
		605: 605,
		606: 606,
		607: 607,
	}
	SeparateConfigMap = make(map[string]bool)
	SeparateConfigMap[","] = true
	SeparateConfigMap[" "] = true
	SeparateConfigMap["\n"] = true
	SeparateConfigMap["\t"] = true
	SeparateConfigMap["\r"] = true
	SeparateConfigMap[";"] = true
	SeparateConfigMap["+"] = true
	SeparateConfigMap["-"] = true
	SeparateConfigMap["*"] = true
	SeparateConfigMap["/"] = true
	SeparateConfigMap["%"] = true
	SeparateConfigMap[">"] = true
	SeparateConfigMap["<"] = true
	SeparateConfigMap["="] = true
	SeparateConfigMap["!"] = true
	SeparateConfigMap["&"] = true
	SeparateConfigMap["|"] = true
}

func ReadType(num int) int {
	switch num {
	case 606, 607:
		return 0
	case 600, 601, 602, 603, 604, 605:
		return 1
	case 500, 501:
		return 2
	case 502, 503, 508:
		return 3
	}
	return 4
}

func IfRunType(a int, b int) bool {
	if ReadType(a) < ReadType(b) {
		return true
	}
	return false
}
func makeValueListStructMap(AstCode ast.AstStruct, line int) map[int]ValueListStruct {
	ValueList := make(map[int]ValueListStruct)
	for ValueListPorint := 0; ValueListPorint <= len(AstCode.ValueList)-1; ValueListPorint++ {
		TmpOpcode := Compile(ast.AstCode(lex.LexRun(AstCode.ValueList[ValueListPorint].Value, line))).Opcode
		ValueList[len(ValueList)] = ValueListStruct{
			Type:  AstCode.ValueList[ValueListPorint].Type,
			Value: TmpOpcode,
		}
	}
	return ValueList
}

func codeForOpcode(Code string, line int) map[int]map[int]OpRun {
	return Compile(ast.AstCode(lex.LexRun(Code, line))).Opcode
}

func funcVarErrorIf(VarName string, line int) string {
	for i := 0; i <= len(VarName)-1; i++ {
		if SeparateConfigMap[string(VarName[i])] {
			ErrorPrint(line, "在声明的变量中 出现不可预料的内容")
		}
	}
	return VarName
}

func splitFuncVar(funcVar string, line int) map[int]string {
	Res := make(map[int]string)
	tmp := strings.Split(funcVar, ",")
	for i := 0; i <= len(tmp)-1; i++ {
		Res[len(Res)] = funcVarErrorIf(strings.TrimSpace(tmp[i]), line)
	}
	return Res
}

func Trim(Opcode map[int]map[int]OpRun) map[int]map[int]OpRun {
	Res := make(map[int]map[int]OpRun)
	for i := 0; i <= len(Opcode)-1; i++ {
		if len(Opcode[i]) > 0 {
			Res[len(Res)] = Opcode[i]
		}
	}
	return Res
}

var funcListMap = make(map[int]FuncStruct)
var ClassList = make(map[string]ClassStruct)
var classComileLock = false
var classUserFunc map[string]FuncStruct

func Compile(Ast map[int]ast.AstStruct) CompileStruct {
	OpcodeMap := make(map[int]map[int]OpRun)
	Opcode := make(map[int]OpRun)
	OpCodeBlockId := 0
	for i := 0; i <= len(Ast)-1; i++ {
		switch Ast[i].Type {
		case 0:
			if len(Ast[i].ValueList) != 0 && (Ast[i].ValueList[0].Type == 0) {
				ValueList := makeValueListStructMap(Ast[i], Ast[i].Line)
				Opcode[len(Opcode)] = OpRun{
					Type:      200,
					Name:      "CALL",
					Text:      Ast[i].Name,
					Line:      Ast[i].Line,
					ValueList: ValueList,
				}
			} else if len(Ast[i].ValueList) == 0 && Ast[i].UName == "TEXT" {
				Opcode[len(Opcode)] = OpRun{
					Type:      0,
					Name:      Ast[i].UName,
					Text:      Ast[i].Name,
					Line:      Ast[i].Line,
					ValueList: make(map[int]ValueListStruct),
				}
			}
		case 203:
			ValueList := makeValueListStructMap(Ast[i], Ast[i].Line)

			Opcode[len(Opcode)] = OpRun{
				Type:      201,
				Name:      Ast[i].UName,
				Text:      Ast[i].Name,
				Line:      Ast[i].Line,
				ValueList: ValueList,
			}
			OpcodeMap[OpCodeBlockId] = Opcode
			OpCodeBlockId++
			Opcode = make(map[int]OpRun)

		case 412:
			if Ast[i+1].Type == 507 {
				ValueList := makeValueListStructMap(Ast[i], Ast[i].Line)
				Opcode[len(Opcode)] = OpRun{
					Type:      301,
					Name:      Ast[i].UName,
					Text:      Ast[i].Name,
					Line:      Ast[i].Line,
					ValueList: ValueList,
				}
				i++
			} else {
				//仅供测试
				ValueList := makeValueListStructMap(Ast[i], Ast[i].Line)
				Opcode[len(Opcode)] = OpRun{
					Type:      300,
					Name:      Ast[i].UName,
					Text:      Ast[i].Name,
					Line:      Ast[i].Line,
					ValueList: ValueList,
				}
			}

		case 413:
			ValueList := makeValueListStructMap(Ast[i], Ast[i].Line)
			Opcode[len(Opcode)] = OpRun{
				Type:      202,
				Name:      Ast[i].UName,
				Text:      Ast[i].Name,
				Line:      Ast[i].Line,
				ValueList: ValueList,
			}

		case 200:
			if !classComileLock {
				ThisFuncOpocde := codeForOpcode(Ast[i].ValueList[1].Value, Ast[i].Line)
				funcListMap[len(funcListMap)] = FuncStruct{
					Name:     Ast[i].Name,
					ValueVar: splitFuncVar(Ast[i].ValueList[0].Value, Ast[i].Line),
					Opcode:   ThisFuncOpocde,
				}
				Opcode[len(Opcode)] = OpRun{
					Type:     204,
					Name:     "INIT_FUNC",
					Text:     Ast[i].Name,
					Line:     Ast[i].Line,
					Register: len(funcListMap) - 1,
				}
				if SeparateConfigMap[Ast[i+1].Name] && Ast[i+1].Type != 408 && Ast[i+1].Name != "\n" {

				} else {
					OpcodeMap[OpCodeBlockId] = Opcode
					OpCodeBlockId++
					Opcode = make(map[int]OpRun)
				}
			} else {
				ThisFuncOpocde := codeForOpcode(Ast[i].ValueList[1].Value, Ast[i].Line)
				classUserFunc[Ast[i].Name] = FuncStruct{
					Name:     Ast[i].Name,
					ValueVar: splitFuncVar(Ast[i].ValueList[0].Value, Ast[i].Line),
					Opcode:   ThisFuncOpocde,
				}
			}
		case 201:
			OpcodeTmp := make(map[int]map[int]OpRun)
			classUserFunc = make(map[string]FuncStruct)
			if len(Ast[i].ClassExtends) > 0 {
				for ci := 0; ci <= len(Ast[i].ClassExtends)-1; ci++ {
					ClassTmpName := Ast[i].ClassExtends[ci]
					rClass := ClassList[ClassTmpName]
					for rop := 0; rop <= len(rClass.Opcode)-1; rop++ {
						OpcodeTmp[len(OpcodeTmp)] = rClass.Opcode[rop]
					}
					for ck, cv := range rClass.ClassUserFunc {
						classUserFunc[ck] = cv
					}
				}
			}

			classComileLock = true
			tmpOpcode := codeForOpcode(Ast[i].ValueList[0].Value, Ast[i].Line)
			for oi := 0; oi <= len(tmpOpcode)-1; oi++ {
				OpcodeTmp[len(OpcodeTmp)] = tmpOpcode[oi]
			}
			classComileLock = false

			ClassList[Ast[i].Name] = ClassStruct{
				ClassName:     Ast[i].Name,
				Opcode:        OpcodeTmp,
				ClassUserFunc: classUserFunc,
			}
		case 205, 206, 207:
			ValueList := makeValueListStructMap(Ast[i], Ast[i].Line)

			Opcode[len(Opcode)] = OpRun{
				Type:      Ast[i].Type,
				Name:      Ast[i].UName,
				Text:      Ast[i].Name,
				Line:      Ast[i].Line,
				ValueList: ValueList,
			}
			if Ast[i].Type == 207 {
				if Ast[i+1].Type == 206 {
					ErrorPrint(Ast[i].Line, "else已经结束语句，不允许再后面接else if")
				} else {
					OpcodeMap[OpCodeBlockId] = Opcode
					OpCodeBlockId++
					Opcode = make(map[int]OpRun)
				}
			} else {
				if Ast[i+1].Type != 206 && Ast[i+1].Type != 207 {
					OpcodeMap[OpCodeBlockId] = Opcode
					OpCodeBlockId++
					Opcode = make(map[int]OpRun)
				}
			}
		case 408:
			if _, ok := OpcodeMap[OpCodeBlockId]; ok {
				OpCodeBlockId++
			}
			Opcode = make(map[int]OpRun)
		case 302:
			ValueList := makeValueListStructMap(Ast[i], Ast[i].Line)
			for _, ValueList_v := range ValueList {
				ErrorIf := false
				if ValueList_v.Type != 1 {
					ErrorIf = true
				}
				if ErrorIf {
					ErrorPrint(Ast[i].Line, "自增只允许在变量后面不允许在 函数/类后面")
				}
			}
			Opcode[len(Opcode)] = OpRun{
				Type:      302,
				Name:      Ast[i].UName,
				Text:      Ast[i].Name,
				Line:      Ast[i].Line,
				ValueList: ValueList,
			}
		case 303:
			ValueList := makeValueListStructMap(Ast[i], Ast[i].Line)
			for _, ValueList_v := range ValueList {
				ErrorIf := false
				if ValueList_v.Type != 1 {
					ErrorIf = true
				}
				if ErrorIf {
					ErrorPrint(Ast[i].Line, "自减只允许在变量后面不允许在 函数/类后面")
				}
			}
			Opcode[len(Opcode)] = OpRun{
				Type:      303,
				Name:      Ast[i].UName,
				Text:      Ast[i].Name,
				Line:      Ast[i].Line,
				ValueList: ValueList,
			}
		case 414:
			Opcode[len(Opcode)] = OpRun{
				Type: 400,
				Name: Ast[i].UName,
				Text: Ast[i].Name,
				Line: Ast[i].Line,
			}
		default:
			if _, ok := TmpRunCode[Ast[i].Type]; ok {
				i--
				if i == 0 {
					ErrorPrint(Ast[i].Line, "出现意外的表达式")
				}
				var Register = make(map[int]Erun)
				var TmpA = make(map[int]int)

				for z := i; z <= len(Ast)-1; z++ {
					i++
					if Ast[z].Type == 408 || Ast[z].Type == 414 {
						i--
						break
					}
					if _, ok := TmpRunCode[Ast[z].Type]; !ok {
						TmpMap := make(map[int]ast.AstStruct)
						TmpMap[0] = Ast[z]
						Register[len(Register)] = Erun{
							Type:   false,
							OpErun: Compile(TmpMap).Opcode,
						}
					} else {
						if len(TmpA) > 0 {
							if IfRunType(TmpA[len(TmpA)-1], Ast[z].Type) {
								TmpA[len(TmpA)] = Ast[z].Type
							} else {
								for TmpALen := len(TmpA) - 1; TmpALen >= 0; TmpALen-- {
									if !IfRunType(TmpA[TmpALen], Ast[z].Type) {
										Register[len(Register)] = Erun{
											Type:    true,
											RunType: TmpA[TmpALen],
										}
										delete(TmpA, TmpALen)
										if len(TmpA) == 0 {
											TmpA[len(TmpA)] = Ast[z].Type
											break
										}
									} else {
										TmpA[len(TmpA)] = Ast[z].Type
										break
									}
								}
							}
						} else {
							TmpA[len(TmpA)] = Ast[z].Type
						}
					}

					fmt.Println(TmpA, z, Ast[z])

				}
				i--
				if len(TmpA) != 0 {
					for y := len(TmpA) - 1; y >= 0; y-- {
						Register[len(Register)] = Erun{
							Type:    true,
							RunType: TmpA[y],
						}
					}
					TmpA = make(map[int]int)
				}

				//fmt.Println(Opcode)
				Opcode[len(Opcode)-1] = OpRun{
					Type:     203,
					Name:     "ERUN",
					Text:     "ERUN",
					Line:     Ast[i].Line,
					Register: Register,
				}
				//fmt.Println(Register,1)
			}
		}

		OpcodeMap[OpCodeBlockId] = Opcode

	}
	return CompileStruct{
		Opcode:    Trim(OpcodeMap),
		FuncList:  funcListMap,
		ClassList: ClassList,
	}
}
