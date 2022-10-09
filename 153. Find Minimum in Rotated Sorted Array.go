package main

func findMin(nums []int) int {
	left := 0
	length := len(nums) - 1
	right := length
	var mid int
	for left < right {
		mid = (left + right) / 2
		if left+1 == right {
			if nums[left] < nums[right] {
				return nums[left]
			} else {
				return nums[right]
			}
		}
		if (nums[left] < nums[mid] && nums[mid] < nums[right]) || (nums[left] > nums[mid]) {
			right = mid
		} else {
			left = mid
		}
	}
	return nums[left]
}
