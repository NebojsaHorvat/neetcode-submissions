func lengthOfLongestSubstring(s string) int {
	first :=0; last := 0
	max := 0
	var inWindow [128]bool
	for first < len(s){
		// add char in window
		if !inWindow[s[first]]{
			inWindow[s[first]] = true
			if first-last + 1 > max{
				max = first - last + 1
			}
			first ++
			continue
		}
		// remove char from window and all chars with it
		for s[last] != s[first]{
			inWindow[s[last]] = false
			last++
		}
		last++
		first++
	}
	return max
}
