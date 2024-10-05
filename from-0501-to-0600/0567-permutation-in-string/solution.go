type charCounter struct {
    counter map[byte]int
}

func newCharCounter(s string) *charCounter {
    c := make(map[byte]int)
    for i := 0; i < len(s); i ++ {
        c[s[i]] += 1
    }
    return &charCounter{counter: c}
}

func checkInclusion(s1 string, s2 string) bool {
    n1 := len(s1)
    n2 := len(s2)
    if n1 > n2 {
        return false
    }
    c1, c2 := newCharCounter(s1), newCharCounter(s2[:n1])
    if c1.eq(c2) {
        return true
    }
    for i := n1; i < n2; i ++ {
        c2.add(s2[i])
        c2.rm(s2[i - n1])
        if c1.eq(c2) {
            return true
        }
    }
    return false
}

func (c *charCounter) add(b byte) {
    c.counter[b] += 1
}

func (c *charCounter) rm(b byte) {
    c.counter[b] -= 1
}

func (c1 *charCounter) eq(c2 *charCounter) bool {
    for k, v := range c1.counter {
        if c2.counter[k] != v {
            return false
        }
    }
    for k, v := range c2.counter {
        if c1.counter[k] != v {
            return false
        }
    }
    return true
}