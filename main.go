package main

import (
	"fmt"

	"github.com/rajsekharde/go-dsa/queue"
	"github.com/rajsekharde/go-dsa/stack"
	"github.com/rajsekharde/go-dsa/hashset"
)

func main() {
	// testQueue()
	// testStack()
	// testHashSet()
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

func testHashSet() {
	s := hashset.NewHashSet[int]()
	s.Add(10)
	s.Add(13)
	s.Add(6)
	fmt.Println(s.Contains(10))
	fmt.Println(s.Contains(24))
	fmt.Println(s.Size())
	s.Remove(13)
	s.Display()
}