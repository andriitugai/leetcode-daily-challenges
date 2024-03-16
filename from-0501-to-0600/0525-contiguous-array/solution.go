func findMaxLength(nums []int) int {
    zeros, ones := 0, 0
    result := 0
    diffIndex := make(map[int]int)

    for i, num := range nums {
        if num == 0 {
            zeros += 1
        } else {
            ones += 1
        }

        if _, ok := diffIndex[ones - zeros]; !ok {
            diffIndex[ones - zeros] = i
        }

        if ones == zeros {
            result = ones + zeros
        } else {
            idx := diffIndex[ones - zeros]
            if i - idx > result {
                result = i - idx
            }
        }
    }
    return result
}