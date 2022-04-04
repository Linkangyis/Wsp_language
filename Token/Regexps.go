package token

import(
    "regexp"
)

func regexps(a string,b string)(string){
    re := regexp.MustCompile(a)
    match := re.ReplaceAllString(b,"\n")
    return match
}
func Code_Notes(text string)(string){
    match:=regexps(`//(.*)\n`,text)
    match=regexps(`#(.*)\n`,match)
    return match
}