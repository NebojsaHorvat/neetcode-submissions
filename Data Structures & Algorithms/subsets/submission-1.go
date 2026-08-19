func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	ret := [][]int{[]int{nums[0]}, []int{}}
	// fmt.Println("Ret:", ret,"\n")
	var dfs func(int)
	dfs = func (pos int){
		if pos == len(nums) {
			return
		}
		newRet := make([][]int,0, len(ret)*2)
		for i:=0; i<len(ret); i++{
				// add number from pos
				newSet := make([]int,len(ret[i]))
				copy(newSet,ret[i])
				// fmt.Println("ret[i]:",ret[i])
				newSet = append(newSet,nums[pos])
				newRet = append(newRet, newSet)
				// fmt.Println("Adding new set:", newSet)
				// copy existing
				oldSet := make([]int,len(ret[i]))
				copy(oldSet, ret[i])
				newRet = append(newRet,oldSet)
				// fmt.Println("Copy od set:",oldSet)
				// fmt.Println("NewRet:", newRet,"\n")
		}
		ret = newRet
		dfs(pos+1)
	}
	dfs(1)

	return ret
}
