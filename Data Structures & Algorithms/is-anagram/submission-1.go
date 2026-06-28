func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	//a-z (26)
	arr := make([]byte, 26)

	//v is a rune.
	for _ , v := range s {
		arr[v-97]++
	}

	for _ , v := range t {
		if arr[v-97] > 0 {
			arr[v-97]--
		}
	}

	for _ , v := range arr {
		if v > 0 {
			return false
		}
	}

	return true
	
}
