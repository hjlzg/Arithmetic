package main

import (
	"fmt"
)

func main(){

	// const n=4

	// //获取打印的最大数
	// maxNum:=GetMaxNum(n)
	// for i:=1;i<=maxNum;i++{
	// 	fmt.Println(i)
	// }

	Print1ToMaxOfDigits(4)
}

func GetMaxNum(n int)int{
	if n==1{
		return 9
	}

	num:=9
	maxNumDir:=make(map[int]int)
	maxNumDir[1]=9

	for i:=2;i<=n;i++{
		num=num*10
		maxNumDir[i]=num+maxNumDir[i-1]
	}
	return maxNumDir[n]
}
//1,9
//2,99 90+9  9*10
//3,999 900+99  9*100
//4,9999 9000+999 9*1000

func Print1ToMaxOfDigits(n int) {
	if n <= 0 {
		return
	}

	number := make([]int, n)
	for i := 0; i < 10; i++ {
		number[0] = i
		print1ToMaxOfDigitsRecursively(number, n, 0)
	}
}

func print1ToMaxOfDigitsRecursively(number []int, length int, index int) {
	if index == length-1 {
		printNumber(number)
		return
	}

	for i := 0; i < 10; i++ {
		number[index+1] = i
		print1ToMaxOfDigitsRecursively(number, length, index+1)
	}
}

func printNumber(number []int) {
	var isBeginning0 = true
	length := len(number)
	for i := 0; i < length; i++ {
		if isBeginning0 && number[i] != 0 {
			isBeginning0 = false
		}

		if !isBeginning0 {
			fmt.Printf("%d", number[i])
			if i == length-1 {
				fmt.Println()
			}
		}
	}
}
