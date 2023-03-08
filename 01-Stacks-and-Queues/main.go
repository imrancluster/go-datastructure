package main

import "fmt"

func main() {

	// example of stack
	myStack := Stack{}
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	fmt.Println(myStack)

	myStack.Pop()
	myStack.Pop()
	fmt.Println(myStack)
	// end

	// example of queue
	myQueue := Queue{}
	myQueue.Enqueue(100)
	myQueue.Enqueue(200)
	myQueue.Enqueue(300)
	fmt.Println(myQueue)

	myQueue.Dequeue()
	myQueue.Dequeue()
	fmt.Println(myQueue)

}
