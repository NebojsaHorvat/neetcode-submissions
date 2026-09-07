func characterReplacement(s string, k int) int {
	frequencyMap := make(map [byte]int) 
	ret, start, end, maxFrequency := 0,0,0,0
	for end < len(s){
		frequencyMap[s[end]]++
		if frequencyMap[s[end]] > maxFrequency {
			maxFrequency = frequencyMap[s[end]]
		}

		for start < end && end-start-maxFrequency+1 > k{
			frequencyMap[s[start]]--
			start++
		}
		if end-start +1 > ret{
			ret = end-start+1
		}
		end++
	}
	return ret
}
