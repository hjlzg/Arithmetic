package main

import (
	"fmt"
	"strings"
)

func main(){
	//
	// LengthOfSubString("123")
	result:=LengthOfSubString("pwwkew")
	fmt.Println(result)
}

func LengthOfSubString(s string) int{
	if len(s)==0{
		return 0
	}

	maxLength:=0
	for j:=len(s);j>0;j--{
		for i:=0;i<j;i++{
			value:=s[i:j]
			if len(value)==filterRepeatStr(value)&&maxLength<len(value) {
				fmt.Println(value)
				maxLength=len(value)
			}
		}
	}	
	return maxLength
}

func filterRepeatStr(s string) int{
	value:=""
	for i:=0;i<len(s);i++{
		if strings.Index(value,string(s[i]))==-1{
			value+=string(s[i])
		}
	}
	return len(value)
}