package Leetcode

// 要求：括号配对检测
// PS：如果有一个不能完整匹配，直接返回错误
// 这里使用 ascii 底层编码进行配对
func isValid(s string) bool {
	size := len(s)
	stack := make([]byte, size)
	top := 0

	for i := 0; i < size; i++ {
		cur := s[i]
		switch cur {
		case '(': // 左括号直接压栈
			stack[top] = cur + 1
			top++
		case '[', '{':
			stack[top] = cur + 2
			top++
		case ')', ']', '}': // 右括号与栈顶检测配对
			if top > 0 && stack[top-1] == cur {
				top--
			} else {
				return false
			}
		}
	}

	// 配对成功标志：栈空
	return top == 0
}
