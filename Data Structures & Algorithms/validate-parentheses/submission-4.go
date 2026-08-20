func isValid(s string) bool {
    stack := make([]byte,0)
	for i:=0; i<len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{'{
			stack = append(stack,s[i])
			continue
		}
		if len(stack) == 0 {
			return false
		}
		lastElement := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if s[i] == ')' && lastElement == '('{
			continue
		}else if s[i] == ']' && lastElement == '[' {
			continue
		}else if s[i] == '}' && lastElement == '{' {
			continue
		}
		return false
	}
	if len(stack) == 0 {
		return true
	}else {
		return false
	}
}
