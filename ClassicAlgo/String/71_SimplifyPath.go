package String

import (
	"strings"
)

// 参考：https://www.cnblogs.com/ganganloveu/p/3782727.html
// 核心在于编写一个以'/'为分隔符的split函数
//
// 以及用进出栈来保存最简路径。
//
// path:"/a/./b/../../c/"
//
// split:"a",".","b","..","..","c"
//
// stack:push(a), push(b), pop(b), pop(a), push(c) --> c

func simplifyPath(path string) string {
	lens := len(path)
	stack := make([]string, 0, lens/2)
	dir := make([]byte, 0, lens)

	for i := 0; i < lens; i++ {
		// 使用前，清空 dir，骚操作，比 []byte{} 效率更高
		dir = dir[:0]
		// 获取 dir，拆分斜杠中间的内容
		for i < lens && path[i] != '/' {
			dir = append(dir, path[i])
			i++
		}

		str := string(dir)
		switch str {
		case ".", "":
			// 不做操作
		case "..": // .. 表示上一级操作，清空栈内容
			if len(stack) > 0 {
				// 清空栈内容
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, str)
		}
	}

	return "/" + strings.Join(stack, "/")
}
