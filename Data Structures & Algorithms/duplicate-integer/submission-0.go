func hasDuplicate(nums []int) bool {
    seen := map[int]int{}
    for _ , v := range nums {
        seen[v]++
        if seen[v] > 1 {
            return true
        }
    }
    return false
}
