func minOperations(boxes string) []int {
    n := len(boxes)
    result := make([]int, n)

    leftBalls, leftMoves := 0, 0
    for i := 0; i < n; i ++ {
        result[i] += leftMoves
        if boxes[i] == '1' {
            leftBalls += 1
        }
        leftMoves += leftBalls
    }

    rightBalls, rightMoves := 0, 0
    for i := n - 1; i >= 0; i -- {
        result[i] += rightMoves
        if boxes[i] == '1' {
            rightBalls += 1
        }
        rightMoves += rightBalls
    }

    return result
}