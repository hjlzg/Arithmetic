package main

import(
	"strings"
	"fmt"
)

func main(){
	str:="12aA"
	fmt.Print(isUniqueString(str))
}

func isUniqueString(s string) bool {
	if strings.Count(s,"") > 3000{
		return  false
	}
	for _,v := range s {
		fmt.Print(" ",v)
		if v > 127 {
			return false
		}
		if strings.Count(s,string(v)) > 1 {
			return false
		}
	}
	return true
}