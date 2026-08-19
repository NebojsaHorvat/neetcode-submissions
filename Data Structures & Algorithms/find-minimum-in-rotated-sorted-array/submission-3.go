func findMin(nums []int) int {
	start := 0; end := len(nums)-1
	min := nums[0]
	for start <= end {
		p := start + (end-start)/2
		fmt.Println("p:",p)
		if nums[p] < min {
			min = nums[p]
		}
		if nums[start] < min{
			min = nums[start]
		}
		if nums[start] <= nums[p]{
			start = p+1
			fmt.Println("s:",start)
		}else{
			end = p-1
			fmt.Println("e:",end)
		}
	}
	return min
}
