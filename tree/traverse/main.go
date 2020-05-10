package main

import "fmt"

type TreeNode struct {
	val int
	left *TreeNode
	right *TreeNode
}

func main(){
	tree:=createTree()
	leftTraverse(tree)
	fmt.Println()
	middleTraverse(tree)

	fmt.Println()
	rightTraverse(tree)
}

func createTree() *TreeNode{
	//创建s树
	sTree:=&TreeNode{3,nil,nil}
	sTree.left=&TreeNode{4,nil,nil}
	sTree.right=&TreeNode{5,nil,nil}
	sTree.left.left=&TreeNode{1,nil,nil}
	sTree.left.right=&TreeNode{2,nil,nil}

	return sTree
}

func leftTraverse(t *TreeNode){
	if t==nil{
		return
	}

	fmt.Printf("%v ",t.val)
	leftTraverse(t.left)
	leftTraverse(t.right)
}

func middleTraverse(t *TreeNode){
	if t==nil{
		return
	}

	middleTraverse(t.left)
	fmt.Printf("%v ",t.val)
	middleTraverse(t.right)
}

func rightTraverse(t *TreeNode){
	if t==nil{
		return
	}

	rightTraverse(t.left)
	rightTraverse(t.right)
	fmt.Printf("%v ",t.val)
}
