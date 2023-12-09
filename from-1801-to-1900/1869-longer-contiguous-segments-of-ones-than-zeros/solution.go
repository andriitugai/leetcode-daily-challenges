func checkZeroOnes(s string) bool {
    max0, max1 := 0, 0
    prev := int(s[0] - '0')
    curr1, curr0 := 0, 0
    if prev == 1 {
        curr1 += 1
    } else {
        curr0 += 1
    }

    for i := 1; i < len(s); i ++ {
        c := int(s[i] - '0')
        if prev == c && c == 1 {
            curr1 += 1
        } else if prev == c && c == 0 {
            curr0 += 1
        } else if prev != c {
            if c == 1 {
                if curr0 > max0 {
                    max0 = curr0
                }
                curr0 = 0
                curr1 = 1
            } else {
                if curr1 > max1 {
                    max1 = curr1
                }
                curr0 = 1
                curr1 = 0
            }
            prev = c
        }
    }
    if curr0 > max0 {
        max0 = curr0
    }
    if curr1 > max1 {
        max1 = curr1
    }
    return max1 > max0
}