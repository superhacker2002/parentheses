package parentheses

import (
	"math/rand"
)

type Parentheses struct{}

func New() Parentheses {
	return Parentheses{}
}

func (p Parentheses) IsBalanced(str string) bool {
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

// GenerateRandom function requires random seed to be generated before call
func (p Parentheses) GenerateRandom(length uint) string {
	const Parentheses = "({[]})"
	result := make([]rune, length)
	for i := uint(0); i < length; i++ {
		result[i] = rune(Parentheses[rand.Intn(len(Parentheses))])
	}
	return string(result)
}
