package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value int
		prev  *node
	}
)

func New() *Stack {
	return &Stack{nil, 0}
}

func (stack *Stack) Pop() (int, error) {
	if stack.length == 0 {
		return 0, errors.New("stack is empty")
	}

	n := stack.top
	stack.top = n.prev
	stack.length--

	return n.value, nil
}

func (stack *Stack) Push(value int) {
	n := &node{value, stack.top}
	stack.top = n
	stack.length++
}

func calculate(scanner *bufio.Scanner) (int, error) {
	stack := New()

	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())

		if err == nil { // is int
			stack.Push(number)
		} else {
			switch scanner.Text()[0] {
			case ' ':
			case '\n':
			case '=':
				result, err := stack.Pop()

				if err != nil {
					return 0, err
				}

				return result, nil
			case '+':
				first, err := stack.Pop()
				if err != nil {
					return 0, err
				}

				second, err := stack.Pop()
				if err != nil {
					return 0, err
				}

				stack.Push(first + second)
			case '-':
				first, err := stack.Pop()
				if err != nil {
					return 0, err
				}

				second, err := stack.Pop()
				if err != nil {
					return 0, err
				}

				stack.Push(-first + second)
			case '*':
				first, err := stack.Pop()
				if err != nil {
					return 0, err
				}

				second, err := stack.Pop()
				if err != nil {
					return 0, err
				}

				stack.Push(first * second)
			case '/':
				first, err := stack.Pop()
				if err != nil {
					return 0, err
				}

				second, err := stack.Pop()
				if err != nil {
					return 0, err
				}

				stack.Push(second / first)
			default:
				return 0, errors.New("undefined symbol")
			}
		}
	}

	result, err := stack.Pop()

	if err != nil {
		return 0, err
	}

	return result, nil
}

func main() {
	input := bytes.NewBufferString(`2 3 4 5 6 * + - / =`)
	result, err := calculate(input)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Result:", result)
	}
}
