func countStudents(students []int, sandwiches []int) int {
    studMap := make(map[int]int)
    for i := 0; i < len(students); i++ {
        studMap[students[i]]++
    }

    unEat := len(students)
    for i := 0; i < len(sandwiches); i++ {
        if studMap[sandwiches[i]] == 0 {
            break
        }
        studMap[sandwiches[i]]--
        unEat--
    }

    return unEat
}