package main

import (
	"fmt"
	"strings"
)

func main(){
	str:="123 123 321"
	fmt.Println(replaceBlank(str))
}


func replaceBlank(str string) string{
	return strings.Replace(str," ","",-1)
}