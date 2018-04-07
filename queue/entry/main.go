package main

import (
	"spiderProject/queue"
	"fmt"
)

func main() {
	q := queue.Queue{1, 2, 3}
	fmt.Println(q)
	q.Push(4)
	q.Push(5)
	fmt.Println(q)

	q.Pop()
	q.Pop()
	fmt.Println(q)

	q.Pop()
	q.Pop()
	q.Pop()
	fmt.Println(q.IsEmpty())
}
