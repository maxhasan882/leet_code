package main

func productExceptSelf(nums []int) []int {
	allProduct := 1
	zeroCount := 0
	for i := range nums {
		if nums[i] != 0 {
			allProduct *= nums[i]
		} else {
			zeroCount += 1
		}
	}
	var result []int
	for i := range nums {
		if zeroCount > 1 {
			result = append(result, 0)
			continue
		}
		if zeroCount == 1 && nums[i] == 0 {
			result = append(result, allProduct)
		} else if zeroCount == 1 && nums[i] != 0 {
			result = append(result, 0)
		} else {
			result = append(result, allProduct/nums[i])
		}
	}
	return result
}
