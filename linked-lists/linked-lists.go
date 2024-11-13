package main

import "fmt"

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

// Prepend: 將節點插入開頭
func (ll *LinkedList) Prepend(n *Node) {
	// 檢查linked list是否是空的
	if ll.checkIsEmpty() {
		ll.head = n
		ll.tail = n
		return
	}

	n.Next = ll.head // 新節點的Next指向原本的開頭節點的記憶體位置
	ll.head = n      // linked list的頭改為新節點
}

// Insert: 將節點插入特定位置
// 原本: A -> C
// 插入index 1: A -> B -> C，要更改節點A的指向
func (ll *LinkedList) Insert(index int, n *Node) {
	if ll.checkIsEmpty() {
		ll.head = n
		ll.tail = n
		return
	}

	// 如果放入位置是linked list開頭，直接用Prepend method
	if index == 0 {
		ll.Prepend(n)
		return
	}

	// 找到要插入的位置的前一個節點
	// current: 要插入的位置的前一個節點
	current := ll.head
	for i := 0; i < index-1; i++ {
		// 如果指定的index超過linked list長度，直接把新節點放到最後
		// 判斷超出範圍的方式：Next是nil，代表沒有指向下一個節點的記憶體位置
		if current.Next == nil {
			ll.Append(n)
			return
		}

		current = current.Next
	}

	// 插入新節點在指定位置，重新安排指向的節點的記憶體位置
	tempPointer := current.Next // 暫時存放新節點要指向的下一個節點的記憶體位置
	current.Next = n            // 插入節點的前一個節點指向插入的節點
	n.Next = tempPointer        // 插入節點的下個節點
}

// PopFirst: 取出第一個節點
func (ll *LinkedList) PopFirst() *Node {
	// 檢查linked list是否為空
	if ll.checkIsEmpty() {
		fmt.Println("Linked list is empty.")
		return nil
	}

	var firstNode *Node
	// 檢查 linked list 是否只有一個值
	if ll.head == ll.tail {
		firstNode = ll.head
		// 取出後，把節點改成預設值nil
		ll.head = nil
		ll.tail = nil

		return firstNode
	}

	// 重新排序節點
	tempNode := ll.head.Next // 暫存原本的第二個節點
	firstNode = ll.head
	ll.head = tempNode // 原本的第二個節點變成第一個節點

	firstNode.Next = nil // 清空取出的節點的指向記憶體位置
	return firstNode
}

/*
Private methods
*/

// checkIsEmpty: 檢查linked list是否是空的
func (ll *LinkedList) checkIsEmpty() bool {
	return ll.head == nil
}

// printLinkedListToSlice: 以slice的方式顯示linked list資料，例如：[1, 2, 3]
func (ll *LinkedList) printLinkedListToSlice() {
	list := make([]interface{}, 0)

	// 檢查linked list是否為空
	if ll.checkIsEmpty() {
		fmt.Println(list, "This linked list is empty.")
		return
	}

	currentNode := ll.head
	for currentNode != nil {
		list = append(list, currentNode.Value)
		// 移動當前節點位置，以便下一輪迴圈取值
		currentNode = currentNode.Next
	}

	fmt.Println(list)
}

func main() {
	myList := New()
	myList.printLinkedListToSlice()

	myList.Append(NewNode("B"))
	myList.printLinkedListToSlice() // ["B"]

	myList.Prepend(NewNode("A"))
	myList.printLinkedListToSlice() // ["A" "B"]
}
