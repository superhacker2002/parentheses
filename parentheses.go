package parentheses

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
