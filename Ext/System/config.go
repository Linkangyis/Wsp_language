package main

import (
	public "Wsp/Public"
)

var ConfigMap = make(map[string]func(map[int]string) public.TypeDLL)

func IoCeng(value func(map[int]string) string) func(map[int]string) public.TypeDLL {
	return func(Value map[int]string) public.TypeDLL {
		return public.TypeDLL{
			Is_Array: false,
			Text:     value(Value),
		}
	}
}

func LOADINIT(Value map[int]string) string {
	return "Sys"
}

func INITS() {
	/* v 请在这里配置函数*/
	ConfigMap["Input"] = IoCeng(Input)
	ConfigMap["Rand"] = IoCeng(Rand)
	ConfigMap["Print"] = IoCeng(Print)
	ConfigMap["Println"] = IoCeng(Println)
	ConfigMap["Stick"] = IoCeng(Stick)
	ConfigMap["Exit"] = IoCeng(Exit)
	ConfigMap["Free"] = IoCeng(Free)
	ConfigMap["Panic"] = IoCeng(Panic)
	ConfigMap["Eval"] = IoCeng(Eval)
	ConfigMap["Exec"] = IoCeng(Exec)
	ConfigMap["Println"] = IoCeng(Println)
	ConfigMap["LOADINIT"] = IoCeng(LOADINIT)
	/* ^ 请在这里配置函数*/
}
