func threeConsecutiveOdds(arr []int) bool {
    i := 0
    for i < len(arr)-2 {
        if arr[i] % 2 == 1 {
            if arr[i+1] % 2 == 1 {
                if arr[i+2] % 2 == 1 {
                    return true
                } else {
                    i += 3
                }
            } else {
                i += 2
            }
        } else {
            i += 1
        }
    }
    return false
}