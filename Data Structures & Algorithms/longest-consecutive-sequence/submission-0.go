func longestConsecutive(nums []int) int {
    numSet := make(map [int]bool)
    for _, num := range nums {
        numSet[num]=true
    }
    longest := 0
    for _, num := range nums{
        if numSet[num-1] == false{
            locLongest := 1
            locNum := num
            for numSet[locNum+1] {
                locLongest++
                locNum++
            }
            if locLongest > longest{
                longest = locLongest
            }
        }
    }
    return longest
}
