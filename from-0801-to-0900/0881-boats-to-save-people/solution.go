func numRescueBoats(people []int, limit int) int {
    result := 0
    sort.Ints(people)
    left, right := 0, len(people) - 1
    for left <= right {
        if left == right {
            return result + 1
        }
        
        if people[right] + people[right - 1] <= limit {
            right -= 2
        } else if people[right] + people[left] <= limit {
            right -= 1
            left += 1
        } else {
            right -= 1
        }
        result += 1
    }
    return result
}