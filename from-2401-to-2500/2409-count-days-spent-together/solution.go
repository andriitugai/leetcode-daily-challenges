func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
    dim := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
    dayOfYear := func(date string) int {
        mm, _ := strconv.Atoi(date[:2])
        dd, _ := strconv.Atoi(date[3:])

        doy := 0
        for i := 0; i < mm - 1; i ++ {
            doy += dim[i]
        }
        return doy + dd
    }
    sA := dayOfYear(arriveAlice)
    eA := dayOfYear(leaveAlice)
    sB := dayOfYear(arriveBob)
    eB := dayOfYear(leaveBob)

    result := 0
    if sA <= sB && eA >= sB {
        if eB > eA {
            result = eA - sB + 1
        } else {
            result = eB - sB + 1
        }
    } else if sB <= sA && eB >= sA {
        if eA > eB {
            result = eB - sA + 1
        } else {
            result = eA - sA + 1
        }
    }
    return result
}