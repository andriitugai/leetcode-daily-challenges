func topStudents(positive_feedback []string, negative_feedback []string, report []string, student_id []int, k int) []int {
    feedback := make(map[string]int)
    for _, w := range positive_feedback {
        feedback[w] = 3
    }
    for _, w := range negative_feedback {
        feedback[w] = -1
    }
    students := make([][]int, len(report))
    for i := 0; i < len(report); i ++ {
        points := 0
        for _, w := range strings.Split(report[i], " ") {
            if fb, ok := feedback[w]; ok {
                points += fb
            }
        }
        students[i] = []int{student_id[i], points}
    }

    sort.Slice(students, func(i, j int) bool {
        if students[i][1] == students[j][1] {
            return students[i][0] < students[j][0]
        }
        return students[i][1] > students[j][1]
    })
    result := make([]int, k)
    for i := 0; i < k; i ++ {
        result[i] = students[i][0]
    }
    return result
}