package main

type Node struct {
	Value interface{}
	Next  *Node // 指向下一個node struct的記憶體位置
}

type LinkedList struct {
	head *Node // 指向第一個node struct的記憶體位置
	tail *Node // 指向最後一個node struct的記憶體位置
}

// 建立新的記憶體空間，用來存放 LinkedList struct
// 同時初始化所有欄位為零值(head = nil, tail = nil)
func New() *LinkedList {
	return new(LinkedList)
}

func NewNode(value interface{}) *Node {
	n := new(Node)
	n.Value = value
	return n
}

// 實作methods
/*
Public methods
*/

// Append: 將節點插入結尾
func (ll *LinkedList) Append(n *Node) {
	// 檢查linked list是否是空的
	if ll.checkIsEmpty() {
		ll.head = n
		ll.tail = n
		return
	}

	ll.tail.Next = n // 原本是 nil，現在指向新節點的記憶體位置
	ll.tail = n      // 把linked list的最後一個節點改成新節點
}

/*
Private methods
*/

// checkIsEmpty: 檢查linked list是否是空的
func (ll *LinkedList) checkIsEmpty() bool {
	return ll.head == nil
}
