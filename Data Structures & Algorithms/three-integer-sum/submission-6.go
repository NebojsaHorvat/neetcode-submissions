import "slices"

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	n := len(nums)
	ret := [][]int{}
	for i:=0; i<n; i++{
		if i>0 && nums[i] == nums[i-1]{
			continue
		}
		for j:=i+1; j<n; j++{
			if j>i+1 && nums[j] == nums[j-1]{
				continue
			}
			for k:=j+1; k<n; k++{
				if k>j+1 && nums[k] == nums[k-1]{
					continue
				}
				if nums[i]+nums[j]+nums[k] > 0{
					break
				}
				if nums[i]+nums[j]+nums[k] == 0 {
					ret = append(ret, []int{nums[i],nums[j],nums[k]})
				}
			}
		}
	}
	return ret
}
