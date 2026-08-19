func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	ret := [][]int{}

	var dfs func(int, []int)
	dfs = func(i int, subSet []int){
		// subSetCopy := make([]int,len(subSet))
		// copy(subSetCopy, subSet)
		// ret = append(ret, subSetCopy)
		ret = append(ret, append([]int{},subSet...))
		for j:=i; j<len(nums); j++{
			if j > i && nums[j-1] == nums[j]{
				continue
			}
			// dfs(j+1, subSet)
			subSet = append(subSet,nums[j])
			dfs(j+1, subSet)
			subSet = subSet[:len(subSet)-1]
		}
	}

	dfs(0,[]int{})
	return ret
}
