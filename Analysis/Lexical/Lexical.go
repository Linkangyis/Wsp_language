package lex

func LexRun(CodeText string, initline int) map[int]*LexicalStruct {
	CodeText = ";;;" + CodeText + ";;;"
	Code := []rune(CodeText)
	LexInit()
	var TmpString string
	var TmpStringMap = make(map[int]TmpLexical)
	var Line int = initline
	var TextLock int = 0
	var PointerLock bool = false
	var TextLockType string

	var LockButs bool = true
	var LockButsType int

	var ZhuShi bool = false
	var ZhuShiType int

	for pointer := 0; pointer <= len(Code)-1; pointer++ {
		CodeThisPointer := string(Code[pointer])
		ShangYi := ""
		if pointer != 0 {
			ShangYi = string(Code[pointer-1])
		}
		if CodeThisPointer == "\n" {
			Line++
		}

		if TextLock == 0 && !ZhuShi {
			if CodeThisPointer == "/" && string(Code[pointer+1]) == "/" {
				ZhuShi = true
				ZhuShiType = 0
			} else if CodeThisPointer == "/" && string(Code[pointer+1]) == "*" {
				ZhuShi = true
				ZhuShiType = 1
				pointer += 1
			}
		}

		if ZhuShi {
			CodeThisPointer = string(Code[pointer])
			if ZhuShiType == 0 && (CodeThisPointer == "\n") {
				ZhuShi = false
			} else if ZhuShiType == 1 && (CodeThisPointer == "*" && string(Code[pointer+1]) == "/") {
				ZhuShi = false
				pointer += 1
				continue
			} else {
				continue
			}
		}

		if ShangYi != "\\" {
			if CodeThisPointer == "\"" {
				if LockButsType == 0 {
					LockButsType = 1
				}
				if !LockButs && LockButsType == 1 {
					LockButs = true
				} else if LockButsType == 1 {
					LockButsType = 0
					LockButs = false
				}
			} else if CodeThisPointer == "'" {
				if LockButsType == 0 {
					LockButsType = 2
				}
				if !LockButs && LockButsType == 2 {
					LockButs = true
				} else if LockButsType == 2 {
					LockButsType = 0
					LockButs = false
				}
			} else if CodeThisPointer == "`" {
				if LockButsType == 0 {
					LockButsType = 3
				}
				if !LockButs && LockButsType == 3 {
					LockButs = true
				} else if LockButsType == 3 {
					LockButsType = 0
					LockButs = false
				}
			}
		}

		var DerferLis bool
		var DerferLtmp TmpLexical

		if !PointerLock {
			if TextLock == 0 && (CodeThisPointer == "\"" || CodeThisPointer == "'" || CodeThisPointer == "`") {
				if CodeThisPointer == "\"" {
					TextLockType = "\""
				} else if CodeThisPointer == "'" {
					TextLockType = "'"
				} else if CodeThisPointer == "`" {
					TextLockType = "`"
				}
				TextLock = 1

				TmpStringMap[len(TmpStringMap)] = TmpLexical{
					Text: TextLockType,
					Type: true,
					Line: Line,
				}
				continue

			} else if CodeThisPointer == "\"" || CodeThisPointer == "'" || CodeThisPointer == "`" {
				if CodeThisPointer == TextLockType && ShangYi != "\\" {
					TextLock = 2
					DerferLis = true
					DerferLtmp = TmpLexical{
						Text: TextLockType,
						Type: true,
						Line: Line,
					}
				}
			}
		}
		if TextLock == 0 && !PointerLock && (CodeThisPointer == "{" || CodeThisPointer == "(" || CodeThisPointer == "[") {
			if CodeThisPointer == "{" {
				LockCompete.SetType(CodeThisPointer, "}")
			} else if CodeThisPointer == "(" {
				LockCompete.SetType(CodeThisPointer, ")")
			} else if CodeThisPointer == "[" {
				LockCompete.SetType(CodeThisPointer, "]")
			}
			PointerLock = true
		}

		var Deferis bool = false
		var DeferTmp TmpLexical
		if PointerLock && LockButs {
			if LockCompete.Add(CodeThisPointer) {
				if LockCompete.Num == 1 {
					TmpStringMap[len(TmpStringMap)] = TmpLexical{
						Text: CodeThisPointer,
						Type: true,
						Line: Line,
					}
					continue
				}
			}
			if LockCompete.Done(CodeThisPointer) {
				if LockCompete.Wait() {
					Deferis = true
					DeferTmp = TmpLexical{
						Text: CodeThisPointer,
						Type: true,
						Line: Line,
					}
				}
			}
		}

		if !LockCompete.Wait() || TextLock == 1 {
			TmpString += CodeThisPointer
			continue
		} else if LockCompete.Wait() && PointerLock {
			MapIds := len(TmpStringMap)
			PointerLock = false
			TmpStringMap[MapIds] = TmpLexical{
				Text: TmpString,
				Type: true,
				Line: Line,
			}
			TmpString = ""

			if Deferis {
				MapIds = len(TmpStringMap)
				TmpStringMap[MapIds] = DeferTmp
			}

			continue
		} else if TextLock == 2 {
			MapIds := len(TmpStringMap)
			TmpStringMap[MapIds] = TmpLexical{
				Text: TmpString,
				Type: true,
				Line: Line,
			}
			TmpString = ""
			TextLock = 0

			if DerferLis {
				MapIds = len(TmpStringMap)
				TmpStringMap[MapIds] = DerferLtmp
			}
			continue
		}

		CodeNextPointer := "<NIL>"
		if pointer != len(Code)-1 {
			CodeNextPointer = string(Code[pointer+1])
		}
		TmpString += CodeThisPointer
		if !JudgeSeparate(CodeNextPointer) && !JudgeSeparate(CodeThisPointer) {
			continue
		}
		MapId := len(TmpStringMap)
		if JudgeToken(TmpString, true) {
			TmpStringMap[MapId] = TmpLexical{
				Text: TmpString,
				Type: true,
				Line: Line,
			}
		} else {
			TmpStringMap[MapId] = TmpLexical{
				Text: TmpString,
				Type: false,
				Line: Line,
			}
		}
		TmpString = ""
	}
	if TmpString != "" {
		MapId := len(TmpStringMap)
		if JudgeToken(TmpString, true) {
			TmpStringMap[MapId] = TmpLexical{
				Text: TmpString,
				Type: true,
				Line: Line,
			}
		} else {
			TmpStringMap[MapId] = TmpLexical{
				Text: TmpString,
				Type: false,
				Line: Line,
			}
		}
	}
	//fmt.Println(TmpStringMap)
	ResMap := make(map[int]*LexicalStruct)

	for pointer := 0; pointer <= len(TmpStringMap)-1; pointer++ {
		ThisStruct := TmpStringMap[pointer]
		ResMapId := len(ResMap)
		if ThisStruct.Type {
			TokenStruct, Type := TokenConfigMap[ThisStruct.Text]
			Text := ThisStruct.Text
			if Type {
				if TokenStruct.Replace != "" {
					Text = TokenStruct.Replace
				}
				if TokenStruct.Hide == false {
					ResMap[ResMapId] = &LexicalStruct{
						Type: TokenStruct.Type,
						Text: Text,
						Name: TokenStruct.Name,
						Line: ThisStruct.Line,
					}
				}
			} else {
				ResMap[ResMapId] = &LexicalStruct{
					Type: 0,
					Text: Text,
					Name: "TEXT",
					Line: ThisStruct.Line,
				}
			}
		} else {
			Text := ThisStruct.Text
			ResMap[ResMapId] = &LexicalStruct{
				Type: 0,
				Text: Text,
				Name: "TEXT",
				Line: ThisStruct.Line,
			}
		}
	}

	return ResMap
}
