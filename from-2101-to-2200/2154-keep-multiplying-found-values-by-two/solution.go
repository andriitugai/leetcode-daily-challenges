func findFinalValue(nums []int, original int) int {
    contains := func(a int) bool {
        for _, num := range nums {
            if num == a {
                return true
            }
        }
        return false
    }
    for contains(original) {
        original *= 2
    }
    return original
}