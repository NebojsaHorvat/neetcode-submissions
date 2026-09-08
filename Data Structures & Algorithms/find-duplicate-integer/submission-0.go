func findDuplicate(nums []int) int {
    for _,num := range nums {
		if num < 0{
			num = -num
		}
		if nums[num] < 0{
			return num
		}
		nums[num] = -nums[num]
	}
	return -1
}
