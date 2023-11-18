func rowAndMaximumOnes(mat [][]int) []int {
    maxOnes := 0
    maxOnesRow := 0
    for r := 0; r < len(mat); r ++ {
        ones := 0
        for c := 0; c < len(mat[0]); c ++ {
            if mat[r][c] == 1 {
                ones += 1
            }
        }
        if ones > maxOnes {
            maxOnes = ones
            maxOnesRow = r
        }
    }
    return []int{maxOnesRow, maxOnes}
}