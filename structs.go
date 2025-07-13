package main

import "fmt"

type stack struct {
	index int
	data  [5]int
}

func (s *stack) push(v int) {
	s.data[s.index] = v
	s.index++
}

func (s *stack) pop() int {
	s.index--

	deleted := s.data[s.index]
	s.data[s.index] = 0

	return deleted
}

func main() {
	// declara ponteiro para a pilha
	stack := new(stack)

	for i, _ := range stack.data {
		stack.push(i + 1)
	}

	fmt.Println(stack.data)
	stack.pop()
	fmt.Println(stack.data)
}
