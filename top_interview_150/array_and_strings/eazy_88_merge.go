package array_and_strings

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
