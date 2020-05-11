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

	fmt.Println()
	treeHigh:=treeHigh(tree)
	fmt.Printf("树的高度：%v",treeHigh)

	fmt.Println()
	treeHigh2:=treeHighByQueue(tree)
	fmt.Printf("树的高度(非递归)：%v",treeHigh2)


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

//前序遍历
func leftTraverse(t *TreeNode){
	if t==nil{
		return
	}

	fmt.Printf("%v ",t.val)
	leftTraverse(t.left)
	leftTraverse(t.right)
}

//中序遍历
func middleTraverse(t *TreeNode){
	if t==nil{
		return
	}

	middleTraverse(t.left)
	fmt.Printf("%v ",t.val)
	middleTraverse(t.right)
}

//后序遍历
func rightTraverse(t *TreeNode){
	if t==nil{
		return
	}

	rightTraverse(t.left)
	rightTraverse(t.right)
	fmt.Printf("%v ",t.val)
}

//递归方法实现
func treeHigh(t *TreeNode) int{
	if t==nil{
		return 0
	}

	leftTreeHigh:=treeHigh(t.left)
	rightTreeHigh:=treeHigh(t.right)

	if leftTreeHigh>rightTreeHigh{
		return leftTreeHigh+1
	}else{
		return rightTreeHigh+1
	}
}

func treeHighByQueue(t *TreeNode) int{
	if t==nil{
		return 0
	}

	layers:=0
	nodes:=[]*TreeNode{t}
	for len(nodes)>0{
		layers++
		size:=len(nodes) //每层的节点数
		count:=0
		for count<size{
			count++
			curNode:=nodes[0]
			nodes=nodes[1:]
			if curNode.left!=nil{
				nodes=append(nodes,curNode.left)
			}
			if curNode.right!=nil{
				nodes=append(nodes,curNode.right)
			}
		}
	}
	return layers
}
