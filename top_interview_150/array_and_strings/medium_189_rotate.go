package array_and_strings

/*
给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。



示例 1:

输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]
示例 2:

输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释:
向右轮转 1 步: [99,-1,-100,3]
向右轮转 2 步: [3,99,-1,-100]


提示：

1 <= nums.length <= 105
-231 <= nums[i] <= 231 - 1
0 <= k <= 105


进阶：

尽可能想出更多的解决方案，至少有 三种 不同的方法可以解决这个问题。
你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？
*/
// 189. 轮转数组
// my wrong solution
//func rotate(nums []int, k int) {
//	count, temp, lenNums := 0, 0, len(nums)
//	k = k % lenNums
//	if k == 0 {
//		return
//	}
//	count = lenNums / k
//	if lenNums%k != 0 {
//		count += 1
//	}
//	if count == 1 {
//		count += 1
//	}
//	for i := 1; i < count; i++ {
//		for j := 0; j < k; j++ {
//			temp = nums[(j+i*k)%lenNums]
//			nums[(j+i*k)%lenNums] = nums[j]
//			nums[j] = temp
//		}
//	}
//	for i := lenNums % k; i < k-lenNums%k; i++ {
//		temp = nums[(i+i*k)%lenNums]
//		nums[(j+i*k)%lenNums] = nums[j]
//		nums[j] = temp
//	}
//
//	println(11)
//}

/*
方法一中使用额外数组的原因在于如果我们直接将每个数字放至它最后的位置，这样被放置位置的元素会被覆盖从而丢失。因此，从另一个角度，我们可以将被替换的元素保存在变量 temp\textit{temp}temp 中，从而避免了额外数组的开销。

我们从位置 000 开始，最初令 temp=nums[0]\textit{temp}=\textit{nums}[0]temp=nums[0]。根据规则，位置 000 的元素会放至 (0+k) mod n(0+k)\bmod n(0+k)modn 的位置，令 x=(0+k) mod nx=(0+k)\bmod nx=(0+k)modn，此时交换 temp\textit{temp}temp 和 nums[x]\textit{nums}[x]nums[x]，完成位置 xxx 的更新。然后，我们考察位置 xxx，并交换 temp\textit{temp}temp 和 nums[(x+k) mod n]\textit{nums}[(x+k)\bmod n]nums[(x+k)modn]，从而完成下一个位置的更新。不断进行上述过程，直至回到初始位置 000。

容易发现，当回到初始位置 000 时，有些数字可能还没有遍历到，此时我们应该从下一个数字开始重复的过程，可是这个时候怎么才算遍历结束呢？我们不妨先考虑这样一个问题：从 000 开始不断遍历，最终回到起点 000 的过程中，我们遍历了多少个元素？

由于最终回到了起点，故该过程恰好走了整数数量的圈，不妨设为 aaa 圈；再设该过程总共遍历了 bbb 个元素。因此，我们有 an=bkan=bkan=bk，即 ananan 一定为 n,kn,kn,k 的公倍数。又因为我们在第一次回到起点时就结束，因此 aaa 要尽可能小，故 ananan 就是 n,kn,kn,k 的最小公倍数 lcm(n,k)\text{lcm}(n,k)lcm(n,k)，因此 bbb 就为 lcm(n,k)/k\text{lcm}(n,k)/klcm(n,k)/k。

这说明单次遍历会访问到 lcm(n,k)/k\text{lcm}(n,k)/klcm(n,k)/k 个元素。为了访问到所有的元素，我们需要进行遍历的次数为

nlcm(n,k)/k=nklcm(n,k)=gcd(n,k)\frac{n}{\text{lcm}(n,k)/k}=\frac{nk}{\text{lcm}(n,k)}=\text{gcd}(n,k)
lcm(n,k)/k
n
​
 =
lcm(n,k)
nk
​
 =gcd(n,k)
其中 gcd\text{gcd}gcd 指的是最大公约数。

作者：力扣官方题解
链接：https://leetcode.cn/problems/rotate-array/solutions/551039/xuan-zhuan-shu-zu-by-leetcode-solution-nipk/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
// 189. 轮转数组 Best solution 1
func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	for start, count := 0, gcd(k, n); start < count; start++ {
		pre, cur := nums[start], start
		for ok := true; ok; ok = cur != start {
			next := (cur + k) % n
			nums[next], pre, cur = pre, nums[next], next
		}
	}
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

/* Best solution 2:
func reverse(a []int) {
    for i, n := 0, len(a); i < n/2; i++ {
        a[i], a[n-1-i] = a[n-1-i], a[i]
    }
}

func rotate(nums []int, k int) {
    k %= len(nums)
    reverse(nums)
    reverse(nums[:k])
    reverse(nums[k:])
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/rotate-array/solutions/551039/xuan-zhuan-shu-zu-by-leetcode-solution-nipk/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
