package array_and_strings

import (
	"testing"
)

func Test_productExceptSelf(t *testing.T) {
	var nums = []int{1, 2, 3, 4}
	answer := productExceptSelf(nums)
	println(answer)
	nums = []int{-1, 1, 0, -3, 3}
	answer = productExceptSelf(nums)
	println(answer)
}
