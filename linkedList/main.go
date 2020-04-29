package LinkedList

import(
	"fmt"
	"reflect"
)

//链表中的数据类型
type ElementType interface{}

//节点结构体
type Node struct{
	node ElementType  //节点
	next *Node		 //下一节点指针
}

//链表结构体
type LinkedList struct{
	Head *Node
}

//添加新节点
func NewNode(x ElementType,next *Node) *Node{
	return &Node{x,next}
}

//创建新链表
func NewLinkedList() *LinkedList{
	head:=&Node{0,nil}
	return &LinkedList{head}
}

//判断是否为空
func (list *LinkedList) IsEmpty() bool{
	return list.Head.next==nil
}

//链表的长度
func (list *LinkedList) Length() int{
	// return int(reflect.ValueOf(list.Head.Node).Int())

	if list.Head.next==nil{
		return 0
	}

	temp:=list.Head
	number:=0
	for temp!=nil{
		temp=temp.next
		number++
	}
	return number
}

//向链表后端Append元素
func (list *LinkedList) Append(node *Node){
	current:=list.Head
	for {
		if current.next==nil{
			break
		}
		current=current.next
	}
	current.next=node
}

func (list *LinkedList) PreAppend(node *Node){
	current=list.Head
	node.next=current.next
	current.next=node
}



