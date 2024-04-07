func longestMonotonicSubarray(nums []int) int {
    incr, decr := 1, 1
    incrMax, decrMax := 1, 1
    prev := nums[0]

    for _, num := range nums[1:] {
        if num > prev {
            if decr > decrMax {
                decrMax = decr
            }
            incr += 1
            decr = 1
        } else if num < prev {
            if incr > incrMax {
                incrMax = incr
            }
            incr = 1
            decr += 1
        } else {
            if incr > incrMax {
                incrMax = incr
            }
            if decr > decrMax {
                decrMax = decr
            }
            incr, decr = 1, 1
        }
        prev = num
    }
    if incr > incrMax {
        incrMax = incr
    }
    if decr > decrMax {
        decrMax = decr
    }
    if incrMax > decrMax {
        return incrMax
    }
    return decrMax
}