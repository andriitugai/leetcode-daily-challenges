func zeroFilledSubarray(nums []int) int64 {
    var result int64 = 0
    var cont int64 = 0
    for _, num := range(nums) {
        if num == 0 {
            cont += 1
            result += cont
        } else {
            cont = 0
        }
    }
    return result
}