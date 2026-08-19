func lengthOfLongestSubstring(s string) int {
	first :=0; last := 0
	longest := 0
	previousDeletedPos := 0
	// map rune -> possition
	unique := make(map [rune]int)
	for pos, char := range s{
		// found duplicate
		if val, ok := unique[char]; ok{
			unique[char] = pos
			first = val+1
			last = pos
			if first > last {
				last = first
			}
			// fmt.Println("Found duplicate: ",string(rune(char)), " first position: ",first, " last position: ",last)
			for i:= val-1; i>=previousDeletedPos; i--{
				if rune(s[i]) == char{
					continue
				}
				// fmt.Println("Deleting: ",string(rune(s[i])))
				delete(unique,rune(s[i]))
			}
			previousDeletedPos = val+1
			
		}else{
			unique[char] = pos
			last = pos
			// fmt.Println("Adding: ",string(rune(char)))
			if last-first+1 > longest {
				longest = last-first+1
				// fmt.Println("Found new possition: ", longest, " - first: ",first," - last:",last)
			}
		}
	}
	return longest
}
