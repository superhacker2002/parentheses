package parentheses

import (
	"fmt"
	"math/rand"
)

func IsBalanced(str string) bool {
	stack := make([]rune, 0)
	pares := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, char := range str {
		if bracket, ok := pares[char]; ok {
			if len(stack) > 0 && bracket == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}

// GenerateRandomSequence function requires random seed to be generated before call
func GenerateRandomSequence(length uint) string {
	fmt.Println(length)
	const parentheses = "({[]})"
	result := make([]rune, length)
	for i := uint(0); i < length; i++ {
		result[i] = rune(parentheses[rand.Intn(len(parentheses))])
	}
	return string(result)
}
