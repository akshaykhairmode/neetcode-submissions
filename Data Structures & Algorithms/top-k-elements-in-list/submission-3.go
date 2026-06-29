func topKFrequent(nums []int, k int) []int {

	//Optimized solution
	// index - [0, 1, 2, 3 , 4 , 5 , 6] - How many times a value came. (len(nums))
	// value - [[], [1], [2], [3, 4] ]

	freqTable := func() map[int]int {
		o := map[int]int{}
		for _ , v := range nums {
			o[v]++
		}
		return o
	}()

	countTable := make([][]int,len(nums)+1) // len(nums)

	for num , count := range freqTable {
		countTable[count] = append(countTable[count],num)
	}

	out := []int{}
	for len(out) < k {
		for i := len(countTable) - 1; i > 0; i-- {
			if len(countTable[i]) > 0 {
				out = append(out,countTable[i]...)

				if len(out) >= k {
					return out[0:k]
				}
			}
		}
	}

	return out[0:k]
	
}
