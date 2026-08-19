func productExceptSelf(nums []int) []int {
	prod := 1
	zeros := 0
	for _, num := range nums{
		if num == 0 {
			zeros++
			continue
		}
		prod *= num
	}
	ret := make([]int,len(nums))
	if zeros > 1 {
		return ret
	}
	for i, num := range nums{
		if zeros == 1{
			if num == 0 {
				ret[i] = prod
			}
		}else{
			ret[i] = prod / num
		}
	}
	return ret
}
