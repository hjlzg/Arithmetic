package main

import "fmt"

type TreeNode struct {
    val int
    left *TreeNode
    right *TreeNode
}

func main(){
    //创建s树
     sTree:=&TreeNode{3,nil,nil}
     sTree.left=&TreeNode{4,nil,nil}
     sTree.right=&TreeNode{5,nil,nil}
     sTree.left.left=&TreeNode{1,nil,nil}
     sTree.left.right=&TreeNode{2,nil,nil}

     //创建t树
     tTree:= &TreeNode{4,nil,nil}
     tTree.left=&TreeNode{1,nil,nil}
     tTree.right=&TreeNode{2,nil,nil}

     value:=isSubTree2(sTree,tTree)
     fmt.Printf("t是s的子树:%v",value)
}

func isSubTree2(s *TreeNode,t *TreeNode) bool{
    if s==nil{
        return false
    }
    return check2(s,t)||isSubTree2(s.left,t)||isSubTree2(s.right,t)
}

func check2(s,t *TreeNode) bool{
    if s==nil&&t==nil{
        return true
    }

    if s==nil||t==nil{
        return false
    }

    if s.val==t.val{
        return check2(s.right,t.right)&&check2(s.left,s.right)
    }
    return false
}



