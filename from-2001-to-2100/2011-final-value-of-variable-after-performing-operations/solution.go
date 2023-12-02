func finalValueAfterOperations(operations []string) int {
    result := 0
    for _, op := range operations {
        if op == "--X" || op == "X--" {
            result -= 1
        } else {
            result += 1
        }
    }
    return result
}