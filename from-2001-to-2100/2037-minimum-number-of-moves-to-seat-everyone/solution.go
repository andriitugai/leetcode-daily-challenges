func minMovesToSeat(seats []int, students []int) int {
    abs := func(a int) int {
        if a < 0 {
            return -a
        }
        return a
    }

    sort.Ints(seats)
    sort.Ints(students)

    moves := 0
    for i := 0; i < len(seats); i ++ {
        moves += abs(students[i] - seats[i])
    }
    
    return moves
}