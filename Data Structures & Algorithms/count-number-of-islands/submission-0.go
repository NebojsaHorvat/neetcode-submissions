func numIslands(grid [][]byte) int {
    visited := make(map[int]bool)
	islands := 0

	var search func(int, int)
	search = func(i,j int){
		key := i*len(grid[0])+j
		if i<0 || i>=len(grid) || j<0 || j>=len(grid[0]){
			return
		}
		if visited[key]{
			return
		}
		visited[key]=true

		if grid[i][j] == '0'{
			return
		}

		search(i+1,j)
		search(i-1,j)
		search(i,j+1)
		search(i,j-1)
	}

	for i:=0; i<len(grid); i++ {
		for j:=0; j<len(grid[0]); j++{
			key := i*len(grid[0])+j
			if visited[key]{
				continue
			}
			if grid[i][j] == '0'{
				visited[key] = true
				continue
			}
			islands++
			search(i,j)
		}
	}
	return islands
}
