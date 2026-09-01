func combinationSum(nums []int, target int) [][]int {
	ret:= [][]int{}
	sort.Ints(nums)
	var dfs func(index int, subSet []int, sum int)
	dfs = func(index int, subSet []int, sum int){
		if index >= len(nums) || sum > target{
			return
		}
		if sum == target{
			var retSubSet []int
			retSubSet = append(retSubSet, subSet...)
			ret = append(ret, retSubSet)
			return
		}
		dfs(index+1, subSet, sum )
		subSet = append(subSet, nums[index])
		dfs(index, subSet, sum+nums[index])
		subSet = subSet[:len(subSet)-1]
	}
	dfs(0,[]int{},0)
	return ret
}
