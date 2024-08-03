func canBeEqual(target []int, arr []int) bool {
    if len(arr) != len(target) {
        return false
    }
    sort.Ints(arr)
    sort.Ints(target)
    for i := 0; i < len(arr); i ++ {
        if arr[i] != target[i] {
            return false
        }
    }
    return true
}