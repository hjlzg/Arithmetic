package main

import (
	"fmt"
)

func main(){
	// l1:=&ListNode{
	// 	Val:2,
	// 	Next:&ListNode{
	// 		Val:4,
	// 		Next:&ListNode{
	// 			Val:3,
	// 			Next:nil,
	// 		},
	// 	},
	// }

	// l2:=&ListNode{
	// 	Val:5,
	// 	Next:&ListNode{
	// 		Val:6,
	// 		Next:&ListNode{
	// 			Val:4,
	// 			Next:nil,
	// 		},
	// 	},
	// }

	// fmt.Println(addTwoNumbers(l1,l2))

	fmt.Println(lengthOfLongestSubstring("pwwkew"))

}

type ListNode struct{
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode,l2 *ListNode) *ListNode{
	if l1==nil{
		return l2
	}
	if l2==nil{
		return l1
	}

	sum:=l1.Val+l2.Val
	nextNode:=addTwoNumbers(l1.Next,l2.Next)
	if sum<10{
		return &ListNode{Val:sum,Next:nextNode}
	}else{
		tempNode:=&ListNode{
			Val:1,
			Next:nil,
		}
		return &ListNode{
			Val:sum-10,
			Next:addTwoNumbers(nextNode,tempNode),
		}
	}
}

func lengthOfLongestSubstring(s string) int {
    dir:=make(map[string]int)
    for i:=0;i<len(s);i++{
       if _, ok := dir[string(s[i])]; ok {
            continue
        }else{
            dir[string(s[i])]=i
        }
    }
    return len(dir)
}