func majorityElement(nums []int) []int {
    // Booyer - Moor majority algorithm
    m1, m2 := 0, 0
    c1, c2 := 0, 0

    for _, num := range nums {
        if num == m1 {
            c1 += 1
        } else if num == m2 {
            c2 += 1
        } else if c1 == 0 {
            m1 = num
            c1 = 1
        } else if c2 == 0 {
            m2 = num
            c2 = 1
        } else {
            c1 -= 1
            c2 -= 1
        }
    }

    c1, c2 = 0, 0
    for _, num := range nums {
        if num == m1 {
            c1 += 1
        } else if num == m2 {
            c2 += 1
        }
    }

    result := []int{}
    if c1 > len(nums) / 3 {
        result = append(result, m1)
    }
    if c2 > len(nums) / 3 {
        result = append(result, m2)
    }

    return result
}