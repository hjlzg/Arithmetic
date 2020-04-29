//翻转字符串

package main

import (
	"fmt"
)

func main(){
	str:="123234feqf34qwf"
	reStr:=""
	// for i:=len(str)-1;i>=0;i--{
	// 	reStr+=string(str[i])
	// }
	
	userMap:=make(map[int]string)
	temp:=0
	for _,v:=range str{
		fmt.Println(temp,v)
		userMap[temp]=string(v)
		temp++
	}
	fmt.Println(len(userMap))

	if userMap!=nil{
		for i:=len(userMap)-1;i>=0;i--{
			reStr+=userMap[i]
		}
	}

	fmt.Println(reStr)
}