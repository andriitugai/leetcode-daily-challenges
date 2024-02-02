func sequentialDigits(low int, high int) []int {
    seqNums := []int{}
    for n := 1; n <= 9; n ++ {
        num := n
        nextNum := n + 1

        for num <= high && nextNum <= 9 {
            num = num * 10 + nextNum
            if num > high {
                break
            }
            if num >= low && num <= high {
                seqNums = append(seqNums, num)
            }
            nextNum += 1
        }
    }
    sort.Ints(seqNums)
    return seqNums
}