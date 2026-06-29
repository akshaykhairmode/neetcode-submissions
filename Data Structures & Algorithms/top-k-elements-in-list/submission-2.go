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

	//fmt.Println(freqTable)
	// fmt.Println(len(countTable))
	// fmt.Println(countTable)

	out := []int{}
	// fmt.Println(len(out),k)
	for len(out) < k {
		// fmt.Println("26",countTable,len(out))
		for i := len(countTable) - 1; i > 0; i-- { //6 > 0 true
			// fmt.Println("28",len(out),i)
			if len(countTable[i]) > 0 {
				// fmt.Println("30",countTable[i],len(out))
				out = append(out,countTable[i]...)
				// fmt.Println(len(out))

				if len(out) >= k {
					return out[0:k]
				}
			}
			//fmt.Println(i)
		}
	}

	return out[0:k]
	
}
