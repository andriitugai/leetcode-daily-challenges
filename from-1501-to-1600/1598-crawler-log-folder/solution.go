func minOperations(logs []string) int {
    level := 0
    for _, log := range logs {
        if log == "../" {
            if level > 0 {
                level -= 1
            }
        } else if log == "./" {
            continue
        } else {
            level += 1
        }
    }
    return level
}