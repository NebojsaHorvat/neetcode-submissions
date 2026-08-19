func groupAnagrams(strs []string) [][]string {
	anagrams  := make(map [[26]int][]string)
	for _, str := range strs{
		histogram := [26]int{}
		for _,char := range str{
			histogram[char-'a']++
		}
		anagrams[histogram] = append(anagrams[histogram],str)
	}
	ret := [][]string{}
	for _,anagram := range anagrams{
		ret = append(ret,anagram)
	}
	return ret
}
