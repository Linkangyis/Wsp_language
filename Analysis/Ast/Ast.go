package ast

import (
	"Wsp/Analysis/Lexical"
)

var SeparateConfigMap map[string]bool

func init() {
	SeparateConfigMap = make(map[string]bool)
	SeparateConfigMap[","] = true
	SeparateConfigMap[" "] = true
	SeparateConfigMap["\n"] = true
	SeparateConfigMap["\t"] = true
	SeparateConfigMap["\r"] = true
	SeparateConfigMap["$"] = true
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

func AstCode(LexCode map[int]*lex.LexicalStruct) map[int]AstStruct {
	ResLine := 0
	Res := make(map[int]AstStruct)
	var VarAstLock bool = false
	for i := 0; i <= len(LexCode)-1; i++ {
		if i == len(LexCode)-2 {
			break
		}
		This := *LexCode[i]
		ThisNext := *LexCode[i+1]
		ThisNextNext := *LexCode[i+2]
		ResStruct := AstStruct{
			Type:  This.Type,
			UName: This.Name,
			Name:  This.Text,
			Line:  This.Line,
		}

		if i > 0 && (*LexCode[i-1]).Type == 201 {
			if (*LexCode[i+1]).Type == 204 {
				ResLine--
				ResStruct = AstStruct{
					Type:  201,
					UName: "CLASS",
					Name:  This.Text,
					Line:  This.Line,
				}
				TmpClassExtends := make(map[int]string)
				for z := i + 1; z <= len(LexCode)-1; z++ {
					This := *LexCode[z]
					if This.Type == 406 {
						break
					} else if This.Type != 414 {
						TmpClassExtends[len(TmpClassExtends)] = This.Text
					}
					i++
				}
				ResStruct.ClassExtends = TmpClassExtends
				This = *LexCode[i]
				ThisNext = *LexCode[i+1]
				ThisNextNext = *LexCode[i+2]
			}
		}

		if ThisNext.Type == 402 || ThisNext.Type == 404 || ThisNext.Type == 406 || (ThisNext.Type == 501 && ThisNextNext.Type == 504) {
			if lex.JudgeSeparate(This.Text) {
				Res[ResLine] = AstStruct{
					Type:  This.Type,
					UName: This.Name,
					Name:  This.Text,
					Line:  This.Line,
				}
				ResLine++
				ResStruct = AstStruct{
					Type:  413,
					UName: "EVAL",
					Name:  "EVAL",
					Line:  This.Line,
				}
			}
			if VarAstLock {
				VarAstLock = false
				ResStruct = AstStruct{
					Type:  412,
					UName: "VAR",
					Name:  This.Text,
					Line:  This.Line,
				}
			}
			if i > 0 && (*LexCode[i-1]).Type == 201 {
				if (*LexCode[i+1]).Type != 204 {
					ResLine--
					ResStruct = AstStruct{
						Type:  201,
						UName: "CLASS",
						Name:  This.Text,
						Line:  This.Line,
					}
				}
			}

			if i > 0 && (*LexCode[i-1]).Type == 207 {
				if This.Type == 205 {
					ResLine--
					ResStruct = AstStruct{
						Type:  206,
						UName: "ELSEIF",
						Name:  "else if",
						Line:  This.Line,
					}
				}
			}

			if i > 0 && (*LexCode[i-1]).Type == 200 {
				ResLine--
				ResStruct = AstStruct{
					Type:  200,
					UName: "NEW_FUNC",
					Name:  This.Text,
					Line:  This.Line,
				}
			}

			ValueList := make(map[int]ValueStruct)
			for z := i + 1; z <= len(LexCode)-1; z++ {
				This := *LexCode[z]
				ThisNext := *LexCode[z+1]

				if This.Type == 402 {
					ValueList[len(ValueList)] = ValueStruct{
						Type:  0,
						Value: (*LexCode[z+1]).Text,
					}
					z += 2
				} else if This.Type == 404 {
					ValueList[len(ValueList)] = ValueStruct{
						Type:  1,
						Value: (*LexCode[z+1]).Text,
					}
					z += 2
				} else if This.Type == 406 {
					ValueList[len(ValueList)] = ValueStruct{
						Type:  2,
						Value: (*LexCode[z+1]).Text,
					}
					z += 2
				} else if This.Type == 501 && ThisNext.Type == 504 {
					ValueList[len(ValueList)] = ValueStruct{
						Type:  3,
						Value: (*LexCode[z+2]).Text,
					}
					z += 2

				} else {
					break
				}
				i += 3
			}
			ResStruct.ValueList = ValueList
			Res[ResLine] = ResStruct

		} else if This.Type == 409 || This.Type == 410 || This.Type == 411 {
			ResStruct := AstStruct{
				Type:  ThisNext.Type,
				UName: ThisNext.Name,
				Name:  "\"" + ThisNext.Text + "\"",
				Line:  ThisNext.Line,
			}
			Res[ResLine] = ResStruct
			i += 2
		} else if This.Type == 412 {
			VarAstLock = true
			continue
		} else if VarAstLock {
			VarAstLock = false
			ResStruct := AstStruct{
				Type:  412,
				UName: "VAR",
				Name:  This.Text,
				Line:  This.Line,
			}
			Res[ResLine] = ResStruct
		} else if This.Type == 504 {
			if ThisNext.Type == 507 {
				i++
				ResStruct := AstStruct{
					Type:  601,
					UName: "DAYUDENGYU",
					Name:  "DAYUDENGYU",
					Line:  ThisNext.Line,
				}
				Res[ResLine] = ResStruct
			} else {
				ResStruct := AstStruct{
					Type:  600,
					UName: "DAYU",
					Name:  "DAYU",
					Line:  ThisNext.Line,
				}
				Res[ResLine] = ResStruct
			}
		} else if This.Type == 507 && ThisNext.Type == 507 {
			i++
			ResStruct := AstStruct{
				Type:  602,
				UName: "DENGYU",
				Name:  "DENGYU",
				Line:  ThisNext.Line,
			}
			Res[ResLine] = ResStruct
		} else if This.Type == 505 {
			if ThisNext.Type == 507 {
				i++
				ResStruct := AstStruct{
					Type:  603,
					UName: "XIAOYUENGYU",
					Name:  "XIAOYUENGYU",
					Line:  ThisNext.Line,
				}
				Res[ResLine] = ResStruct
			} else {
				ResStruct := AstStruct{
					Type:  604,
					UName: "XIAOY",
					Name:  "XIAOY",
					Line:  ThisNext.Line,
				}
				Res[ResLine] = ResStruct
			}
		} else if This.Type == 506 && ThisNext.Type == 507 {
			i++
			ResStruct := AstStruct{
				Type:  605,
				UName: "BUDENGYU",
				Name:  "BUDENGYU",
				Line:  ThisNext.Line,
			}
			Res[ResLine] = ResStruct
		} else if This.Type == 509 && ThisNext.Type == 509 {
			i++
			ResStruct := AstStruct{
				Type:  606,
				UName: "LUOJIYU",
				Name:  "LUOJIYU",
				Line:  ThisNext.Line,
			}
			Res[ResLine] = ResStruct
		} else if This.Type == 510 && ThisNext.Type == 510 {
			i++
			ResStruct := AstStruct{
				Type:  607,
				UName: "LUOJIFEI",
				Name:  "LUOJIFEI",
				Line:  ThisNext.Line,
			}
			Res[ResLine] = ResStruct
		} else if i != 0 && (SeparateConfigMap[(*LexCode[i-1]).Text] && This.Type == 501 && ThisNext.Type == 0) {
			i++
			ResStruct := AstStruct{
				Type:  0,
				UName: "TEXT",
				Name:  "-" + ThisNext.Text,
				Line:  ThisNext.Line,
			}
			//ResLine--
			Res[ResLine] = ResStruct
		} else if i != 0 && (SeparateConfigMap[(*LexCode[i-1]).Text] && This.Type == 500 && ThisNext.Type == 0) {
			i++
			ResStruct := AstStruct{
				Type:  0,
				UName: "TEXT",
				Name:  "+" + ThisNext.Text,
				Line:  ThisNext.Line,
			}
			//ResLine--
			Res[ResLine] = ResStruct
		} else if ResLine > 1 && Res[ResLine-1].Type == 412 && (This.Type == 500 && ThisNext.Type == 500) {
			i++
			if i != 0 && (SeparateConfigMap[(*LexCode[i-1]).Text] && (*LexCode[i]).Type == 500 && (*LexCode[i+1]).Type == 0) {
				ErrorPrint((*LexCode[i]).Line, "未定义行为")
			}
			TmpStruct := Res[ResLine-1]
			TmpStruct.Type = 302
			ResLine--
			Res[ResLine] = TmpStruct
		} else if ResLine > 1 && Res[ResLine-1].Type == 412 && (This.Type == 501 && ThisNext.Type == 501) {
			i++
			if i != 0 && (SeparateConfigMap[(*LexCode[i-1]).Text] && (*LexCode[i]).Type == 501 && (*LexCode[i+1]).Type == 0) {
				ErrorPrint((*LexCode[i]).Line, "未定义行为")
			}
			TmpStruct := Res[ResLine-1]
			TmpStruct.Type = 303
			ResLine--
			Res[ResLine] = TmpStruct
		} else if This.Type == 413 {
			ResStruct := AstStruct{
				Type:  414,
				UName: This.Name,
				Name:  This.Text,
				Line:  This.Line,
			}
			Res[ResLine] = ResStruct
		} else {
			ResStruct := AstStruct{
				Type:  This.Type,
				UName: This.Name,
				Name:  This.Text,
				Line:  This.Line,
			}
			Res[ResLine] = ResStruct
		}
		ResLine++
	}
	return Res
}
