package array_and_strings

/*
给你一个整数数组 citations ，其中 citations[i] 表示研究者的第 i 篇论文被引用的次数。计算并返回该研究者的 h 指数。

根据维基百科上 h 指数的定义：h 代表“高引用次数” ，一名科研人员的 h 指数 是指他（她）至少发表了 h 篇论文，并且 至少 有 h 篇论文被引用次数大于等于 h 。如果 h 有多种可能的值，h 指数 是其中最大的那个。



示例 1：

输入：citations = [3,0,6,1,5]
输出：3
解释：给定数组表示研究者总共有 5 篇论文，每篇论文相应的被引用了 3, 0, 6, 1, 5 次。
     由于研究者有 3 篇论文每篇 至少 被引用了 3 次，其余两篇论文每篇被引用 不多于 3 次，所以她的 h 指数是 3。
示例 2：

输入：citations = [1,3,1]
输出：1


提示：

n == citations.length
1 <= n <= 5000
0 <= citations[i] <= 1000
*/
// 274. H 指数
func hIndex(citations []int) int {
	hNums := make([]int, len(citations)+1)
	for _, value := range citations {
		num := value
		if num >= len(hNums) {
			num = len(hNums) - 1
		}
		for ; num > 0; num-- {
			hNums[num] += 1
		}
	}
	for i := len(hNums) - 1; i >= 0; i-- {
		if i <= hNums[i] {
			return i
		}
	}
	return 0
}

/* best solution
func hIndex(citations []int) (h int) {
	n := len(citations)
	counter := make([]int, n+1)
	for _, citation := range citations {
		if citation >= n {
			counter[n]++
		} else {
			counter[citation]++
		}
	}
	for i, tot := n, 0; i >= 0; i-- {
		tot += counter[i]
		if tot >= i {
			return i
		}
	}
	return 0
}

func hIndex(citations []int) (h int) {
    sort.Ints(citations)
    for i := len(citations) - 1; i >= 0 && citations[i] > h; i-- {
        h++
    }
    return
}

方法三：二分搜索
我们需要找到一个值 h，它是满足「有 h 篇论文的引用次数至少为 h」的最大值。小于等于 h 的所有值 x 都满足这个性质，而大于 h 的值都不满足这个性质。同时因为我们可以用较短时间（扫描一遍数组的时间复杂度为 O(n)O(n)O(n)，其中 nnn 为数组 citations\textit{citations}citations 的长度）来判断 x 是否满足这个性质，所以这个问题可以用二分搜索来解决。

设查找范围的初始左边界 leftleftleft 为 000，初始右边界 rightrightright 为 nnn。每次在查找范围内取中点 midmidmid，同时扫描整个数组，判断是否至少有 mid 个数大于 mid。如果有，说明要寻找的 h 在搜索区间的右边，反之则在左边。

func hIndex(citations []int) int {
    // 答案最多只能到数组长度
    left,right:=0,len(citations)
    var mid int
    for left<right{
        // +1 防止死循环
        mid=(left+right+1)>>1
        cnt:=0
        for _,v:=range citations{
            if v>=mid{
                cnt++
            }
        }
        if cnt>=mid{
            // 要找的答案在 [mid,right] 区间内
            left=mid
        }else{
            // 要找的答案在 [0,mid) 区间内
            right=mid-1
        }
    }
    return left
}
*/
