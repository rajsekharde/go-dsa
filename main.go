package main

import (
	"fmt"

	"github.com/rajsekharde/go-dsa/queue"
	"github.com/rajsekharde/go-dsa/stack"
)

func main() {
	// testQueue()
	// testStack()
}

func testQueue() {
	q := queue.NewQueue[int]()
	q.Push(10)
	q.Push(20)
	q.Push(30)
	q.Display()
	fmt.Println(q.Size())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	q.Display()
	fmt.Println(q.Size())
}

func testStack() {
	s := stack.NewStack[string]()
	s.Push("a")
	s.Push("b")
	s.Push("c")
	s.Display()
	fmt.Println(s.Size())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	s.Display()
	fmt.Println(s.Size())
}