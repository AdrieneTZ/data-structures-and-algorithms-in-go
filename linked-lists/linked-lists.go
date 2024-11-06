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
