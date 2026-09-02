func islandsAndTreasure(grid [][]int) {
	directions := [][2]int{
		[2]int{0,-1},
		[2]int{0,1},
		[2]int{1,0},
		[2]int{-1,0},
	}
	n:= len(grid)
	m:= len(grid[0])
	visited := make([]bool,n*m)
	q := [][3]int{}
	
	for i:=0; i<n; i++{
		for j:=0; j<m; j++{
			if grid[i][j] == 0{
				q = append(q, [3]int{i,j,0})
			}
		}
	}
	index := 0
	for index < len(q){
		el := q[index]
		index++
		i:= el[0]; j:= el[1]
		dist := el[2]
		if grid[i][j] != 0 && dist < grid[i][j]  {
			grid[i][j] = dist
		}
		for _, direction := range directions{
			newI := i+direction[0]
			newJ := j+direction[1]
			if newI < 0 || newI >= n || newJ<0 || newJ >=m || visited[newI*m+newJ] || grid[newI][newJ] == -1 || grid[newI][newJ] == 0{
				continue
			}
			visited[newI*m+newJ] = true
			q = append(q, [3]int{newI,newJ,dist+1})
		}
	}
	
}
