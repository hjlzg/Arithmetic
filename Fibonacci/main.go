package main

import (
	"fmt"
)

func main(){
	//1,1,2,3
	const n=100
	for i:=0;i<n;i++ {
		value:=fibValue(i)
		fmt.Println(value)

	}
	

}

//递归实现
func fibValue(n int) int{
	if n==0||n==1 {
		return 1
	}else{
		return fibValue(n-1)+fibValue(n-2)
	}
}

//字典优化肥波切尔数列
func fibValueSEO(n int) int{
	if n<2{
		return 1
	}

	fibDir:=make(map[int]int)
	fibDir[0]=1
	fibDir[1]=1
	for i:=2;i<=n;i++{
		fibDir[i]=fibDir[i-1]+fibDir[i-2]
	}
	return fibDir[n]

}

//数组实现
func fibarray(term int) []int {
    farr := make([]int, term)
    farr[0], farr[1] = 1, 1

    for i:= 2; i < term; i++ {
        farr[i] = farr[i-1] + farr[i-2]
    }
    return farr
}
