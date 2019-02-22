package main

import "fmt"

// 单链表
// 物理存储方式：链式存储结构
// 使用基本数据类型和指针，练习完成数组的实现

type Node struct {
	Value interface{}
	Next  *Node
}

type LinkedList struct {
	Head   *Node
	Length int
}

func BuildLinkedList() (*LinkedList) {
	return &LinkedList{
		Head: &Node{
			Value: 0,
			Next:  nil,
		},
		Length: 0,
	}
}

// 以下为尾插法初始化
func BuildLinkedListTailBySlice(slice []interface{}) (*LinkedList) {
	list := &LinkedList{Head: &Node{Value: 0, Next: nil,}, Length: 0,}
	temp := list.Head
	for _, value := range slice {
		node := &Node{Value: value, Next: nil,}
		temp.Next = node
		temp = node
	}
	return list
}

// 以下为头插法初始化
func BuildLinkedListHeadBySlice(slice []interface{}) (*LinkedList) {
	list := &LinkedList{Head: &Node{Value: 0, Next: nil,}, Length: 0,}
	for _, value := range slice {
		node := &Node{Value: value, Next: list.Head.Next,}
		list.Head.Next = node
	}
	return list
}

func (list *LinkedList) Clear() () {
	temp := list.Head
	for temp.Next != nil {
		temp.Next = temp.Next.Next
	}
}

func (list *LinkedList) DeleteIndex(index int) () {
	if index >= list.Length || index < 0 {
		panic(fmt.Errorf("not valid index"))
	}
	temp := list.Head
	for i := 0; i < index; i++ {
		temp = temp.Next
	}
	temp.Next = temp.Next.Next
}

func (list *LinkedList) InsertIndex(index int, value interface{}) () {
	if index > list.Length || index < 0 {
		panic(fmt.Errorf("not valid index"))
	}
	temp := list.Head
	for i := 0; i < index; i++ {
		temp = temp.Next
	}
	newNode := &Node{
		Value: value,
		Next:  temp.Next,
	}
	temp.Next = newNode
	list.Length++
}

func (list *LinkedList) Insert(node *Node, value interface{}) {
	newNode := &Node{
		Value: value,
		Next:  node.Next,
	}
	node.Next = newNode
	list.Length++
}

func (list *LinkedList) Get(index int) (interface{}) {
	if index >= list.Length || index < 0 {
		panic(fmt.Errorf("not valid index"))
	}
	temp := list.Head
	for i := -1; i < index; i++ {
		temp = temp.Next
	}
	return temp.Value
}

func (list *LinkedList) Print() () {
	temp := list.Head.Next
	for temp != nil {
		fmt.Println(temp.Value)
		temp = temp.Next
	}
}

func main() {
	fmt.Println("------------before insert index--------------")
	list := BuildLinkedList()
	list.InsertIndex(0, 0)
	list.InsertIndex(1, 1)
	list.Print()
	fmt.Println("-----------before delete index---------------")
	list.DeleteIndex(1)
	list.Print()
	fmt.Println("-------------before head build-------------")
	list = BuildLinkedListHeadBySlice([]interface{}{5, 4, 3, 2, 1})
	list.Print()
	fmt.Println("--------------before tail build------------")
	list = BuildLinkedListTailBySlice([]interface{}{5, 4, 3, 2, 1})
	list.Print()
	fmt.Println("------------before clear--------------")
	list.Clear()
	list.Print()
	fmt.Println("-------------end main-------------")
}
