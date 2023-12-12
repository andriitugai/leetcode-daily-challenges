func maximumTime(time string) string {
    mt := strings.Split(time, "")
    if mt[0] == "?" {
        if mt[1] == "?" || mt[1] < "4" {
            mt[0] = "2"
        } else {
            mt[0] = "1"
        }
        
    }
    if mt[1] == "?" {
        if mt[0] == "2" {
            mt[1] = "3"
        } else {
            mt[1] = "9"
        }
    }
    if mt[3] == "?" {
        mt[3] = "5"
    }
    if mt[4] == "?" {
        mt[4] = "9"
    }
    return strings.Join(mt, "")
}