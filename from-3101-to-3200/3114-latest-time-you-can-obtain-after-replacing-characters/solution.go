func findLatestTime(s string) string {
    t := strings.Split(s, ":")
    hh, mm := t[0], t[1]

    if hh == "??" {
        hh = "11"
    } else {
        if hh[0] == '?' {
            if hh[1] < '2' {
                hh = "1" + hh[1:]
            } else {
                hh = "0" + hh[1:]
            }
        } else if hh[1] == '?' {
            if hh[0] == '1' {
                hh = "11"
            } else {
                hh = "09"
            }
        }
    }
    
    if mm == "??" {
        mm = "59"
    } else {
        if mm[0] == '?' {
            mm = "5" + mm[1:]
        } else if mm[1] == '?' {
            mm = mm[:1] + "9"
        }
    }
    return hh + ":" + mm
}