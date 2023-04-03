package parentheses

func isBalanced(str string) bool {
	stack := make([]rune, 0)
	pares := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, bracket := range str {
		if _, ok := pares[bracket]; ok {
			stack = append(stack, bracket)
		} else if len(stack) > 0 && pares[bracket] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
