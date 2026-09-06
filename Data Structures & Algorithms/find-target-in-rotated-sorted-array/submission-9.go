func search(nums []int, target int) int {
	n := len(nums)
	start := 0; end := n-1

	for start <= end {
		p := start + (end-start)/2
		fmt.Println("p:",p)
		if nums[p] == target {
			return p
		}
		// Clasic case
		if nums[p] < nums[end] {
			if target > nums[p] && target <= nums[end]{
				start = p+1
				fmt.Println("NewStart:",start)
			}else {
				end = p-1
				fmt.Println("NewEnd:",end)
			}
		}else{
			if target > nums[p] || target <= nums[end]{
				start = p+1
				fmt.Println("NewStart:",start)
			}else {
				end = p-1
				fmt.Println("NewEnd:",end)
			}
		}
	}
	return -1
}
