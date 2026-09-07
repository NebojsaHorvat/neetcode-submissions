func exist(board [][]byte, word string) bool {
	n:= len(board)
	m:= len(board[0])
	directions := [][]int{
		[]int{1,0},
		[]int{-1,0},
		[]int{0,1},
		[]int{0,-1},
	}
	var dfs func(i,j, wIndex int) bool 
	dfs = func(i,j, wIndex int) bool{
		if board[i][j] != word[wIndex]{
			return false
		}
		if wIndex == len(word)-1{
			return true
		}
		temp := board[i][j]
		board[i][j] = '#'
		resp := false
		for _, direction := range directions{
			newI := i+direction[0]
			newJ := j+direction[1]
			if newI<0 || newI > n-1 || newJ<0 || newJ>m-1 || board[newI][newJ] == '#' {
				continue
			}
			resp = resp || dfs(newI, newJ, wIndex+1)
		}
		board[i][j] = temp
		return resp
	}
	for i:=0; i<n; i++{
		for j:=0; j<m; j++{
			if dfs(i,j,0){
				return true
			}
		}
	}
	return false
}
