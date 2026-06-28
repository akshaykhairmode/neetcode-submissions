import (
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	hashMap := map[string][]string{}
	
	for i , str := range strs {
		strBytes := []byte(str)
		slices.Sort(strBytes)
		hashMap[string(strBytes)] = append(hashMap[string(strBytes)],strs[i])
	}

	out := [][]string{}

	for _ , v := range hashMap {
		out = append(out,v)
	}

	return out

}
