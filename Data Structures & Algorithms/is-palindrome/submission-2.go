func isPalindrome(s string) bool {
	start := 0
	end := len(s)-1
	s = strings.ToLower(s)
	for start < end{
		if !isDigit(s[start]) && !isLetter(s[start]){
			start++
			// fmt.Println("Skipping ",string(s[start]))
			continue
		}
		if !isDigit(s[end]) && !isLetter(s[end]){
			end--
			// fmt.Println("Skipping ",string(s[end]))
			continue
		}
		// fmt.Println("Comparing ", s[start], s[end])
		if s[start] != s[end]{
			return false
		}
		start++
		end--
	}
	return true
}

func isDigit(c byte)bool{
	return c >= '0' && c<= '9'
}
func isLetter(c byte)bool{
	return c >= 'a' && c<= 'z'
}