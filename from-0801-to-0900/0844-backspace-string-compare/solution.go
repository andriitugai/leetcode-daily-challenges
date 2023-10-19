unc backspaceCompare(s string, t string) bool {
    sr := []byte{}
    for i := 0; i < len(s); i++ {
        if s[i] == '#' {
            if len(sr) > 0 {
                sr = sr[:len(sr)-1]
            }
        } else {
            sr = append(sr, s[i])
        }
    }
    tr := []byte{}
    for i := 0; i < len(t); i++ {
        if t[i] == '#' {
            if len(tr) > 0 {
                tr = tr[:len(tr)-1]
            }
        } else {
            tr = append(tr, t[i])
        }
    }
    if len(tr) != len(sr) {
        return false
    }
    for i := 0; i < len(sr); i++ {
        if sr[i] != tr[i] {
            return false
        }
    }
    return true
}