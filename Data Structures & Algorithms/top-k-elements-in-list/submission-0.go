func topKFrequent(nums []int, k int) []int {

	//Unoptimized solution
	freqMap := map[int]int{}
	for _ , v := range nums {
		freqMap[v]++
	}

	out := []int{}
	for range k {
		maxValue := 0
		maxNum := 0
		for num , count := range freqMap {
			if count > maxValue {
				maxNum = num
				maxValue = count
			}
			fmt.Println(num,count,maxNum,maxValue)
		}
		out = append(out,maxNum)
		freqMap[maxNum] = 0
	}

	return out

}
