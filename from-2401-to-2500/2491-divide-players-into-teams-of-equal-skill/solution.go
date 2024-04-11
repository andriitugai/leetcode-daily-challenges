func dividePlayers(skill []int) int64 {
    sort.Ints(skill)
    left, right := 0, len(skill) - 1
    es := skill[left] + skill[right]
    var result int64 = 0
    for left < right {
        if skill[left] + skill[right] != es {
            return -1
        }
        result += int64(skill[left] * skill[right])
        left += 1
        right -= 1
    }
    return result
}