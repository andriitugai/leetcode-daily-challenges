func countPairs(nums []int) int {
    result := 0
    for i := 0; i < len(nums); i ++ {
        for j := i + 1; j < len(nums); j ++ {
            if almostEqual(nums[i], nums[j]) {
                result += 1
            }
        }
    }
    return result
}

func almostEqual(a, b int) bool {
    as := strconv.Itoa(a)
    bs := strconv.Itoa(b)

    for len(as) < len(bs) {
        as = "0" + as
    }
    for len(bs) < len(as) {
        bs = "0" + bs
    }

    diff := 0
    missMatchIdx1, missMatchIdx2 := -1, -1
    for i := 0; i < len(as); i ++ {
        if as[i] != bs[i] {
            diff += 1
            if diff == 1 {
                missMatchIdx1 = i
            } else if diff == 2 {
                missMatchIdx2 = i
            } else {
                return false
            }
        }
    }
    return diff == 0 || (missMatchIdx1 >= 0 && missMatchIdx2 >= 0 && 
        as[missMatchIdx1] == bs[missMatchIdx2] && as[missMatchIdx2] == bs[missMatchIdx1])
}