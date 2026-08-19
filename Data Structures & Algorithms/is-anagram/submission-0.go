import "slices"
func isAnagram(s string, t string) bool {
	appearedS := make([]int,26)
	appearedT := make([]int,26)
	for _, letter := range s{
		letter = letter - 'a'
		appearedS[letter] += 1
	}
	for _, letter := range t{
		letter = letter - 'a'
		appearedT[letter] += 1
	}
	return slices.Equal(appearedS, appearedT)
}
