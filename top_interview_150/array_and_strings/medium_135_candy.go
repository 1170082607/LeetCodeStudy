package array_and_strings

/*
n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。

你需要按照以下要求，给这些孩子分发糖果：

每个孩子至少分配到 1 个糖果。
相邻两个孩子评分更高的孩子会获得更多的糖果。
请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目 。



示例 1：

输入：ratings = [1,0,2]
输出：5
解释：你可以分别给第一个、第二个、第三个孩子分发 2、1、2 颗糖果。
示例 2：

输入：ratings = [1,2,2]
输出：4
解释：你可以分别给第一个、第二个、第三个孩子分发 1、2、1 颗糖果。
     第三个孩子只得到 1 颗糖果，这满足题面中的两个条件。


提示：

n == ratings.length
1 <= n <= 2 * 104
0 <= ratings[i] <= 2 * 104
*/
// 135. candy
func candy(ratings []int) int {
	n := len(ratings)
	nums := make([]int, n)
	nums[0] = 1
	for i := 0; i < n-1; i++ {
		if ratings[i] >= ratings[i+1] {
			nums[i+1] = 1
			continue
		}
		nums[i+1] = nums[i] + 1
	}
	for i := n - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] {
			if nums[i-1] < nums[i]+1 {
				nums[i-1] = nums[i] + 1
			}
			continue
		}
	}
	result := 0
	for i := 0; i < n; i++ {
		result += nums[i]
	}
	return result
}

/*
best On O1
func candy(ratings []int) int {
    n := len(ratings)
    ans, inc, dec, pre := 1, 1, 0, 1
    for i := 1; i < n; i++ {
        if ratings[i] >= ratings[i-1] {
            dec = 0
            if ratings[i] == ratings[i-1] {
                pre = 1
            } else {
                pre++
            }
            ans += pre
            inc = pre
        } else {
            dec++
            if dec == inc {
                dec++
            }
            ans += dec
            pre = 1
        }
    }
    return ans
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/candy/solutions/533150/fen-fa-tang-guo-by-leetcode-solution-f01p/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
