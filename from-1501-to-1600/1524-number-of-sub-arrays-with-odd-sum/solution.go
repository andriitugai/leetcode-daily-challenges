func numOfSubarrays(arr []int) int {
    MOD := 1000000007
    oddity := make([]int, len(arr))
    curr := 0
    for i, num := range arr {
        curr = (curr + num) % 2
        oddity[i] = curr
    }

    odd, even := 0, 0
    result := 0

    for _, v := range oddity {
        if v == 0 {
            even += 1
            result += odd
        } else {
            odd += 1
            result += even + 1
        }
        result %= MOD
    }
    return result
}