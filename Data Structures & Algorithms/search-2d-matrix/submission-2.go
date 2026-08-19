func searchMatrix(matrix [][]int, target int) bool {
	start := 0; end := len(matrix)-1
	targetRow := 0
	for start <= end{
		p := start + (end-start)/2
		if p == len(matrix)-1{
			targetRow = p
			break
		}
		if target >= matrix[p][0] && target < matrix[p+1][0]{
			targetRow = p
			break
		}
		if target > matrix[p][0]{
			start = p + 1
			continue
		}
		end = p -1
	}
	start = 0; end = len(matrix[0])-1
	for start <= end{
		p := start + (end-start)/2
		if target == matrix[targetRow][p]{
			return true
		}
		if target > matrix[targetRow][p]{
			start = p +1
			continue
		}
		end = p-1
	}
	
	return false
}
