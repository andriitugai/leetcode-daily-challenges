func countBeautifulPairs(nums []int) int {
    gcd := func(a, b int) int {
        for b > 0 {
            a, b = b, a % b
        }
        return a
    }
    result := 0
    for i := 0; i < len(nums)-1; i ++ {
        for j := i + 1; j < len(nums); j ++ {
            if gcd(int(strconv.Itoa(nums[i])[0] - '0'), nums[j] % 10 ) == 1 {
                fmt.Println(nums[i], nums[j])
                result += 1
            }
        }
    }
    return result
}