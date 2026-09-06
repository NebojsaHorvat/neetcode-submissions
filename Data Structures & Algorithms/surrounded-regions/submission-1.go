func solve(board [][]byte) {
	n := len(board)
	m := len(board[0])
	directions := [][]int{
		[]int{1,0},
		[]int{-1,0},
		[]int{0,1},
		[]int{0,-1},
	}
	var dfs func(i,j int) 
	dfs = func(i,j int){
		if board[i][j] == '#' || board[i][j] == 'X' {
			return
		}
		board[i][j] = '#'
		// Return true if 'O' is on the edge
		for _,direction := range directions{
			newI := i + direction[0]
			newJ := j + direction[1]
			if newI < 0 || newI >=n || newJ <0 || newJ>=m || board[newI][newJ] == 'X' || board[newI][newJ] == '#'{
				continue
			}
			dfs(newI, newJ)
		}
	}
	for i:=0; i<n; i++{
		dfs(i,0)
		dfs(i,m-1)
	}
	for j:=0; j<m; j++{
		dfs(0,j)
		dfs(n-1,j)
	}

	for i:=0; i<n; i++{
		for j:=0; j<m; j++{
			if board[i][j] == 'O'{
				board[i][j] = 'X'
			}else if board[i][j] == '#'{
				board[i][j] = 'O'
			}
		}
	}
}
