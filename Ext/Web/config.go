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
	return "Web"
}

func INITS() {
	/* v 请在这里配置函数*/
	ConfigMap["New_Web"] = IoCeng(New_Web)
	ConfigMap["Print"] = IoCeng(Print)
	ConfigMap["Start"] = IoCeng(Start)
	ConfigMap["Start_Ssl"] = IoCeng(Start_Ssl)
	ConfigMap["Header_Set"] = IoCeng(Header_Set)
	ConfigMap["New_WebFiles"] = IoCeng(New_WebFiles)
	ConfigMap["SslGoto"] = IoCeng(SslGoto)
	ConfigMap["GET"] = IoCeng(GET)
	ConfigMap["POST"] = IoCeng(POST)
	ConfigMap["WebPath"] = IoCeng(WebPath)
	ConfigMap["LOADINIT"] = IoCeng(LOADINIT)
	/* ^ 请在这里配置函数*/
}
