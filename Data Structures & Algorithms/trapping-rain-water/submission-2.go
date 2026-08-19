func trap(heights []int) int {

	poolVolumes := make([][]int,0)
	end := 1
	for end < len(heights){
		// Found peek, go back and find the highest point below the peek
		if end+1 < len(heights) && heights[end+1] <= heights[end] || end == len(heights)-1{
			fmt.Println("Found Peek:",heights[end], " inx:",end)
			leftMax := 0
			leftMaxIndex := 0
			start := end-1
			for start >= 0 {
				if heights[start] > leftMax{
					leftMax = min(heights[start],heights[end])
					leftMaxIndex = start
					if heights[start] > heights[end]{
						fmt.Println("breaking on start: ",start)
						break
					}
				}
				start--
			}
			if leftMax > 0 && end-leftMaxIndex > 1{
				fmt.Println("Found leftMax:",heights[leftMaxIndex]," inx:",leftMaxIndex)
				localStart := leftMaxIndex + 1
				localPoolVolume := 0
				for localStart < end{
					localPoolVolume += leftMax - heights[localStart]
					localStart++
				}
				poolVolumes = append(poolVolumes, []int{leftMaxIndex,end,localPoolVolume})
			}
		}
		end++
	}
	fmt.Println(poolVolumes)
	// remove overpaling ranges
	poolVolumesNonOverlaping := make([][]int,0)
	for i:=0; i<len(poolVolumes); i++{
		overlaping := false
		for j:=0; j<len(poolVolumes); j++{
			if j==i{
				continue
			}
			if poolVolumes[j][0] <= poolVolumes[i][0] && poolVolumes[j][1] >= poolVolumes[i][1]{
				overlaping = true
			}
		}
		if !overlaping{
			poolVolumesNonOverlaping = append(poolVolumesNonOverlaping,poolVolumes[i])
		}
	}

	fmt.Println(poolVolumesNonOverlaping)
	poolVolume := 0
	for _, volume := range poolVolumesNonOverlaping {
		poolVolume += volume[2]
	}

	return poolVolume
}
