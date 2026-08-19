func combinationSum(nums []int, target int) [][]int {
	ret:= [][]int{}
	sort.Ints(nums)
	var dfs func (int,[]int,int)
	dfs = func (i int, subSet []int, sum int){
		if sum == target{
			ret = append(ret, subSet)
			return
		}else if sum > target{
			return
		}
		for j:=i; j<len(nums); j++{
			num := nums[j]
			subSetNew := make([]int,len(subSet),len(subSet)+1)
			copy(subSetNew,subSet)
			subSetNew = append(subSetNew, num)
			sumNew := sum + num
			dfs(j, subSetNew, sumNew)
		}
	}
	dfs(0,[]int{},0)
	return ret
}
