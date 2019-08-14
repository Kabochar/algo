package CartoonAlgo

// 全排列

// 全排列下一个数
// <<漫画算法>>
// 1，从后向前看逆序区域，找到逆序区域的前一位，也就是数字置换的边界
// 2，让逆序区域的前一位和逆序区域中大于它的最小数字交换位置
// 3，把原来的逆序区域转为顺序状态
// Time: O(n)
// Space: O(n)

// 全排列全部可能性