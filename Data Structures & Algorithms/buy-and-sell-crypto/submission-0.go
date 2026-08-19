func maxProfit(prices []int) int {
	min := prices[0]
	diff := 0
	for i := 1; i<len(prices); i++ {
		if prices[i]- min > diff{
			diff = prices[i]-min
		}
		if prices[i] < min{
			min = prices[i]
		}
	}
	return diff
}
