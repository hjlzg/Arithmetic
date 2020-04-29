package main

import(
	"fmt"
)

func main(){
	fmt.Print("%b \n",8)
	str:="12345"
	reverStr,state:=reverString(str)
	if(state){
		fmt.Println(reverStr)
	}
}

func reverString(s string) (string, bool) {
    str := []rune(s)
    len := len(str)
    if len > 5000 {
        return string(str), false
    }
    for i := 0; i < len/2; i++ {
		fmt.Println(i)
        str[i], str[len-1-i] = str[len-1-i], str[i]
    }
    return string(str), true
}