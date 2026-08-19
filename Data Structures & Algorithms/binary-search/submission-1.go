func search(nums []int, target int) int {
	start:=0; end:=len(nums)-1
	for start <= end{
		pivot := start + (end - start)/2
		if nums[pivot] == target {
			return pivot
		}
		if target > nums[pivot] {
			start = pivot + 1
		}else{
			end = pivot -1
		}
	}
	return -1
}	
