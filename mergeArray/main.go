package main

import (
	"fmt"
)

func main(){

	array:=[][]int{{1,3},{2,4},{5,6}}
	mergeArray:=make([]int,0)

	arrayLength:=len(array)
	arrayWidth:=len(array[0])

	//数组组装成切片
	for i:=0;i<arrayLength;i++{
		for j:=0;j<arrayWidth;j++{
			mergeArray=append(mergeArray,array[i][j])
		}
	}
	fmt.Println(arrayOrder(mergeArray))

}

func arrayOrder(slice []int) []int{
	tempSelice:=slice
	for i:=0;i<len(tempSelice);i++{
		for j:=i;j<len(tempSelice);j++{
			if slice[i]>tempSelice[j]{
				tempSelice[j],tempSelice[i]=tempSelice[i],tempSelice[j]
			}
		}
	}
	return tempSelice
}