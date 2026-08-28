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

	for i:=0; i<n; i++ {
		for j:=0; j<m; j++{
			if grid[i][j] == 0{
				q = append(q,[3]int{i,j,0})
				visited[i*m+j] = true
			}
		}
	}
	
	iter := 0
	for iter < len(q) {
		el := q[iter]
		dist := el[2]
		for _, direction := range directions{
			newI:= el[0]+direction[0]
			newJ:= el[1]+direction[1]
			newDist := dist + 1
			if newI <0 || newI>= n || newJ<0 || newJ>=m || grid[newI][newJ] == -1 || grid[newI][newJ] == 0 || visited [newI*m+newJ] {
				continue
			}
			visited[newI*m+newJ] = true
			grid[newI][newJ] = newDist
			q = append(q, [3]int{newI, newJ, newDist})
		}
		iter++
	}
}
