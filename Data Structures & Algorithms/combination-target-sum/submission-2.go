func combinationSum(nums []int, target int) [][]int {
	ret:= [][]int{}
	sort.Ints(nums)
	var dfs func (int,[]int,int)
	dfs = func (i int, subSet []int, sum int){
		if sum == target{
			subSetNew := make([]int,len(subSet))
			copy(subSetNew, subSet)
			ret = append(ret, subSetNew)
			return
		}else if sum > target{
			return
		}
		for j:=i; j<len(nums); j++{
			num := nums[j]
			subSet = append(subSet, num)
			dfs(j, subSet, sum + num)
			subSet = subSet[:len(subSet)-1]
		}
	}
	dfs(0,[]int{},0)
	return ret
}
