func permute(nums []int) [][]int {
	ret := [][]int{}
	var dfs func([]int,[]int)
	dfs = func(permutation []int, remains []int){
		if len(remains) == 0 {
			permCopy := make([]int,len(permutation))
			copy(permCopy, permutation)
			ret = append(ret, permCopy)
			return
		}
		n := len(remains)
		for i:=0; i<n; i++{
			permutation = append(permutation, remains[i])
			remains[0], remains[i] = remains[i], remains[0]
			dfs(permutation,remains[1:])
			remains[i],remains[0] = remains[0], remains[i]
			permutation = permutation[:len(permutation)-1]
		}
	}

	dfs([]int{}, nums)
	return ret
}
