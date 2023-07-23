package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := []string{"2", "1", "+", "3", "*"}
	fmt.Println(evalRPN(s))
}

func isValid(s string) bool {
	stack := []byte{}
	m := map[byte]byte{')': '(', ']': '[', '}': '{'}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else if len(stack) > 0 && stack[len(stack)-1] == m[s[i]] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}

func removeDuplicates(s string) string {
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}

func evalRPN(tokens []string) int {
	stack := []int{}
	for i := 0; i < len(tokens); i++ {
		m, err := strconv.Atoi(tokens[i])
		if err == nil {
			stack = append(stack, m)
			continue
		}
		m1 := stack[len(stack)-1]
		m2 := stack[len(stack)-2]
		stack = stack[:len(stack)-2]
		if tokens[i] == "+" {
			stack = append(stack, m2+m1)
		}
		if tokens[i] == "-" {
			stack = append(stack, m2-m1)
		}
		if tokens[i] == "*" {
			stack = append(stack, m2*m1)
		}
		if tokens[i] == "/" {
			stack = append(stack, m2/m1)
		}
	}
	return stack[len(stack)-1]
}
