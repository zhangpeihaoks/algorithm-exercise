// 如果字符串中不含有任何 'aaa'，'bbb' 或 'ccc' 这样的字符串作为子串，那么该字符串就是一个「快乐字符串」。
//
// 给你三个整数 a，b ，c，请你返回 任意一个 满足下列全部条件的字符串 s：
//
// s 是一个尽可能长的快乐字符串。
// s 中 最多 有a 个字母 'a'、b 个字母 'b'、c 个字母 'c' 。
// s 中只含有 'a'、'b' 、'c' 三种字母。
//
// 如果不存在这样的字符串 s ，请返回一个空字符串 ""。
//
// 示例 1：
//
// 输入：a = 1, b = 1, c = 7
// 输出："ccaccbcc"
// 解释："ccbccacc" 也是一种正确答案。
//
// 示例 2：
//
// 输入：a = 2, b = 2, c = 1
// 输出："aabbc"
//
// 示例 3：
//
// 输入：a = 7, b = 1, c = 0
// 输出："aabaa"
// 解释：这是该测试用例的唯一正确答案。
//
// 提示：
//
// 0 <= a, b, c <= 100
// a + b + c > 0
//
// Related Topics 贪心 字符串 堆（优先队列） 👍 212 👎 0
package main

// leetcode submit region begin(Prohibit modification and deletion)
func longestDiverseString(a int, b int, c int) string {
	var ans string
	for a+b+c > 0 {
		//	当前要加入字符串的字符
		var choose string
		choose = "_"
		//	当前最多的球
		chlen := 0
		xsize := len(ans)
		if a > chlen && (xsize < 2 || ans[xsize-1] != 'a' || ans[xsize-2] != 'a') {
			chlen = a
			choose = "a"
		}
		if b > chlen && (xsize < 2 || ans[xsize-1] != 'b' || ans[xsize-2] != 'b') {
			chlen = b
			choose = "b"
		}
		if c > chlen && (xsize < 2 || ans[xsize-1] != 'c' || ans[xsize-2] != 'c') {
			chlen = c
			choose = "c"
		}
		if choose == "_" {
			break
		}
		ans += choose
		if choose == "a" {
			a--
		} else if choose == "b" {
			b--
		} else if choose == "c" {
			c--
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
