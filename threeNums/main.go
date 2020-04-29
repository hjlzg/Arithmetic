package main

import (
	"fmt"
	"sort"
	"reflect"
)

func main(){

	nums:=[]int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
    if len(nums)<3 {
        return nil
    }

    var ret [][]int
    for i:=0;i<len(nums);i++{
        for j:=i+1;j<len(nums);j++{
            for k:=j+1;k<len(nums);k++{
                if nums[i]+nums[j]+nums[k]==0{  
                    var res []int    
                    res=append(res,nums[i])
                    res=append(res,nums[j])
					res=append(res,nums[k])

					sort.Ints(res)
				    fmt.Println("res",res)
					isContain:=false
					for _,value :=range ret {
						if reflect.DeepEqual(value, res) {
							isContain=true
							break
						}
					}

					if(!isContain){
						ret=append(ret,res)
					}
                }

            }
        }
    }

    return ret

}