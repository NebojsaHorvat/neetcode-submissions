func findMin(nums []int) int {
	start := 0; end := len(nums)-1
	min := nums[0]

	for start <= end{
		p := start + (end-start)/2
		if nums[p] <= min{
			min = nums[p]
		}
		if nums[end] < nums[p]{
			start = p+1
		}else{
			end = p-1
		}
	}
	return min
}
