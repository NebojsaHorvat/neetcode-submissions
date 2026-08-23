func maxAreaOfIsland(grid [][]int) int {
    visited := make(map[int]bool)
	max := 0
	currentMax := 0
	var search func(i,j int)
	search = func(i,j int){
		key := i*len(grid[0])+j
		if i<0 || i>=len(grid) || j<0 || j>=len(grid[0]) || visited[key]{
			return 
		}
		if visited[key] || grid[i][j]==0{
			return 
		}
		visited[key] = true
		currentMax++
		search(i+1, j)
		search(i-1, j)
		search(i, j+1)
		search(i, j-1)
	}
	for i:=0; i<len(grid); i++{
		for j:=0; j<len(grid[0]); j++{
			key := i*len(grid[0])+j
			if visited[key] || grid[i][j] == 0{
				continue
			}
			currentMax = 0
			search(i,j)
			if currentMax > max{
				max = currentMax
			}

		}
	}
	return max
}
