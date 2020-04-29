package main

import (
	"strings"
	"fmt"
)

func main(){

	s1:="黄金"
	s2:="黄龙"
	fmt.Print(isRegroup(s1,s2))
}

func isRegroup(s1,s2 string) bool{
	sl1:=len([]rune(s1))
	sl2:=len([]rune(s2))

	if sl1>5000||sl2>500||sl1!=sl2{
		return false
	}

	for _,v:=range s1{
		fmt.Printf("%v \n",v)
		if(strings.Count(s1,string(v))!=strings.Count(s2,string(v))){
			return false
		}
	}
	return true
}