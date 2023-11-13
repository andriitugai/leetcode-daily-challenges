func countTime(time string) int {
    hh := time[:2]
    hhVars := 1
    if hh == "??" {
        hhVars = 24
    } else if hh[0] == '?' {
        if int(hh[1] - '0') > 3 {
            hhVars = 2
        } else {
            hhVars = 3
        }
    } else if hh[1] == '?' {
        if int(hh[0] - '0') == 2 {
            hhVars = 4
        } else {
            hhVars = 10
        }
    }

    mm := time[3:]
    mmVars := 1
    if mm == "??" {
        mmVars = 60
    } else if mm[0] == '?' {
        mmVars = 6
    } else if mm[1] == '?' {
        mmVars = 10
    }

    return hhVars * mmVars
}