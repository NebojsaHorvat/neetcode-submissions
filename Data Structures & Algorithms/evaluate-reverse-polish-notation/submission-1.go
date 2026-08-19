func evalRPN(tokens []string) int {
	if len(tokens) == 0{
		return 0
	}
	num, _ := strconv.Atoi(tokens[0])
	nums := []int{num}

	for _, token := range tokens[1:]{
		if token == "+"{
			num := nums[len(nums)-1]
			nums = nums[:len(nums)-1]
			nums[len(nums)-1] += num
		}else if token == "-"{
			num := nums[len(nums)-1]
			nums = nums[:len(nums)-1]
			nums[len(nums)-1] -= num
		}else if token == "*"{
			num := nums[len(nums)-1]
			nums = nums[:len(nums)-1]
			nums[len(nums)-1] *= num
		}else if token == "/"{
			num := nums[len(nums)-1]
			nums = nums[:len(nums)-1]
			nums[len(nums)-1] /= num
		}else {
			num, _ := strconv.Atoi(token)
			nums = append(nums,num)
		}
	}

	return nums[0]
}
