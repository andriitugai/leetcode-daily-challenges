func haveConflict(event1 []string, event2 []string) bool {
    e1shh, _ := strconv.Atoi(event1[0][:2])
    e1smm, _ := strconv.Atoi(event1[0][3:])
    e1ehh, _ := strconv.Atoi(event1[1][:2])
    e1emm, _ := strconv.Atoi(event1[1][3:])

    e2shh, _ := strconv.Atoi(event2[0][:2])
    e2smm, _ := strconv.Atoi(event2[0][3:])
    e2ehh, _ := strconv.Atoi(event2[1][:2])
    e2emm, _ := strconv.Atoi(event2[1][3:])

    e1s := e1shh * 60 + e1smm
    e1e := e1ehh * 60 + e1emm

    e2s := e2shh * 60 + e2smm
    e2e := e2ehh * 60 + e2emm

    if e2s >= e1s && e2s <= e1e || e1s >= e2s && e1s <= e2e {
        return true
    }
    return false
}