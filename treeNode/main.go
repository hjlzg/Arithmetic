package main

import (
	"fmt"
	"math"
)

type treeNode struct{
	value string
	left,right *treeNode
}

func main(){

	//创建一颗树
	root:=treeNode{"A",nil,nil}
	root.left=&treeNode{value:"B"}
	root.right=&treeNode{value:"C"}
	root.left.left=&treeNode{value:"D"}
	root.left.right = &treeNode{value:"F"}
	// root.left.right.left = new(treeNode)
	// root.left.right.left.value = "E"
	root.left.right.left=&treeNode{value:"E"}
	root.left.right.left.right=&treeNode{value:"J"}
 
	root.right.left = &treeNode{value:"G"}
	root.right.left.right = &treeNode{value:"H"}
	root.right.right = &treeNode{value:"I"}

	root.traverse()
	maxDeep:=root.printMaxDeep()
	fmt.Println("最大深度：",maxDeep)
}

func (node *treeNode)traverse(){
	if(node == nil){
	   return
	}
	fmt.Print(node.value + " ")
	node.left.traverse()
	node.right.traverse()
 }

 func (node *treeNode)printMaxDeep() int{
	 if node==nil {
		 return 0
	 }
	 left:=node.left.printMaxDeep()+1
	 right:=node.right.printMaxDeep()+1

	 return int(math.Max(float64(left),float64(right)))
 }
 