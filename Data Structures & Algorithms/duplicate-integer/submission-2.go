func hasDuplicate(nums []int) bool {
    hashMap := make(map[int]int, len(nums))
    for i, v := range nums{
        if _, ok := hashMap[v]; ok{
            return true
        }
        hashMap[v] = i
    }
    return false
}
