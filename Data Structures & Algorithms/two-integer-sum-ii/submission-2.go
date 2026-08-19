func twoSum(nums []int, target int) []int {
	
	i:=0
	j:=1
	maxNotFound := true
	for target != nums[i]+nums[j]{
		if maxNotFound && j<len(nums)-1 && nums[i]+nums[j] < target{
			j++
			fmt.Println("j:",j)
			continue
		}
		maxNotFound = false
		if nums[i+1]+nums[j] <= target{
			i++
			fmt.Println("i",i)
			continue
		}
		j--
		fmt.Println("j:",j)
	}
	return []int{i+1,j+1}
	
}
