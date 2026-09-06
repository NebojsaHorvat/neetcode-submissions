func search(nums []int, target int) int {
	n := len(nums)
	start := 0; end := n-1

	for start <= end {
		p := start + (end-start)/2
		if nums[p] == target {
			return p
		}
		// Clasic case
		if nums[p] < nums[end] {
			if target > nums[p] && target <= nums[end]{
				start = p+1
			}else {
				end = p-1
			}
		}else{
			if target > nums[p] || target <= nums[end]{
				start = p+1
			}else {
				end = p-1
			}
		}
	}
	return -1
}
