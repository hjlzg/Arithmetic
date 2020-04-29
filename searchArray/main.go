package main

import (
	"fmt"
)

func main(){
	a := [][]int{  
		{0, 1, 2, 3} ,   /*  第一行索引为 0 */
		{4, 5, 6, 7} ,   /*  第二行索引为 1 */
		{8, 9, 10, 11},   /* 第三行索引为 2 */
	   }

	result:=searchArray(a,12)
	fmt.Println("result:",result)

}

func searchArray(array [][]int,target int) bool{
	if len(array)==0||len(array[0])==0{
		return false
	}

	var row=0
	var array1=array[0]
	var column=len(array1)-1

	for column>0&&row<len(array)-1 {
		if array[row][column]==target {
			return true
		}else if array[row][column]>target {
			column--
		}else{
			row++
		}
	}
	return false
}