package parentheses

import (
	"math/rand"
	"time"
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

func Generate(length int) string {
	parentheses := "({[]})"
	result := make([]rune, length)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		result[i] = rune(parentheses[rand.Intn(len(parentheses))])
	}
	return string(result)
}
