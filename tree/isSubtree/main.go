package main

func main(){
}

type Treenode struct {
    val int
    eft *treenode
    right *treenode
 }

func isSubtree(s *TreeNode, t *TreeNode) bool {
    if s == nil {
        return false
    }
    return check(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func check(a, b *TreeNode) bool {
    if a == nil && b == nil {
        return true
    }
    if a == nil || b == nil {
        return false
    }
    if a.Val == b.Val {
        return check(a.Left, b.Left) && check(a.Right, b.Right)
    }
    return false
}