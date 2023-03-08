package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

// Insert a node at the beginning of a linked list
func (ll *LinkedList) InsertAtBeginning(data int) {
	newNode := &Node{
		data: data,
		next: ll.head,
	}
	ll.head = newNode
}

// Print all nodes in a linked list
func (ll *LinkedList) PrintList() {
	currentNode := ll.head
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}
}

func main() {
	//@TODO: implement example of linked lists
}
