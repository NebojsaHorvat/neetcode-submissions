func permute(nums []int) [][]int {
	retPerms := [][]int{{}}
	for _, num := range nums{
		newPerms := [][]int{}
		for _, perm := range retPerms{
			for i:=0; i<= len(perm); i++{
				newPerm := make([]int,len(perm), len(perm)+1)
				copy(newPerm,perm)
				newPerm = append(newPerm[:i],append([]int{num},newPerm[i:]...)...)
				newPerms = append(newPerms, newPerm)
			}
		}
		retPerms = newPerms
	}
	return retPerms
}
