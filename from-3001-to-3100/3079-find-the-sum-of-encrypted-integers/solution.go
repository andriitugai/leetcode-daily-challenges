func sumOfEncryptedInt(nums []int) int {
    max := func(a, b int) int {
        if a > b {
            return a
        }
        return b
    }
    encrypt := func(num int) int {
        len, dig, res := 0, 0, 0
        for num > 0 {
            dig = max(dig, num % 10)
            num /= 10
            len += 1
        }
        for len > 0 {
            res = res * 10 + dig
            len -= 1
        }
        return res
    }
    result := 0
    for _, num := range nums {
        result += encrypt(num)
    }
    return result
}