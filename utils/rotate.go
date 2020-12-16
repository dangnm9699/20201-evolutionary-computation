package utils

func Exec(nums []int, k int) []int {
	if k < 0 || len(nums) == 0 {
		return nums
	}
	r := len(nums) - k%len(nums)
	nums = append(nums[r:], nums[:r]...)
	return nums
}
