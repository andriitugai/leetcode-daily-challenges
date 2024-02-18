func longestCommonPrefix(arr1 []int, arr2 []int) int {
    prefixes := make(map[int]bool)
    for _, num := range arr1 {
        for num > 0 {
            prefixes[num] = true
            num /= 10
        }
    }

    maxPrefix := 0
    for _, num := range arr2 {
        for num > 0 {
            if prefixes[num] && num > maxPrefix {
                maxPrefix = num
                break
            }
            num /= 10
        }
    }
    width := 0
    for maxPrefix > 0 {
        maxPrefix /= 10
        width += 1
    }
    return width
}