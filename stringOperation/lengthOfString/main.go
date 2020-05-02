package main

import (
	"fmt"
	"strings"
)

//例如：str:="pwwkew"
//输出：qwke 5
//解题思路：设置一个中间变量，用于存放不重复的字符串

func main(){
	str:="pwwkew"
	fmt.Print("不重复字符串长度为:",lengthOfSubString(str))
}

func lengthOfSubString(str string) int{
	if len(str)==0{
		return 0
	}

	result:=""
	for _,v:=range str{
		if strings.Index(result,string(v))==-1{
			fmt.Println(string(v))
			result+=string(v)
		}
	}
	return len(result)
}