func canFinish(numCourses int, prerequisites [][]int) bool {
    prereqMap := make(map[int][]int)
	for _, prereq := range prerequisites{
		pre := prereq[1]
		course := prereq[0]
		prereqMap[course] = append(prereqMap[course], pre)
	}

	var dfs func(int, []int) bool
	dfs = func(course int, visited []int) bool{
		if visited[course] == 2{
			return false
		}
		if visited[course] == 1{
			return true
		}
		visited[course]++
		for _, prereq := range prereqMap[course]{
			if dfs(prereq,visited){
				return true
			}
		}
		visited[course] = 2
		return false
	}

	visited := make([]int, numCourses)
	for i:=0; i<numCourses; i++{
		if dfs(i, visited){
			return false
		}
	}
	return true
}
