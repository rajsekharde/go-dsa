package main

import (
	"fmt"

	"github.com/rajsekharde/go-dsa/queue"
)

func main() {
	testQueue()
}

func testQueue() {
	q := queue.NewQueue[int]()
	q.Push(10)
	q.Push(20)
	q.Push(30)
	q.Display()
	q.Pop()
	fmt.Println(q.Front())
	fmt.Println(q.Size())
}