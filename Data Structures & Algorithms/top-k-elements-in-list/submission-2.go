func topKFrequent(nums []int, k int) []int {

	histogram := make(map [int]int)
	for _,num := range nums{
		histogram[num]++
	}
	arr := make([][2]int, 0 ,len(histogram))
	for k,v := range histogram {
		arr = append(arr, [2]int{k,v})
	}
	sort.Slice(arr, func(i,j int) bool{
		return arr[i][1] > arr[j][1]
	})
	ret := make([]int,k)
	for i:=0; i<k; i++ {
		ret[i]=arr[i][0]
	}
	return ret
	}
