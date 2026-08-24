func numIslands(grid [][]byte) int {
	islands := 0
	directions := [][2]int{
		[2]int{1,0},
		[2]int{-1,0},
		[2]int{0,1},
		[2]int{0,-1},
	}
	var bfs func(i,j int)
	bfs = func(i,j int){
		queue := [][2]int{[2]int{i,j}}
		for len(queue) > 0{
			el := queue[len(queue)-1]
			queue = queue[:len(queue)-1]
			grid[el[0]][el[1]] = '0'
			for _, direction := range directions{
				newI := el[0] + direction[0]
				newJ := el[1] + direction[1]
				if newI<0 || newI >= len(grid) || newJ<0 || newJ>=len(grid[0]) || grid[newI][newJ] == '0'{
					continue
				}
				queue = append(queue, [2]int{newI,newJ})
			}
		}
	}

	for i:=0; i<len(grid); i++{
		for j:=0; j<len(grid[0]); j++{
			if grid[i][j] == '0'{
				continue
			}
			bfs (i,j)
			islands++
		}
	}
	return islands
}


