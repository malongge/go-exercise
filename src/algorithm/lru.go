package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type ILRU interface {
	getCache() (int, int)
}

type LRU struct {
	link *Node
}

func (lru *LRU) appendNode(node *Node) (*Node) {
	if lru.link == nil {
		lru.link = node
		return node
	}
	tail := lru.link
	for tail.next != nil{
		tail = tail.next
	}
	tail.next = node

	return tail.next

}

func (lru *LRU) insertNode(node *Node, value int) int {
	tail := lru.link
	preNode := tail
	for tail != nil {
		if tail.value == value {
			preNode = tail
			break
		}
		tail = tail.next

	}
	if tail == nil {
		return -1
	}

	node.next = preNode.next
	preNode.next = node
	return 0
}

func (lru *LRU) printLink() {
	node := lru.link
	for node != nil {
		fmt.Printf("%d\n", node.value)
		node = node.next
	}
}

func main() {
	lru := new(LRU)
	node1 := Node{1, nil}
	node2 := Node{2, nil}
	node4 := Node{4, nil}
	node3 := Node{3, nil}
	lru.appendNode(&node1)
	lru.appendNode(&node2)
	lru.appendNode(&node4)
	lru.insertNode(&node3, 2)
	lru.printLink()
}
