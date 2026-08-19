func isValid(s string) bool {
    stack := make([]byte,0)
	for i:=0; i<len(s); i++ {
		char := s[i]
		if char == '[' || char == '{' || char == '('{
			stack = append(stack,char)
		}else if char == ']' && len(stack) > 0 && stack[len(stack)-1] == '['{
			stack =stack[:len(stack)-1]
		}else if char == '}' && len(stack) > 0 && stack[len(stack)-1] == '{'{
			stack =stack[:len(stack)-1]
		}else if char == ')' && len(stack) > 0 && stack[len(stack)-1] == '('{
			stack =stack[:len(stack)-1]
		}else{
			return false
		}
	}
	if len(stack) == 0 {
		return true
	}else {
		return false
	}
}
