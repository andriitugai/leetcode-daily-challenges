func isValidSerialization(preorder string) bool {
    vals := strings.Split(preorder, ",")
    slots := 1

    for _, val := range vals {
        if slots == 0 {
            return false
        }
        if val == "#" {
            slots -= 1
        } else {
            slots += 1
        }
    }
    return slots == 0
}