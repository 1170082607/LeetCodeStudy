package array_and_strings

/*
给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。

请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。

注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。
*/
// 88. 合并两个有序数组
func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1Index, nums2Index := m-1, n-1
	for i := m + n - 1; i >= 0; i-- {
		if nums1Index < 0 {
			nums1[i] = nums2[nums2Index]
			nums2Index--
			continue
		}
		if nums2Index < 0 {
			nums1[i] = nums1[nums1Index]
			nums1Index--
			continue
		}
		if nums1[nums1Index] > nums2[nums2Index] {
			nums1[i] = nums1[nums1Index]
			nums1Index--
		} else {
			nums1[i] = nums2[nums2Index]
			nums2Index--
		}
	}
	println(11)
}

/*
Best solution
class Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        p1, p2, p = m - 1, n - 1, m + n - 1
        while p2 >= 0:  # nums2 还有要合并的元素
            # 如果 p1 < 0，那么走 else 分支，把 nums2 合并到 nums1 中
            if p1 >= 0 and nums1[p1] > nums2[p2]:
                nums1[p] = nums1[p1]  # 填入 nums1[p1]
                p1 -= 1
            else:
                nums1[p] = nums2[p2]  # 填入 nums2[p1]
                p2 -= 1
            p -= 1  # 下一个要填入的位置
*/

//ProxyCommand "C:\Program Files\Git\mingw64\bin\connect.exe" -S 192.168.1.1:20171 -a none %h %p
