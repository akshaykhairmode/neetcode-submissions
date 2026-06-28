func groupAnagrams(strs []string) [][]string {

	hashMap := map[[26]int][]string{}

	for _ , str := range strs {
		charArr := [26]int{}
		for _ , v := range str {
			charArr[v - 'a']++
		}

		hashMap[charArr] = append(hashMap[charArr],str)
	}

	out := [][]string{}
	for _ , v := range hashMap {
		out = append(out,v)
	}

	return out
}
