package main

import "fmt"

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
	list := BuildLinkedList()
	list.InsertIndex(0, 0)
	list.InsertIndex(1, 1)
	list.Print()
}
