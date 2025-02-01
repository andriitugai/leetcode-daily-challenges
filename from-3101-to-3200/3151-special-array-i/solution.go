func isArraySpecial(nums []int) bool {
    prev := nums[0] & 1
    var curr int
    for _, num := range(nums[1:]) {
        curr = num & 1
        if prev == curr {
            return false
        }
        prev = curr
    }
    return true
}