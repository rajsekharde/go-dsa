# Data Structures & Algorithms library for Golang
Contains implementation of commonly used data structures and algorithms in Golang

## Import and use

Queue:
```bash
go get github.com/rajsekharde/go-dsa/queue
```
```bash
import "github.com/rajsekharde/go-dsa/queue"

func main() {
    q := queue.NewQueue[int]()

    q.Push(10)
    q.Push(20)
    f1, res := q.Front()
    f2, res := q.Pop()
}
```

Stack:
```bash
import "github.com/rajsekharde/go-dsa/stack"

func main() {
    s := stack.NewStack[string]()

    s.Push(10)
    s.Push(20)
    f1, res := s.Top()
    f2, res := s.Pop()
}
```