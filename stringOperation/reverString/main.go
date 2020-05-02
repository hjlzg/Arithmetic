package main

import "fmt"

func main(){
	str:="jin中国long"
	fmt.Println(reverString(str))
}


//注意：需要判断是否含有中文
func reverString(str string) string{
	if len(str)==0{
		return ""
	}

	//不区分中英文
	temp:=[]rune(str)

	for i:=0;i<len(temp)/2;i++{
		temp[i],temp[len(temp)-1-i]=temp[len(temp)-1-i],temp[i]
	}
	return string(temp)
}
