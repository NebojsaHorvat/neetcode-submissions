func maxArea(heights []int) int {
	max := 0
	start:=0; end:= len(heights)-1
	for start < end {
		if (end-start)* min(heights[end],heights[start]) > max {
			max = (end-start)* min(heights[end],heights[start])
		}
		if heights[start] < heights[end]{
			start++
			continue
		}
		end--
	}
	return max
}
