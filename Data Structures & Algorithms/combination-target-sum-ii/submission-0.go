func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	ret := make([][]int,0)
	var dfs func(int, []int, int)
	dfs = func(elementPos int,subSet []int, currentSum int){
		if currentSum == target{ 
			subSetCopy := make([]int,len(subSet))
			copy(subSetCopy, subSet)
			ret = append(ret, subSetCopy)
		}
		for j:=elementPos; j<len(candidates); j++{
			if currentSum + candidates[j] > target{
				continue
			}
			subSet = append(subSet, candidates[j])
			dfs(j+1, subSet, currentSum+candidates[j])
			subSet = subSet[:len(subSet)-1]
			for  j+1<len(candidates) && candidates[j] == candidates[j+1]{
				j++
			}
		}
	}

	dfs(0,[]int{},0)
	return ret
}
