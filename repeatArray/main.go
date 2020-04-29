package main

import "fmt"

func main(){
	nums:=[]int{0,1,3,3,4}
	result:=repeatInArray(nums)
	fmt.Println("result:",result)
}

func repeatInArray(array []int) int{
	if array==nil||len(array)<=0 {
		return -1
	}

	for j:=0;j<len(array);j++ {
		if array[j]<0 || array[j]>len(array)-1 {
			return -1
		}
	}

	for i:=0;i<len(array);i++ {
		for array[i]!=i {
			if array[i]==array[array[i]] {
				return array[i]
			}else{
				array[i],array[array[i]]=array[array[i]],array[i]
			}
		}
	}
	return -1
}