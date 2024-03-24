func minOperations(k int) int {
    if k <= 2 {
        return k - 1
    }
    var ops int
    minOps := k - 1
    for nums := 2; nums * nums <= k; nums ++ {
        ops = nums + k / nums - 2
        if k % nums != 0 {
            ops += 1
        }
        if minOps > ops {
            minOps = ops
        }
    }
    return minOps
}