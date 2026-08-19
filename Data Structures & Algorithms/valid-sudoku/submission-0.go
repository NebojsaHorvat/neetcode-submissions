func isValidSudoku(board [][]byte) bool {
	exists := make(map [string]bool)
	for i:= 0; i<9; i++ {
		for j:= 0; j<9; j++{
			if board[i][j] == '.' {
				continue
			}
			// row
			rowVal := "r" + string(i + 48) + "-" + string(board[i][j])
			if exists[rowVal]{
				return false
			}
			// fmt.Println("Adding " + rowVal)
			exists[rowVal] = true
			// cal
			calVal := "c" + string(j + 48) + "-" + string(board[i][j])
			if exists[calVal] {
				return false
			}
			// fmt.Println("Adding " + calVal)
			exists[calVal] = true
			// cube
			cube := (i/3)*3 + j/3
			cubeVal := "cube"+string(cube+48)+string(board[i][j])
			if exists[cubeVal]{
				return false
			}
			exists[cubeVal] = true
			// fmt.Println("Cube " + string(cube + 48)) 
			// fmt.Println("-----")
		}
	}
	return true
}
