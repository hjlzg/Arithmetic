package main

import (
	"fmt"
	"strings"
)

func main(){
	str1:="中国jinlong123"
	str2:="中国jinlong123"
	fmt.Print("两个字符串是否相同：",completeEqual(str1,str2))
}

func completeEqual(str1,str2 string) bool{
	if len(str1)!=len(str2){
		return false
	}

	for _,v:=range str1{
		fmt.Println(string(v))
		if strings.Index(str1,string(v))!=strings.Index(str2,string(v)){
			return false
		}
	}
	return true
}
