func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var ret [][]int

	var dfs func(int, []int)
	dfs = func(i int, subSet []int){
		if i == len(nums){
			ret = append(ret, append([]int{},subSet...))
			return
		}	
		
		subSet = append(subSet,nums[i])
		dfs(i+1, subSet)
		subSet = subSet[:len(subSet)-1]

		for i+1 < len(nums) && nums[i+1] == nums[i]{
			i++
		}
		dfs(i+1, subSet)
	}

	dfs(0,[]int{})
	return ret
}
