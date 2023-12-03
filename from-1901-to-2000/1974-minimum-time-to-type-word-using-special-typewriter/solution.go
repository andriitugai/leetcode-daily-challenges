func minTimeToType(word string) int {
    chars := make(map[rune]int)
    abc := "abcdefghijklmnopqrstuvwxyz"
    for i, c := range abc {
        chars[c] = i
    }

    time := len(word)
    curr := 'a'
    for _, c := range word {
        t1 := chars[c] - chars[curr]
        if t1 < 0 {
            t1 += len(abc)
        }
        t2 := chars[curr] - chars[c]
        if t2 < 0 {
            t2 += len(abc)
        } 
        t := t1
        if t2 < t1 {
            t = t2
        }
        time += t
        curr = c
    }
    return time
}