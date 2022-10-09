package main

func maximumItem(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func minimumItem(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maXItemFromArray(nums []int) int {
	max := -999999999
	for _, num := range nums {
		max = maximumItem(max, num)
	}
	return max
}

func maxProduct(nums []int) int {
	corrMin, corrMax := 1, 1
	result := maXItemFromArray(nums)
	for _, n := range nums {
		tmp := n * corrMax
		corrMax = maximumItem(n, maximumItem(tmp, n*corrMin))
		corrMin = minimumItem(n, minimumItem(tmp, n*corrMin))
		result = maximumItem(corrMax, result)
	}
	return result
}
