import "slices"

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	ret := [][]int{}
	n := len(nums)
	for i:=0; i<n; i++ {
		if i>0 && nums[i] == nums[i-1]{
			continue
		}
		start := i+1
		end := n-1
		for start < end {
			currentSum := nums[i] + nums[start] + nums[end]
			if currentSum == 0{
				ret = append(ret, []int{nums[i],nums[start],nums[end]})
				start++
				end--
				for start<end && nums[start] == nums[start-1]{
					start++
				}
				for start<end && nums[end] == nums[end+1]{
					end--
				}
 			}else if currentSum < 0 {
				start ++
			}else {
				end--
			}
		}
	}
	return ret
}
