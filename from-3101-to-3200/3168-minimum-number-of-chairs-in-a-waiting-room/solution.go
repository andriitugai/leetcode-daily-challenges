func minimumChairs(s string) int {
    chairs := 0
    maxChairs := 0
    for i := 0; i < len(s); i ++ {
        if s[i] == 'E' {
            chairs += 1
            if chairs > maxChairs {
                maxChairs = chairs
            }
        } else if chairs > 0 && s[i] == 'L' {
            chairs -= 1
        }
    }
    return maxChairs
}