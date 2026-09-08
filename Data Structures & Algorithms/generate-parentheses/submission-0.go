func generateParenthesis(n int) []string {
	ret := []string{}
	stack := []string{}
	var backtrack func(open,closed int)
	backtrack = func(open, closed int){
		if open==closed && open == n{
			ret = append(ret, strings.Join(stack,""))
		}
		if open < n{
			stack = append(stack, "(")
			backtrack(open+1, closed)
			stack = stack[:len(stack)-1]
		}
		if closed < open {
			stack = append(stack, ")")
			backtrack(open, closed+1)
			stack = stack[:len(stack)-1]
		}
	}
	backtrack(0,0)
	
	return ret
}
