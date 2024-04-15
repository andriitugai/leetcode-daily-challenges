func compareVersion(version1 string, version2 string) int {
    vs1, vs2 := strings.Split(version1, "."), strings.Split(version2, ".")
    var v1, v2 int
    p1, p2 := 0, 0
    for  p1 < len(vs1) && p2 < len(vs2) {
        v1, _ = strconv.Atoi(vs1[p1])
        v2, _ = strconv.Atoi(vs2[p2])
        if v1 > v2 {
            return 1
        } else if v1 < v2 {
            return -1
        }
        p1 += 1
        p2 += 1
    }
    for p1 < len(vs1) {
        v1, _ = strconv.Atoi(vs1[p1])
        if v1 > 0 {
            return 1
        }
        p1 += 1
    }
    for p2 < len(vs2) {
        v2, _ = strconv.Atoi(vs2[p2])
        if v2 > 0 {
            return -1
        }
        p2 += 1
    }
    return 0
}