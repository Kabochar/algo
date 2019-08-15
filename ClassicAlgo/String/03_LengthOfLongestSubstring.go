package String

// 求一个字符串不重复的子串的最大的长度

// location[s[i]] == j 表示：
// s中第i个字符串，上次出现在s的j位置，所以，在s[j+1:i]中没有s[i]
//
// location[s[i]] == -1 表示： s[i] 在s中第一次出现
func lengthOfLongestSubstring(s string) int {
	// 假设输入的只有 ASCII 字符
	location := [256]int{}
	for i := range location {
		// 假设所有字符都没有出现过
		location[i] = -1
	}

	maxLen, left := 0, 0
	for i := 0; i < len(s); i++ {
		// location[s[i]] >= left
		// 说明 s[i]已经在 s[left:i+1] 中重复了
		// 并且s[i]上次出现的位置在location[s[i]]
		if location[s[i]] >= left {
			// 移动，把 s[left:i+1] 中去除 s[i] 字符以及前面的
			// 不断遍历字符串内容
			left = location[s[i]] + 1

			// 当前 (i + 1 - left) > maxLen 才能计算 max Len
		} else if (i + 1 - left) > maxLen {
			// 不重复字符长度，
			maxLen = i + 1 - left
		}

		location[s[i]] = i
	}

	return maxLen
}
