package main

func maximum(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func minimum(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maxArea(height []int) int {
	maxx := -99999999
	left, right := 0, len(height)-1
	for {
		maxx = maximum(maxx, (right-left)*minimum(height[left], height[right]))
		if left == right {
			break
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxx
}
