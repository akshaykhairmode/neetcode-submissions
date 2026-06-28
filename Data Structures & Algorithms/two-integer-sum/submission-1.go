func twoSum(nums []int, target int) []int {
	seen := map[int]int{}
	for i , v := range nums {
		seen[v] = i
	}

	for oi , v := range nums {
		needed := target - v 
		if ii , ok := seen[needed]; ok && oi != ii {
			return []int{oi,ii}
		}
	}

	return []int{}
}
