package main

import (
	"fmt"
)

func balancedBracket(s string) string {
	stack := []rune{}
	bracketPairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		switch char {
		case '(', '{', '[':
			stack = append(stack, char)
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != bracketPairs[char] {
				return "NO"
			}
			// Delete last char in stack
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) == 0 {
		return "YES"
	}
	return "NO"
}

func main() {
	inputs := []string{
		"{ [ ( ) ] }",             // Yes
		"{ [ ()]}",                // Yes
		"{ [ ( ] ) }",             // No
		"{ ( ( [ ] ) [ ] ) [ ] }", // Yes
	}

	for _, input := range inputs {
		fmt.Printf("Input: %s\n", input)
		fmt.Printf("Output: %s\n\n", balancedBracket(input))
	}
}
