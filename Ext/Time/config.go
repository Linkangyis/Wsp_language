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
	return "Time"
}

func INITS() {
	/* v 请在这里配置函数*/
	ConfigMap["Time"] = IoCeng(Time)
	ConfigMap["Sleep"] = IoCeng(Sleep)
	ConfigMap["LOADINIT"] = IoCeng(LOADINIT)
	/* ^ 请在这里配置函数*/
}
