func diagonalPrime(nums [][]int) int {
    checkPrime := func(a int) bool {
        if a < 2 {
            return false
        }
        for i := 2; i * i <= a; i ++ {
            if a % i == 0 {
                return false
            }
        }
        return true
    }
    n := len(nums)
    result := 0
    for i := 0; i < n; i ++ {
        if checkPrime(nums[i][i]) {
            if nums[i][i] > result {
                result = nums[i][i]
            }
        }
        if checkPrime(nums[i][n - i - 1]) {
            if nums[i][n - i - 1] > result {
                result = nums[i][n - i - 1]
            }
        }
    }
    return result
}