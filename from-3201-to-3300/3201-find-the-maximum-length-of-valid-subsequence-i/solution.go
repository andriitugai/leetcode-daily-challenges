func maximumLength(nums []int) int {
    n := len(nums)
    alter := 1
    var odd, even int = 0, 0
    if nums[0] % 2 == 0 {
        even = 1
    } else {
        odd = 1
    }

    prevOdd := (nums[0] % 2 == 1)
    for i := 1; i < n; i ++ {
        currEven := (nums[i] % 2 == 0)
        if currEven {
            if prevOdd {
                alter += 1
                prevOdd = false
            }
             even += 1
        } else {
            if !prevOdd {
                alter += 1
                prevOdd = true
            }
            odd += 1
        }
    }
    return max(alter, max(odd, even))
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}