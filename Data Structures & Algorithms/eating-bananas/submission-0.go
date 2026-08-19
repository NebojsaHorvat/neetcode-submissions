func minEatingSpeed(piles []int, h int) int {
	max := piles[0]
	for _, pile := range piles{
		if pile > max{
			max = pile
		}
	}
	start := 1; end := max
	globalK := max
	for start <= end {
		p := start + (end-start)/2
		// Calculate eating time
		eatingTime := 0
		for _, pile := range piles{
			eatingTime += pile/p
			if pile % p != 0{
				eatingTime++
			}
		}
		// -----
		if eatingTime > h{
			start = p+1
			continue
		}
		// eatingTime < h
		if p < globalK{
			globalK = p
		}
		end = p-1
	}
	return globalK
}
