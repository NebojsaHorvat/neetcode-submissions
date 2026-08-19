import "slices"

func threeSum(nums []int) [][]int {
	ret := make([][]int,0)
	slices.Sort(nums)
	fmt.Println(nums)
	for i:=0; i<len(nums)-2; i++{
		// if nums[i] == nums[i+1]{
		// 	continue
		// }
		j := i+1
		k := len(nums)-1
		for j < k {
			if nums[i] + nums[j] + nums[k] == 0{
				// remove doubles
				found := false
				for _,retSlice := range ret{
					if retSlice[0] == nums[i] && retSlice[1]==nums[j] && retSlice[2]== nums[k]{
						found = true
						break
					}
				}
				if !found{
					ret = append(ret, []int{nums[i],nums[j],nums[k]})
				}
				
			}
			if nums[i] + nums[j] + nums[k] < 0{
				j++
				continue
			}
			k--
		}
	}
	return ret
}
