import "slices"

func subsetsWithDup(nums []int) [][]int {
	slices.Sort(nums)
	var ret [][]int

	var dfs func(int, []int)
	dfs = func(i int, subSet []int){
		if i > len(nums)-1 {
			var retSub []int
			retSub = append(retSub, subSet...)
			ret = append(ret, retSub)
			return 
		}
		j:=i
		for j<len(nums) && nums[j]==nums[i]{
			j++
		}
		dfs(j, subSet)
		subSet = append(subSet, nums[i])
		dfs(i+1, subSet)
		subSet = subSet[:len(subSet)-1]
	}
	dfs(0,[]int{})
	return ret
}
